package repository

import (
	"cloudCli/db"
	"cloudCli/domain"
	"github.com/google/uuid"
	"log"
)

/**
 * 参数配置
 * @author jensen.chen
 * @date 2022/7/8
 */
type ParamRepository struct {
}

/**
 * 根据主键查询
 */
func (r *ParamRepository) GetByPrimary(priKey string) (*domain.Param, error) {
	doc := domain.Param{}
	err := db.DbInst.Get(&doc, "select * from sys_param where id=?", priKey)
	return &doc, err
}

/**
保存或更新
*/
func (r *ParamRepository) SaveOrUpdate(param *domain.Param) error {
	if len(param.Id) > 0 {
		return r.Update(param)
	} else {
		return r.Save(param)
	}
}

/**
保存
*/
func (r *ParamRepository) Save(param *domain.Param) error {
	_, err := db.DbInst.Execute("insert into sys_param (id,name,code,val,param_group) values (?,?,?,?,?)",
		uuid.New(), param.Name, param.Code, param.Val, param.Group)
	log.Println(err)
	return err

}

/**
 * 更新
 */
func (r *ParamRepository) Update(param *domain.Param) error {
	_, err := db.DbInst.Execute("update sys_param set name=?,code=?,val=?,param_group=? where id=?",
		param.Name, param.Code, param.Val, param.Group, param.Id)
	return err
}

/**
获取分组配置
*/
func (r *ParamRepository) GetGroupParams(dest *[]domain.Param, group string) error {
	return db.DbInst.Query(dest, "select  * from sys_param where param_group=?", group)
}

/**
获取分组配置
*/
func (r *ParamRepository) GetGroupSpecParam(group string, name string) (*domain.Param, error) {
	doc := domain.Param{}
	err := db.DbInst.Get(&doc, "select * from sys_param where param_group=? and code=?", group, name)
	return &doc, err
}

/**

 */
func (r *ParamRepository) GetAll(dest *[]domain.Param) error {
	return db.DbInst.Query(dest, "select  * from sys_param order by param_group,name")
}
