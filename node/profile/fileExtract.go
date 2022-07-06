package profile

import (
	"archive/zip"
	"cloudCli/domain"
	"cloudCli/utils"
	"cloudCli/utils/timeUtils"
	"errors"
	"strings"
	"time"
)

var extractImpl = map[int]FileExtract{domain.DOC_TYPE_DISKFILE: &DiskFileExtract{}, domain.DOC_TYPE_JARFILE: &JarFileExtract{}}

/**
 *
 * @author jensen.chen
 * @date 2022/7/6
 */
type FileExtract interface {
	Extract(filePath string, nestedPath string) (*domain.DocInfo, error)
}

func ExtractDocInfo(doc *domain.DocInfo) (*domain.DocInfo, error) {
	if doc.Type == domain.DOC_TYPE_JARFILE {
		return extractImpl[domain.DOC_TYPE_JARFILE].Extract(doc.Path, doc.NestedPath)
	} else {
		return extractImpl[domain.DOC_TYPE_DISKFILE].Extract(doc.Path, doc.NestedPath)
	}
}

func ExtractFile(filePath string, nestedPath string) (*domain.DocInfo, error) {
	if strings.Index(filePath, ".jar") > 0 && len(nestedPath) > 0 {
		return extractImpl[domain.DOC_TYPE_JARFILE].Extract(filePath, nestedPath)
	} else {
		return extractImpl[domain.DOC_TYPE_DISKFILE].Extract(filePath, nestedPath)
	}
}

/**
磁盘文件的解析
*/
type DiskFileExtract struct {
}

/**
解析磁盘文件
*/
func (e *DiskFileExtract) Extract(filePath string, nestedPath string) (*domain.DocInfo, error) {
	time := timeUtils.TimeConfig{Time: time.Now()}
	content, cErr := utils.GetFileStringContent(filePath)
	if cErr != nil {
		return nil, cErr
	}
	hashVal, hashErr := utils.GetFileHash(filePath)
	if hashErr == nil {
		newDoc := domain.DocInfo{Name: utils.GetFileName(filePath), NestedPath: nestedPath, Content: content, Path: filePath, CreateTime: time.Unix(), Hash: hashVal, Type: domain.DOC_TYPE_DISKFILE}
		return &newDoc, nil
	} else {
		return nil, hashErr
	}
}

/**
Jar 文件的解析
*/
type JarFileExtract struct {
}

func (e *JarFileExtract) Extract(filePath string, nestedPath string) (*domain.DocInfo, error) {
	time := timeUtils.TimeConfig{Time: time.Now()}
	rc, err := zip.OpenReader(filePath)
	if err == nil {
		defer rc.Close()
	}
	cfgFile, _ := rc.Open(nestedPath)
	if cfgFile != nil {
		content, cErr := utils.GetStringContent(cfgFile)
		if cErr != nil {
			return nil, cErr
		}
		hashVal, hashErr := utils.ConvertReaderToHash(cfgFile)

		if hashErr == nil {
			newDoc := domain.DocInfo{Name: utils.GetFileName(filePath), NestedPath: nestedPath, Content: content, Path: filePath, CreateTime: time.Unix(), Hash: hashVal, Type: domain.DOC_TYPE_JARFILE}
			return &newDoc, nil
		} else {
			return nil, hashErr
		}
	} else {
		return nil, errors.New("Not Found Jar File " + filePath + "  " + nestedPath)
	}

}
