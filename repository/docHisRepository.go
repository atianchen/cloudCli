package repository

import (
	"cloudCli/db"
	"cloudCli/domain"
	"cloudCli/utils/timeUtils"
	"github.com/google/uuid"
	"time"
)

/**
 *文件变更历史
 * @author jensen.chen
 * @date 2022/7/6
 */
type DocHisRepository struct {
}

func (r *DocHisRepository) Save(dh *domain.DocHistory) error {
	time := timeUtils.TimeConfig{time.Now()}
	dh.Ts = time.Unix()

	/*	Id         string
		Name       string
		Path       string
		ModifyTime int64  `db:"modify_time"` //变更时间
		Raw        string //原始的文件内容
		Content    string //变更后的文件让内容
		Status     int    //状态
		HanleTime  int64  `db:"handle_time"` //处理时间
		Handler    string*/
	_, err := db.DbInst.Execute("insert into inspect_doc_his (id,doc_id,name,path,nested_path,modify_time,raw,content,status,handle_time,handler,opinion,ts) "+
		"values (?,?,?,?,?,?,?,?,?,?,?,?,?)",
		uuid.New(), dh.DocId, dh.Name, dh.Path, dh.NestedPath, dh.ModifyTime, dh.Raw, dh.Content, dh.Status, dh.HandleTime, dh.Handler, dh.Opinion, dh.Ts)

	return err

}

/**
 * 更新
 */
func (r *DocHisRepository) Update(dh *domain.DocHistory) error {
	_, err := db.DbInst.Execute("update inspect_doc_his set name=?,path=?,nested_path=?,modify_time=?,raw=?,content=?,status=?,handle_time=?,handler=?,opinion=?,ts=? where id=?",
		dh.Name, dh.Path, dh.NestedPath, dh.ModifyTime, dh.Raw, dh.Content, dh.Status, dh.HandleTime, dh.Handler, dh.Opinion, dh.Ts, dh.Id)
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
 * 执行查询
 */
func (r *DocHisRepository) Query(dest []domain.DocHistory, sql string, args ...any) error {
	return db.DbInst.Query(dest, sql, args)
}
