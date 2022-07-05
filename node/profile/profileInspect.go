package profile

import (
	"cloudCli/cfg"
	"cloudCli/domain"
	"cloudCli/io"
	"cloudCli/node"
	"cloudCli/repository"
	"cloudCli/utils"
	"cloudCli/utils/log"
	"cloudCli/utils/timeUtils"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

/**
 * 配置文件巡检
 * @author jensen.chen
 * @date 2022/6/30
 */
type ProfileInspect struct {
	node.AbstractNode
	configs       []ProfileConfig
	docRepository *repository.DocRepository
}

/**
处理Channel消息
*/
func (p *ProfileInspect) HandleMessage(msg interface{}) {

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
	for _, config := range p.configs {
		files, err := io.FindFile(config.Directory, config.Expression)
		if err != nil {
			log.Error("ProfileInspect Run Error", err.Error())
			return
		}
		for _, cfgFile := range files {
			log.Info("Found Config File ", cfgFile)
			p.registerFile(cfgFile)
		}
	}
}

func (p *ProfileInspect) registerFile(filePath string) {
	doc, err := p.docRepository.FindByPath(filePath)
	if err != nil {
		if doc == nil || len(doc.Id) < 1 {
			log.Info("Add Inspect Doc ", filePath)
			time := timeUtils.TimeConfig{Time: time.Now()}
			hashVal, hashErr := utils.GetFileHash(filePath)
			if hashErr == nil {
				newDoc := domain.DocInfo{Name: utils.GetFileName(filePath), Path: filePath, CreateTime: time.Unix(), Hash: hashVal}
				err := p.docRepository.Save(&newDoc)
				if err != nil {
					log.Error("Save File Error ", filePath, err)
				}
			} else {
				log.Error("Get File Hash Error ", filePath, err)
			}

		}
	} else {
		log.Error("Register File Error ", filePath)
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
