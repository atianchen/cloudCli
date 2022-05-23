package db

/**
 * 数据库操作类
 * @author jensen.chen
 * @date 2022-05-23
 */
type DbHelper interface{

	/**
	 * 保存
	 */
	Save(entity interface{}) interface{}

	/**
	 * 更新
	 */
	Update(entity interface{}) int

	/**
	 * 删除
	 */
	Remove(entity interface{}) int

	/**
	 * 根据主键删除
	 */
	RemoveByPriKey(entity interface{}) int
}