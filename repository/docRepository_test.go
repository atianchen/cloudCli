package repository

import (
	"cloudCli/utils/encrypt"
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	/*	cfg.Load("../config.yml")
		dbManager := node.DbManager{}
		dbManager.Init()
		defer dbManager.Stop()
		doc := domain.DocInfo{}
		doc.Name = "test.xlss"
		repository := DocRepository{}
		fmt.Println(repository.Save(&doc))
		fmt.Println("--------------")
		doc1, err := repository.GetByPrimary("1")
		fmt.Println(err)
		fmt.Println(doc1.Name)*/
	fmt.Println(encrypt.MD5("admin"))
}
