package db

import (
    "database/sql"
)

/**
 * 关系数据库操作
 */
type SqlDb interface{

	Connect(cfg DbConfig)
	
	/**
	 * 执行SQL
	 * @args 参数
	 */
	Execute(sql string,args ...any) (sql.Result,error)

	/**
	 * 返回单条数据
	 */
	Get(dest interface{},sql string ,args...any) error


	/**
	 * 查询
	 * @dest 数组指针
	 */
	Query(dest interface{},sql string,args ...any) error

	/**
	 * 释放数据库链接
	 */
	Release()
}