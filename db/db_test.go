package db

import (
	"cloudCli/db"
	"testing"
)



func TestAdd(t *testing.T){
	db := db.BadgerDb{}
	defer dbHelper.Release()
	db.Connect("d:/temp/db")
	err := db.Set("name","chenzhi",0)
	if (err!=nil){
		t.Errorf("Save Error: %s\n", err.Error())
	}else{
		t.Logf("Success")
	}
}

func TestGet(t *testing.T){
	db := db.BadgerDb{}
	defer db.Release()
	db.Connect("d:/temp/db")
	content,err := db.Get("name")
	if (err!=nil){
		t.Errorf("Save Error: %s\n", err.Error())
	}else{
		t.Logf("%+v",content)
	}
}