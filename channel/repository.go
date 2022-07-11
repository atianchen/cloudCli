package channel

/**
 *
 * @author jensen.chen
 * @date 2022/6/10
 */
/**
通道工厂
*/
var chanRepository = make(map[string]chan interface{})

/**
创建通道
*/
func CreateChan(name string) chan interface{} {
	rs, ok := chanRepository[name]
	if !ok {
		nc := make(chan interface{})
		chanRepository[name] = nc
		return nc
	}
	return rs
}

/**
获取通道
*/
func GetChan(name string) (chan interface{}, bool) {
	rs, ok := chanRepository[name]
	if ok {
		return rs, true
	} else {
		return nil, false
	}

}

/**
关闭Channel
*/
func CloseChan(name string) {
	ch, ok := chanRepository[name]
	if ok {
		close(ch)
		delete(chanRepository, name)
	}
}

/**
释放整个工厂
*/
func Release() {
	/**
	关闭Chan
	*/
	closeCmd := BuildCloseCommand()
	for _, v := range chanRepository {
		v <- closeCmd
		close(v)
	}
}
