package repository

import (
	"cloudCli/db"
	"cloudCli/domain"
	"cloudCli/utils/timeUtils"
	"github.com/google/uuid"
	"log"
	"time"
)

/**
 * 文件信息Repository
 * @author jensen.chen
 * @date 2022-05-23
 */
type DocRepository struct {
}

func (r *DocRepository) Save(doc *domain.DocInfo) error {
	time := timeUtils.TimeConfig{time.Now()}
	doc.Ts = time.Unix()
	_, err := db.DbInst.Execute("insert into inspect_doc (id,name,content,path,nested_path,creator,create_time,modify_time,check_time,hash,type,ts) values (?,?,?,?,?,?,?,?,?,?,?,?)",
		uuid.New(), doc.Name, doc.Content, doc.Path, doc.NestedPath, doc.Creator, doc.CreateTime, doc.ModifyTime, doc.CheckTime, doc.Hash, doc.Type, doc.Ts)
	log.Println(err)
	return err

}

/**
 * 更新
 */
func (r *DocRepository) Update(doc *domain.DocInfo) error {
	_, err := db.DbInst.Execute("update inspect_doc set name=?,content=?,path=?,nested_path=?,creator=?,create_time=?,modify_time=?,check_time=?,hash=?,type=?,ts=? where id=?",
		doc.Name, doc.Content, doc.Path, doc.NestedPath, doc.Creator, doc.CreateTime, doc.ModifyTime, doc.CheckTime, doc.Hash, doc.Type, doc.Ts, doc.Id)
	return err
}

func (r *DocRepository) UpdateContent(docId string, hash string, content string) error {
	_, err := db.DbInst.Execute("update inspect_doc set hash=?,content=? where id=?", hash, content, docId)
	return err
}

func (r *DocRepository) UpdateCheckTime(doc *domain.DocInfo) error {
	_, err := db.DbInst.Execute("update inspect_doc set check_time=? where id=?", doc.CheckTime, doc.Id)
	return err
}

/**
 * 删除
 */
func (r *DocRepository) Remove(doc *domain.DocInfo) error {
	return r.RemoveByPrimary(doc.Id)
}

/**
 * 删除全部
 */
func (r *DocRepository) RemoveAll() error {
	_, err := db.DbInst.Execute("delete from inspect_doc")
	return err
}

/**
 * 根据主键删除
 */
func (r *DocRepository) RemoveByPrimary(priKey string) error {
	_, err := db.DbInst.Execute("delete from inspect_doc where id=?", priKey)
	return err
}

/**
 * 根据主键查询
 */
func (r *DocRepository) GetByPrimary(priKey string) (*domain.DocInfo, error) {
	doc := domain.DocInfo{}
	err := db.DbInst.Get(&doc, "select * from inspect_doc where id=?", priKey)
	return &doc, err
}

/**
获取所有
*/
func (r *DocRepository) GetAll(dest *[]domain.DocInfo) error {
	return db.DbInst.Query(dest, "select  * from inspect_doc")
}

/**
 * 执行查询
 */
func (r *DocRepository) Query(dest *[]domain.DocInfo, sql string, args ...any) error {
	return db.DbInst.Query(dest, sql, args)
}

/**
 * 分页查询
 */
func (r *DocRepository) PageQuery(dest *[]domain.DocInfo, startIndex int, limit int, name string) error {
	sql := "select * from inspect_doc"
	if len(name) > 0 {
		sql += " where name like ?"
	}
	sql += " limit ? offset  ?"
	if len(name) > 0 {
		return db.DbInst.Query(dest, sql, "%"+name+"%", limit, startIndex)
	} else {
		return db.DbInst.Query(dest, sql, limit, startIndex)
	}
}

/**
根据路径查找文件
*/
func (r *DocRepository) FindByPath(path string, nestedPath string) (*domain.DocInfo, error) {
	doc := domain.DocInfo{}
	var err error
	if len(nestedPath) > 0 {
		err = db.DbInst.Get(&doc, "select * from inspect_doc where path=? and nested_path=?", path, nestedPath)
	} else {
		err = db.DbInst.Get(&doc, "select * from inspect_doc where path=?", path)
	}
	return &doc, err
}
