package profile

/**
 *
 * @author jensen.chen
 * @date 2022/7/8
 */
type DocDto struct {
	Id         string
	Name       string
	Path       string
	NestedPath string
	Content    string
	Type       int8
	Hash       string
	CreateTime int64
	CheckTime  int64
}
