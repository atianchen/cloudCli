package db

import (
	"cloudCli/db"
	"testing"
	"fmt"
)

func TestDb(t *testing.T){
	cfg:= db.FileDbConfig{DbFile:"d:/work/temp/sql.db"}
	dbInstance:=&db.SqlLiteDb{}
	dbInstance.Connect(cfg)

/*	rs,err:=dbInstance.Execute(`
		CREATE TABLE IF NOT EXISTS  "userinfo" (
			    "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
			    "username" VARCHAR(64) NULL,
			    "departname" VARCHAR(64) NULL
			);
		`)
   if (err!=nil){
		fmt.Println(err)
	}*/
	rs,err :=dbInstance.Execute("insert into userinfo (username,departname) values (?,?)","jensen1","rd")
	if (err!=nil){
		fmt.Println(err)
	}else{
		fmt.Println(rs.LastInsertId())
	}
	
	row,err1:= dbInstance.Query("select uid,username from userinfo where username=?","chenzhi11")
	if (row!=nil){
		 fmt.Println(row.Next())
		for row.Next() {
			var uid int
			var username string
	        row.Scan(&uid, &username)
	    }
	    row.Close()
	}else{

		fmt.Println(err1)
	}
}