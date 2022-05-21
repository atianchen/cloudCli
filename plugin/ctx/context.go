package ctx

//import "unsafe"
import ("cloudCli/common")

/**
 * @author jensen.chen 
 * @date 2022-05-20
 * 上下文环境变量
 */
type Context interface{
 	common.Extends
}


/**
 * 默认环境变量
 */
type DefaultContext struct{
	common.BaseObj
	common.ModalMap
}


/**
 * 创建默认Context
 */
func CreateContext() Context{
	ctx := &DefaultContext{}
	ctx.AttrMap = make(map[string]interface{})
	return ctx//(Context)(unsafe.Pointer(ctx))
}