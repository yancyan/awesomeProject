package db

import (
	"database/sql"
	"fmt"
	_ "github.com/sijms/go-ora/v2"
	. "project/src/config"
	. "project/src/log"
)

func InitOracleConfig() {
	InitConfig("dev-f1")
	InitLog("app-")

	dbProps := Config.Db
	Log.Infof("the db config is %+v", dbProps)

	osqlInfo := fmt.Sprintf("oracle://%s:%s@%s:%d/%s", dbProps.Username, dbProps.Password,
		dbProps.Host, dbProps.Port, dbProps.Dbname)
	db, err := sql.Open("oracle", osqlInfo)
	if err != nil {
		Log.Fatalf("connect oracle db error: %s:", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//rows, err := db.Query("select to_char(sysdate,'yyyy-mm-dd hh24:mi:ss') AS name from dual")
	//rows, err := db.Query("select * from boss_account.account where id = 10054")
	//if err != nil {
	//	fmt.Println("exec query error:", err.Error())
	//}
	//defer rows.Close()
	//
	//var s sub
	//var id int
	//var name string
	//for rows.Next() {
	//
	//	rows.Scan(&id, &name)
	//	Log.Printf("fetch item: %+v \n", s)
	//	Log.Printf("fetch item: id is %d name is  %s \n", id, name)
	//}

	const sqlQuery = "select * from boss_account.account where id = 10054"

	rs, _ := DoQuery(db, sqlQuery)
	Log.Infoln(rs)
}

func TestOracle() {
	InitOracleConfig()

}

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
