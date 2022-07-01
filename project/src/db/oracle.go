package db

import (
	"database/sql"
	"fmt"
	//_ "github.com/sijms/go-ora/v2"
	. "project/src/config"
	. "project/src/log"
)

//var DbClient *sql.DB
//
//func InitOracleConfig() {
//	InitConfig("dev-f1")
//	InitLog("app-")
//
//	db, err := sql.Open("oracle", GetDbUrl())
//	if err != nil {
//		Log.Fatalf("connect oracle db error: %s:", err.Error())
//	}
//	err = db.Ping()
//	if err != nil {
//		panic(err)
//	}
//
//	DbClient = db
//}

func GetDbUrl() string {
	dbProps := Config.Db
	Log.Infof("the db config is %+v", dbProps)

	return fmt.Sprintf("oracle://%s:%s@%s:%d/%s", dbProps.Username, dbProps.Password,
		dbProps.Host, dbProps.Port, dbProps.Dbname)
}

//func TestOracle() {
//	InitOracleConfig()
//
//	//rows, err := db.Query("select to_char(sysdate,'yyyy-mm-dd hh24:mi:ss') AS name from dual")
//	const sqlQuery = "select * from boss_account.account where id = 10054"
//
//	rs, _ := DoQuery(DbClient, sqlQuery)
//	Log.Infoln(rs)
//}

// DoQuery 查询结果集合转到map
func DoQuery(db *sql.DB, sqlInfo string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(sqlInfo, args...)
	if err != nil {
		return nil, err
	}
	columns, _ := rows.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength) //临时存储每行数据
	for index, _ := range cache {              //为每一列初始化一个指针
		var a interface{}
		cache[index] = &a
	}
	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		_ = rows.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //取实际类型
		}
		list = append(list, item)
	}
	_ = rows.Close()
	return list, nil
}
