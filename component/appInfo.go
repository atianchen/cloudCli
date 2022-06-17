package component

/**
 *
 * @author jensen.chen
 * @date 2022/6/17
 */

type OsTypeEnum int8 //操作系统类型
const (
	OS_WINDOWS       OsTypeEnum = 1
	OS_WINDOWSSERVER OsTypeEnum = 2
	OS_LINUX         OsTypeEnum = 3
	OS_CENTOS        OsTypeEnum = 4
	OS_REDHAT        OsTypeEnum = 5
	OS_UBUNTU        OsTypeEnum = 6
)

/**
应用信息
*/
type AppInfo struct {
	OsType  OsTypeEnum //操作系统类型
	Ip      string     //绑定的IP
	Version string     //操作系统的版本
}
