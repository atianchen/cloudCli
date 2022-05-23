package repository

import (
		"cloudCli/domain"
		"cloudCli/db"
		"errors"
		"log"
	)

/**
 * 文件信息Repository
 * @author jensen.chen
 * @date 2022-05-23
 */
type DocRepository struct{
	
}



func (r *DocRepository) Save(doc *domain.DocInfo) error{
	_,err :=  db.DbInst.Execute("insert into inspect_doc (name,path,modifyDate,lastCheckDate,hash,ts) values (?,?,?,?,?,?)",
		doc.Name,doc.Path,doc.ModifyDate,doc.LastCheckDate,doc.Hash,doc.Ts)
	log.Println(err)
	return err

}	

/**
 * 更新
 */
func (r *DocRepository) Update(doc *domain.DocInfo) error{
	_,err :=  db.DbInst.Execute("update inspect_doc set name=?,path=?,modifyDate=?,lastCheckDate=?,hash=?,ts=? where id=?",
		doc.Name,doc.Path,doc.ModifyDate,doc.LastCheckDate,doc.Hash,doc.Ts,doc.Id)
	return err
}

/**
 * 删除
 */
func (r *DocRepository) Remove(doc *domain.DocInfo) error{
	return r.RemoveByPriKey(doc.Id)
}

/**
 * 根据主键删除
 */
func (r *DocRepository) RemoveByPriKey(priKey int64) error{
	_,err :=  db.DbInst.Execute("delete from inspect_doc where id=?",priKey)
	return err
}

/**
 * 根据主键查询
 */
func (r *DocRepository) GetByPrimary(priKey int64) (*domain.DocInfo,error){
	rows,err := db.DbInst.Query("select id,name,path,modifyDate,lastCheckDate,hash,ts from inspect_doc where id=?",priKey)
	if (rows!=nil){
		defer rows.Close()
		if (rows.Next()){
			doc := &(domain.DocInfo{})
			rows.Scan(&(doc.Id),&(doc.Name),&(doc.Path),&(doc.ModifyDate),&(doc.LastCheckDate),&(doc.Hash),&(doc.Ts))
			return doc,nil
		}else{
			return nil,errors.New("Entity Not Found")
		}

	}else{
		return nil,err
	}
}

/**
 * 执行查询
 */
func (r *DocRepository)  Query(sql string ,args...any) ([]*domain.DocInfo,error){
		rows,err := db.DbInst.Query(sql,args)
	if (rows!=nil){
		defer rows.Close()
		docArr :=[]*domain.DocInfo{}
		for (rows.Next()){
			doc := &domain.DocInfo{}
			rows.Scan(&(doc.Id),&(doc.Name),&(doc.Path),&(doc.ModifyDate),&(doc.LastCheckDate),&(doc.Hash),&(doc.Ts))
			docArr = append(docArr,doc)
		}
		return docArr,nil
	}else{
		return nil,err
	}
}