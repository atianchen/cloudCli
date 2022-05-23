package repository

import (
		"cloudCli/domain"
		"cloudCli/task"
		"testing"
		"fmt"	
		dao "cloudCli/repository"
	)

func TestInsert(t *testing.T){
	dbManager := task.DbManager{}
	dbManager.Init()
	defer dbManager.Stop()
	doc := &domain.DocInfo{}
	doc.Name = "test.xlss"
	repository:=dao.DocRepository{}
	fmt.Println(repository.Save(doc))
}