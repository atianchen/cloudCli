package repository

import (
		"cloudCli/domain"
		"cloudCli/task"
		"cloudCli/cfg"
		"testing"
		"fmt"	
		dao "cloudCli/repository"
	)

func TestInsert(t *testing.T){
	cfg.Load("../config.yml")
	dbManager := task.DbManager{}
	dbManager.Init()
	defer dbManager.Stop()
	doc := domain.DocInfo{}
	doc.Name = "test.xlss"
	repository:=dao.DocRepository{}
	fmt.Println(repository.Save(&doc))
	fmt.Println("--------------")
	doc1,err:=repository.GetByPrimary(1)
	fmt.Println(err)
	fmt.Println(doc1.Name)
}