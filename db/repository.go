package db

/**
 * Dao层封装
 * @author jensen.chen
 * @date 2022-05-23
 */
type Repository interface{

	/**
	 * 保存
	 */
	Save(entity interface{}) interface{}

	/**
	 * 更新
	 */
	Update(entity interface{}) error

	/**
	 * 删除
	 */
	Remove(entity interface{}) error

	/**
	 * 根据主键删除
	 */
	RemoveByPrimary(priKey interface{}) error

	/**
	 * 根据ID查询 
	 */
	GetByPrimary(priKey interface{}) (interface{},error)


	/**
	 * 查询
	 */
	Query(sql string ,args...any) ([]interface{},error)

	/**
	 * 获取建表语句
	 */
	GenrateCreateSql() string
}