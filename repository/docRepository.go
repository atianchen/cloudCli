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
	_, err := db.DbInst.Execute("insert into inspect_doc (id,name,path,creator,create_time,modify_time,check_time,hash,ts) values (?,?,?,?,?,?,?,?,?)",
		uuid.New(), doc.Name, doc.Path, doc.Creator, doc.CreateTime, doc.ModifyTime, doc.CheckTime, doc.Hash, doc.Ts)
	log.Println(err)
	return err

}

/**
 * 更新
 */
func (r *DocRepository) Update(doc *domain.DocInfo) error {
	_, err := db.DbInst.Execute("update inspect_doc set name=?,path=?,creator=?,create_time=?,modify_time=?,check_time=?,hash=?,ts=? where id=?",
		doc.Name, doc.Path, doc.Creator, doc.CreateTime, doc.ModifyTime, doc.CheckTime, doc.Hash, doc.Ts, doc.Id)
	return err
}

/**
 * 删除
 */
func (r *DocRepository) Remove(doc *domain.DocInfo) error {
	return r.RemoveByPrimary(doc.Id)
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
 * 执行查询
 */
func (r *DocRepository) Query(dest []domain.DocInfo, sql string, args ...any) error {
	return db.DbInst.Query(dest, sql, args)
}

/**
根据路径查找文件
*/
func (r *DocRepository) FindByPath(path string) (*domain.DocInfo, error) {
	doc := domain.DocInfo{}
	err := db.DbInst.Get(&doc, "select * from inspect_doc where path=?", path)
	return &doc, err
}
