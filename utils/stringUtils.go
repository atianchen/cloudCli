package utils

import "strings"

/**
 * 字符串工具集合
 * @author jensen.chen
 * @date 2022/6/27
 */
/**
替换空串
*/
func TrimBlank(content string) string {
	return strings.Trim(content, " ")
}
