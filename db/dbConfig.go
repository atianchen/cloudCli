package db

type DbConfig interface{

}

/**
 * 基于文件的数据库配置
 */
type FileDbConfig struct{
	DbFile string
}