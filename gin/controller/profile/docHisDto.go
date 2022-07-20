package profile

/**
 *
 * @author jensen.chen
 * @date 2022/7/19
 */
type DocHisDto struct {
	Id           string `json:"id"`
	DocId        string `json:"docId"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	NestedPath   string `json:"nestedPath"`
	ModifyTime   int64  `json:"modifyTime"`
	Hash         string `json:"hash"`
	Status       int    `json:"status"`
	HandleResult int    `json:"handleResult"`
	HandleTime   int64  `json:"handleTime"`
	Handler      string `json:"handler"`
	Opinion      string `json:"opinion"`
}

//变更明细
type DocHisDetailDto struct {
	Id           string `json:"id"`
	DocId        string `json:"docId"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	NestedPath   string `json:"nestedPath"`
	ModifyTime   int64  `json:"modifyTime"`
	Hash         string `json:"hash"`
	Status       int    `json:"status"`
	HandleResult int    `json:"handleResult"`
	HandleTime   int64  `json:"handleTime"`
	Handler      string `json:"handler"`
	Opinion      string `json:"opinion"`
	Raw          string `json:"raw"`
	Content      string `json:"content"`
}
