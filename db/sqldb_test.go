package db

import (
	"cloudCli/db"
	// "cloud"
	"testing"
	"fmt"
)
type UserInfo struct{
	Id int64
	UserName string
	DepartName string
}

func TestDb(t *testing.T){
	cfg:= db.FileDbConfig{DbFile:"d:/work/temp/sql.db"}
	dbInstance:=&db.SqlLiteDb{}
	dbInstance.Connect(cfg)
	defer dbInstance.Release()
	rs,err:=dbInstance.Execute(`
		CREATE TABLE IF NOT EXISTS  "userinfo" (
			    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
			    "username" VARCHAR(64) NULL,
			    "departname" VARCHAR(64) NULL
			);
		`)
   if (err!=nil){
		fmt.Println(err)
	}
	rs,err =dbInstance.Execute("insert into userinfo (username,departname) values (?,?)","jensen111","rd")
	if (err!=nil){
		fmt.Println(err)
	}else{
		fmt.Println(rs.LastInsertId())
	}
	
	docs :=[]UserInfo{}
	dbInstance.Query(&docs,"select * from userinfo where username=?","jensen111")
	fmt.Println(len(docs))
	for _,doc := range docs {
   	 fmt.Println(doc.UserName)    
   	 fmt.Println(doc.DepartName)  
	}
}