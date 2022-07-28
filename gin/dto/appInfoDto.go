package dto

/**
 *
 * @author jensen.chen
 * @date 2022/7/28
 */
type AppInfoDto struct {
	Name   string `json:"name"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Leader bool   `json:"leader"`
}
