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
	_, err := db.DbInst.Execute("insert into sys_param (id,name,code,val,group) values (?,?,?,?,?)",
		uuid.New(), param.Name, param.Code, param.Val, param.Group)
	log.Println(err)
	return err

}

/**
 * 更新
 */
func (r *ParamRepository) Update(param *domain.Param) error {
	_, err := db.DbInst.Execute("update sys_param set name=?,code=?,val=?,group=? where id=?",
		param.Name, param.Code, param.Val, param.Group, param.Id)
	return err
}

/**
获取分组配置
*/
func (r *ParamRepository) GetGroupParams(dest *[]domain.Param, group string) error {
	return db.DbInst.Query(dest, "select  * from sys_param where group=?", group)
}
