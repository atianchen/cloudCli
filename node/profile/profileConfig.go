package profile

/**
 *
 * @author jensen.chen
 * @date 2022/7/5
 */
type ProfileConfig struct {
	Directory        string `json:"dir"`
	IncludeSubFolder bool   `json:"includeSubFolder"`
	Expression       string `json:"expression"`
	NestedPath       string `json:"nestedPath"`
}
