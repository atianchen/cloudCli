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
	Execute(sql  string,args ...any) (sql.Result,error)


	/**
	 * 查询
	 * @args 参数
	 */
	Query(sql string,args ...any) (*sql.Rows,error)

	/**
	 * 释放数据库链接
	 */
	Release()
}