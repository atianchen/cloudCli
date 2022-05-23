package db

import (
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqlLiteDb struct{
	 db *sql.DB
}

func (s *SqlLiteDb) Connect(cfg DbConfig){

	 db, _ := sql.Open("sqlite3", (cfg.(FileDbConfig)).DbFile)
  	if (db!=nil){
  		s.db = db
  	}
}


func (s *SqlLiteDb) Execute(sql  string,args ...any) (sql.Result,error){
	if (args==nil){
		return s.db.Exec(sql)
	}else{
		stmt, err := s.db.Prepare(sql)
	    if (err!=nil){
	    	return nil,err
	    }
	    return stmt.Exec(args...)
	}
}



func (s *SqlLiteDb) Query(sql string,args ...any) (*sql.Rows,error){
	return s.db.Query(sql,args...)
}

func (s *SqlLiteDb) Release(){
	if (s.db!=nil){
		s.db.Close()
	}
}
