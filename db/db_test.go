package db

import (
	"cloudCli/db"
	"testing"
)



func TestAdd(t *testing.T){
	dbHelper := db.BadgerDbHelper{}
	defer dbHelper.Release()
	dbHelper.Connect("d:/temp/db")
	err := dbHelper.Set("name","chenzhi",0)
	if (err!=nil){
		t.Errorf("Save Error: %s\n", err.Error())
	}else{
		t.Logf("Success")
	}
}

func TestGet(t *testing.T){
	dbHelper := db.BadgerDbHelper{}
	defer dbHelper.Release()
	dbHelper.Connect("d:/temp/db")
	content,err := dbHelper.Get("name")
	if (err!=nil){
		t.Errorf("Save Error: %s\n", err.Error())
	}else{
		t.Logf("%+v",content)
	}
}