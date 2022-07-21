package repository

import (
	"cloudCli/db"
	"cloudCli/domain"
	"cloudCli/gin/dto"
	"cloudCli/utils/timeUtils"
	"github.com/google/uuid"
)

/**
 *文件变更历史
 * @author jensen.chen
 * @date 2022/7/6
 */
type DocHisRepository struct {
}

func (r *DocHisRepository) Save(dh *domain.DocHistory) error {
	dh.Ts = timeUtils.NowUnixTime()

	/*	Id         string
		Name       string
		Path       string
		ModifyTime int64  `db:"modify_time"` //变更时间
		Raw        string //原始的文件内容
		Content    string //变更后的文件让内容
		Status     int    //状态
		HanleTime  int64  `db:"handle_time"` //处理时间
		Handler    string*/
	_, err := db.DbInst.Execute("insert into inspect_doc_his (id,hash,doc_id,name,path,nested_path,modify_time,raw,content,status,handle_time,handler,opinion,ts,creator,handle_result) "+
		"values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		uuid.New(), dh.Hash, dh.DocId, dh.Name, dh.Path, dh.NestedPath, dh.ModifyTime, dh.Raw, dh.Content, dh.Status, dh.HandleTime, dh.Handler, dh.Opinion, dh.Ts, dh.Creator, dh.HandleResult)

	return err

}

/**
 * 更新处理意见
 */
func (r *DocHisRepository) UpdateHandleResult(dh *dto.DocHandleDto) error {
	_, err := db.DbInst.Execute("update inspect_doc_his set handle_time=?,handler=?,opinion=?,handle_result=? where id=?",
		dh.HandleTime, dh.Handler, dh.Opinion, dh.HandleResult, dh.Id)
	return err
}

/**
 * 更新
 */
func (r *DocHisRepository) Update(dh *domain.DocHistory) error {
	_, err := db.DbInst.Execute("update inspect_doc_his set hash=?,name=?,path=?,nested_path=?,modify_time=?,raw=?,content=?,status=?,handle_time=?,handler=?,opinion=?,ts=?,handle_result=? where id=?",
		dh.Hash, dh.Name, dh.Path, dh.NestedPath, dh.ModifyTime, dh.Raw, dh.Content, dh.Status, dh.HandleTime, dh.Handler, dh.Opinion, dh.Ts, dh.HandleResult, dh.Id)
	return err
}

/**
 * 删除
 */
func (r *DocHisRepository) Remove(doc *domain.DocHistory) error {
	return r.RemoveByPrimary(doc.Id)
}

/**
 * 根据主键删除
 */
func (r *DocHisRepository) RemoveByPrimary(priKey string) error {
	_, err := db.DbInst.Execute("delete from inspect_doc_his where id=?", priKey)
	return err
}

/**
 * 根据主键查询
 */
func (r *DocHisRepository) GetByPrimary(priKey string) (*domain.DocHistory, error) {
	doc := domain.DocHistory{}
	err := db.DbInst.Get(&doc, "select * from inspect_doc_his where id=?", priKey)
	return &doc, err
}

/**
 * 分页查询
 */
func (r *DocHisRepository) PageQuery(dest *[]domain.DocHistory, startIndex int, limit int, status int, name string) error {
	sql := "select * from inspect_doc_his where status=? "
	if len(name) > 0 {
		sql += " and name like ?"
	}
	sql += " limit ? offset  ?"
	if len(name) > 0 {
		return db.DbInst.Query(dest, sql, status, "%"+name+"%", limit, startIndex)
	} else {
		return db.DbInst.Query(dest, sql, status, limit, startIndex)
	}
}

/**
获取最后一条未处理的更改历史
*/
func (r *DocHisRepository) GetLastDocHis(path string, nestedPath string, status int) (*domain.DocHistory, error) {
	doc := domain.DocHistory{}
	if err := db.DbInst.Get(&doc, "select *from inspect_doc_his where path=? and nested_path=? and status=?  order by modify_time desc limit 1 offset 0", path, nestedPath, status); err == nil {
		return &doc, nil
	} else {
		return nil, err
	}
}

/**
计算数量
*/
func (r *DocHisRepository) CountDocHis(path string, nestedPath string, status int) (int, error) {
	var num int
	err := db.DbInst.Get(&num, "select count(*) from inspect_doc_his where path=? and nested_path=? and status=?", path, nestedPath, status)
	return num, err
}

/**
 * 执行查询
 */
func (r *DocHisRepository) Query(dest []domain.DocHistory, sql string, args ...any) error {
	return db.DbInst.Query(dest, sql, args)
}
