package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB * sql.DB
func InitMysql() (err error) {
	fmt.Println("InitMysql....")
	if SqlDB == nil {
		SqlDB, _ = sql.Open("mysql", "user:mhqa1234mhqa1234@tcp(fbq-engine-dev-mysql.io.netease.com:32381)/cy_test1")
		CreateTableWithUser()
	}
	return

}
//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        email VARCHAR(64),
        createtime INT(10)
        );`

	ModifyDB(sql)
}
//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := SqlDB.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}
// 查询
func QueryRowDB(sql string) *sql.Row{
	return SqlDB.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return SqlDB.Query(sql)
}