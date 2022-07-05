package utils

import "runtime"

/**
 *
 * @author jensen.chen
 * @date 2022/7/5
 */
/**
判断系统是否linux
*/
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

/**
获取系统路径的分隔符
*/
func SysSeparator() string {
	separator := "\\"
	if IsLinux() {
		separator = "/"
	}
	return separator
}
