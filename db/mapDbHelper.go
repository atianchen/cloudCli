package db

import "time"
/**
 * 基于K-V数据库操作类
 * @author jensen.chen
 * @date 2022-05-23
 */
type MapDbHelper interface{

	/**
	 * 初始化
	 */
	Connect(path string)
	/**
	 * 设置值
	 * @expireTime 过期时间
	 */
	Set(key string,value string,expireTime time.Duration) error 

	Get(key string) (string,error)

	Remove(key string) error 

	/**
	 * 释放链接
	 */
	Release()
}

