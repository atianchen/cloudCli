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
 *
 * @auth0月托管户0or jensen.chen
 * @date 2022/7/7
 */
type SysUserRepository struct {
}

func (r *SysUserRepository) Save(u *domain.SysUser) error {

	/*	Id     string
		Code   string
		RoleId string `db:"role_id"`
		Name   string
		Pwd    string
		Status int*/
	time := timeUtils.TimeConfig{time.Now()}
	u.Ts = time.Unix()
	_, err := db.DbInst.Execute("insert into sys_user (id,code,role_id,name,pwd,status,ts) values (?,?,?,?,?,?,?)",
		uuid.New(), u.Code, u.RoleId, u.Name, u.Pwd, u.Status, u.Ts)
	log.Println(err)
	return err

}

/**
 * 更新
 */
func (r *SysUserRepository) Update(u *domain.SysUser) error {
	_, err := db.DbInst.Execute("update sys_user set code=?,role_id=?,name=?,pwd=?,status=?,ts=? where id=?",
		u.Code, u.RoleId, u.Name, u.Pwd, u.Status, u.Ts, u.Id)
	return err
}

/**
 * 删除
 */
func (r *SysUserRepository) Remove(u *domain.SysUser) error {
	return r.RemoveByPrimary(u.Id)
}

/**
 * 根据主键删除
 */
func (r *SysUserRepository) RemoveByPrimary(priKey string) error {
	_, err := db.DbInst.Execute("delete from sys_user where id=?", priKey)
	return err
}

/**
 * 根据主键查询
 */
func (r *SysUserRepository) GetByPrimary(priKey string) (*domain.SysUser, error) {
	user := domain.SysUser{}
	err := db.DbInst.Get(&user, "select * from sys_user where id=?", priKey)
	return &user, err
}

/**
 * 执行查询
 */
func (r *SysUserRepository) Query(dest *[]domain.SysUser, sql string, args ...any) error {
	return db.DbInst.Query(dest, sql, args)
}

/**
根据CODE查找用户
*/
func (r *SysUserRepository) FindByCode(code string) (*domain.SysUser, error) {
	user := domain.SysUser{}
	err := db.DbInst.Get(&user, "select * from sys_user where code=?", code)
	return &user, err
}
