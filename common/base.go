package common

/**
 * @author jensen.chen
 * @date 2022-05-20
 */
/**
 * 基础对象
 */
type BaseObj struct
{
	Ts int64
	Creator string
}
/**
 * 扩展属性
 */
type Extends interface
{
		/**
	 * 获取属性
	 */
	GetAttr(key string) interface{}

	/**
	 * 添加属性
	 */
	AddAttr(key string,val interface{})

	/**
	 * 移除属性
	 */
	RemoveAttr(key string)

	/**
	 * 判定是否包含属性
	 */
	Contains(key string) bool
}

/**
 * 包含额外属性的对象
 */
type ModalMap struct
{
	AttrMap map[string]interface{}
}
func (c *ModalMap) GetAttr(key string) interface{}{
	 rs, _ := c.AttrMap[key]
	 return rs
}

func (c *ModalMap) AddAttr(key string,val interface{}){

	c.AttrMap[key] = val
}

func (c *ModalMap) RemoveAttr(key string){
	delete(c.AttrMap, key)
}

func (c *ModalMap) Contains(key string) bool{
	  _, ok := c.AttrMap[key]
	  return ok
}
