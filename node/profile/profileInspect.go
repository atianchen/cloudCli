package profile

import (
	"cloudCli/cfg"
	channel2 "cloudCli/channel"
	"cloudCli/domain"
	"cloudCli/io"
	"cloudCli/node"
	"cloudCli/repository"
	"cloudCli/utils"
	"cloudCli/utils/log"
	"cloudCli/utils/timeUtils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

const PROFILE_NODE_NAME = "profileInspect"

/**
 * 配置文件巡检
 * @author jensen.chen
 * @date 2022/6/30
 */
type ProfileInspect struct {
	node.AbstractNode
	configs          []ProfileConfig
	docRepository    *repository.DocRepository
	docHisRepository *repository.DocHisRepository
}

func (p *ProfileInspect) Init() {
	/**
	读取配置文件
	*/
	configFile, err := cfg.GetConfig("cli.profile-inspect.config")
	if err == nil {
		filePath, err := utils.GetFilePath(configFile.(string))
		if err == nil {
			jsonFile, err := os.Open(filePath)
			if err == nil {
				jsonData, err := ioutil.ReadAll(jsonFile)
				if err == nil {
					p.configs = make([]ProfileConfig, 0)
					jsonError := json.Unmarshal(jsonData, &p.configs)
					if jsonError != nil {
						str := jsonError.Error()
						log.Error(str)
					}
				} else {
					log.Error("Read Profile Config Error " + err.Error())
				}
			}
		} else {
			log.Error("Not Found Profile Config" + err.Error())
		}
	} else {
		log.Error("Not Found Profile Config" + err.Error())
	}
}

/**
 * 开始
 */
func (p *ProfileInspect) Start(context *node.NodeContext) {
	/**
	搜索文件
	*/
	p.docRepository = &repository.DocRepository{}
	p.docHisRepository = &repository.DocHisRepository{}
	p.FirstFileScan()
}

/**
第一次文件扫描
*/
func (p *ProfileInspect) FirstFileScan() {
	for _, config := range p.configs {
		files, err := io.FindFile(config.Directory, config.Expression)
		if err != nil {
			log.Error("ProfileInspect Run Error", err.Error())
			return
		}
		for _, cfgFile := range files {
			log.Info("Found Config File ", cfgFile)
			p.registerFile(cfgFile, config.NestedPath)
		}
	}
}

/**
注册监控文件信息
*/
func (p *ProfileInspect) registerFile(filePath string, nestedPath string) {
	doc, err := p.docRepository.FindByPath(filePath, nestedPath)
	if err != nil {
		if doc == nil || len(doc.Id) < 1 {
			log.Info("Add Inspect Doc ", filePath)
			docInfo, err := ExtractFile(filePath, nestedPath)
			if err == nil {
				err = p.docRepository.Save(docInfo)
				if err == nil {
					log.Info(" Save Orgin File Success ", docInfo.Name, doc.Path, nestedPath)
				} else {
					log.Info(" Save Orgin File Error ", err)
				}
			} else {
				log.Info(" Save Orgin File Error ", err)
			}
		} else {
			log.Info(" File Already registered ", doc.Name)
		}
	} else {
		log.Info(" File Already registered ", doc.Name)
	}
}

/**
 * 停止
 */
func (p *ProfileInspect) Stop() {

}

/**
 * 获取名称
 */
func (p *ProfileInspect) Name() string {
	return "profileInspect"
}

func (p *ProfileInspect) GetMsgHandler() node.MsgHandler {
	return p
}

/**
处理Channel消息
*/
func (p *ProfileInspect) HandleMessage(msg interface{}) {
	if msg != nil {
		switch msg.(type) { //判断消息类型
		case *channel2.CommandMessage: //如果是CommandMessage
			{
				switch msg.(*channel2.CommandMessage).Name { //判断消息的指令
				case channel2.MESSAGE_ONTIME: //如果是定时执行的指令，则执行定时任务
					{
						//按时巡检系统
						var docs []domain.DocInfo
						err := p.docRepository.GetAll(&docs)

						if err == nil {
							for _, doc := range docs {
								if len(doc.Path) > 0 {
									p.checkFile(&doc)
								}
							}
						} else {
							log.Error("Inspect Exception ", err)
						}
					}
				case MESSAGE_PROFILE_RESET: //如果是收到的重置命令
					{
						/**
						重置
						*/
						p.reset()
					}
				}

			}
		}

	}
}

/**
重置所有配置
*/
func (p *ProfileInspect) reset() {
	log.Info("Reset Profile Config")
	err := p.docRepository.RemoveAll()
	if err != nil { //有错误，回送载有错误信息的AsyncResponse
		p.Transpot <- channel2.BuildErrorResponse(err)
	} else {
		//执行成功，发送回表示执行成功的AsyncResponse
		p.FirstFileScan()
		p.Transpot <- channel2.BuildEmptyResponse()
	}

}

/**
检测文件内容是否被变更
*/
func (p *ProfileInspect) checkFile(info *domain.DocInfo) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Check File Exception:" + info.Name)
		}
	}()
	time := timeUtils.TimeConfig{time.Now()}
	info.CheckTime = time.Unix()
	p.docRepository.UpdateCheckTime(info)
	doc, err := ExtractDocInfo(info)
	if err == nil {
		/**
		HASH值变化，记录历史
		*/
		if doc.Hash != info.Hash {
			/**
			判断是否有未处理的预警信息
			*/
			if num, _ := p.docHisRepository.CountDocHis(doc.Path, doc.NestedPath, domain.DOCHIS_STATUS_PENDING); num < 1 {
				log.Info("File Changed Detected ", doc.Path)
				log.Info("Send Alarm Mail ", doc.Path)
				if err := SendMailAlarm(p.saveChangeHis(info, doc)); err != nil {
					log.Info("Mail Send Error ", err.Error())
				} //发送警告邮件
			} else {
				log.Info("Skip Alarm ", doc.Path)
			}

		} else {
			log.Info("Consistency Check PASS ", doc.Path)
		}
	}
}

/**
记录文件变更历史
*/
func (p *ProfileInspect) saveChangeHis(od *domain.DocInfo, nd *domain.DocInfo) *domain.DocHistory {

	time := timeUtils.TimeConfig{time.Now()}

	docHis := domain.DocHistory{
		DocId:      od.Id,
		Path:       od.Path,
		NestedPath: od.NestedPath,
		Name:       od.Name,
		ModifyTime: time.Unix(),
		Raw:        od.Content,
		Content:    nd.Content,
		Status:     domain.DOCHIS_STATUS_PENDING,
	}
	err := p.docHisRepository.Save(&docHis)
	if err != nil {
		log.Error("Save Change History Exception ", err)
	} else {
		log.Error("Save Change History ", od.Path)
	}
	return &docHis
}

/**
更新变更历史的处理信息
@handler 处理人
@opinion 处理意见
*/
func (p *ProfileInspect) UpdateHistoryStatus(id string, handler string, opinion string) error {
	docHis, err := p.docHisRepository.GetByPrimary(id)
	if err == nil {
		time := timeUtils.TimeConfig{time.Now()}
		docHis.Handler = handler
		docHis.Opinion = opinion
		docHis.HandleTime = time.Unix()
		return p.docHisRepository.Update(docHis)
	} else {
		return errors.New("Spec DocHistory Not Found")
	}
}

/**
还原配置文件的原始内容
@id 变更记录ID
*/
func (p *ProfileInspect) RestoreChangedContent(id string, handler string, opinion string) error {
	return nil
}
