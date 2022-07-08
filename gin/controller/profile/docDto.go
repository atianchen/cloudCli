package profile

/**
 *
 * @author jensen.chen
 * @date 2022/7/8
 */
type DocDto struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	NestedPath string `json:"nestedPath"`
	Content    string `json:"content"`
	Type       int8   `json:"type"`
	Hash       string `json:"hash"`
	CreateTime int64  `json:"createTime"`
	CheckTime  int64  `json:"checkTime"`
}
