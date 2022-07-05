package io

import (
	"cloudCli/utils"
	"io/ioutil"
	"os"
	"regexp"
)

/**
 * 文件搜索工具库
 * @author jensen.chen
 * @date 2022/7/5
 */

/**
搜索文件
*/
func FindFile(directory string, expression string) ([]string, error) {

	reg, err := regexp.Compile(expression)
	if err == nil {
		var files []string
		dirInfo, err := os.Stat(directory)
		if err == nil && dirInfo.IsDir() {
			rs, _ := ioutil.ReadDir(directory)
			for _, fileInfo := range rs {
				if reg.MatchString(fileInfo.Name()) {
					files = append(files, directory+utils.SysSeparator()+fileInfo.Name())
				}
			}
		}
		return files, nil
	} else {
		return nil, err
	}

}
