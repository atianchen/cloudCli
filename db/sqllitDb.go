package db

import (
    "database/sql"
    "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqlLiteDb struct{
	 db *sqlx.DB
}

func (s *SqlLiteDb) Connect(cfg DbConfig){
	 db, _ := sqlx.Connect("sqlite3", (cfg.(FileDbConfig)).DbFile)//sql.Open("sqlite3", (cfg.(FileDbConfig)).DbFile)
  	if (db!=nil){
  		s.db = db
  	}
}


func (s *SqlLiteDb) Execute(sql  string,args ...any) (sql.Result,error){
	return s.db.Exec(sql,args...)
/*	if (args==nil){
		return s.db.Exec(sql)
	}else{
		stmt, err := s.db.Prepare(sql)
	    if (err!=nil){
	    	return nil,err
	    }
	    return stmt.Exec(args...)
	}*/
}


func (s *SqlLiteDb) Get(dest interface{},sql string ,args...any) error{
	return  s.db.Get(dest,sql,args...)
}


func (s *SqlLiteDb) Query(dest interface{},sql string,args ...any) error{
	return s.db.Select(dest,sql,args...)
}

/**
 * 获取原始链接
 */
func (s *SqlLiteDb) Raw() *sqlx.DB{
	return s.db
}

func (s *SqlLiteDb) Release(){
	if (s.db!=nil){
		s.db.Close()
	}
}
