package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	go_ora "github.com/sijms/go-ora/v2"
	// init 自动注册驱动，与gorm驱动注册冲突，先注释
	_ "github.com/sijms/go-ora/v2"
	. "project/src/config"
	. "project/src/log"
)

var DbClient *sql.DB

//
func InitOracleConfig() {

	db, err := sql.Open("oracle", GetDbUrl())
	if err != nil {
		Log.Fatalf("connect oracle db error: %s:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DbClient = db
}

func GetDbUrl() string {
	dbProps := Config.Db
	Log.Infof("the db config is %+v", dbProps)

	// conn, err := go_ora.NewConnection("oracle://user:pass@server/service_name")

	urlOptions := map[string]string{
		//"TRACE FILE": "trace.log",
		//"SERVER": "server2, server3",
		//"PREFETCH_ROWS": "500",
		//"SSL": "enable",
		//"SSL Verify": "false",
	}

	buildUrl := go_ora.BuildUrl(dbProps.Host, int(dbProps.Port), dbProps.Dbname, dbProps.Username, dbProps.Password, urlOptions)
	Log.Infoln("BuildUrl => ", buildUrl)

	return buildUrl

	//url := fmt.Sprintf("oracle://%s:%s@%s:%d/%s", dbProps.Username, dbProps.Password,
	//	dbProps.Host, dbProps.Port, dbProps.Dbname)
	//
	//Log.Infoln("fmt.Sprintf URL => ", url)
	//return url

}

func TestOracle() {
	InitConfig("dev-f1")
	InitLog("app-")

	//InitOracleConfig()
	//
	////rows, err := db.Query("select to_char(sysdate,'yyyy-mm-dd hh24:mi:ss') AS name from dual")
	//const sqlQuery = "select * from boss_account.account where id = 10054"
	//
	//rs, _ := DoQuery(DbClient, sqlQuery)
	//Log.Infoln(rs)

	//stmt, _ := DbClient.Prepare("select company_id, code from boss_account.account where id = :1")
	//// check for error
	//defer stmt.Close()
	//
	//rows, _ := stmt.Query(10054)
	//defer rows.Close()
	//
	//for rows.Next() {
	//	// define vars
	//	var code string
	//	var companyId int
	//	err := rows.Scan(&companyId, &code)
	//	if err != nil {
	//		panic(err)
	//	}
	//	Log.Infoln("=============> ", companyId, code)
	//	// check for error
	//}

	//
	conn, err := go_ora.NewConnection(GetDbUrl())
	if err != nil {
		panic(err)
	}
	err = conn.Open()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	stmt := go_ora.NewStmt("select company_id, code from boss_account.account where id = :id", conn)
	defer stmt.Close()
	stmt.AddParam("id", 10054, 40, go_ora.InOut /* or go_ora.Output*/)
	//rows, _ := stmt.Query(nil)
	//defer rows.Close()
	//
	//columns := rows.Columns()
	//
	//values := make([]driver.Value, len(columns))
	//
	//for {
	//	err = rows.Next(values)
	//	if err != nil {
	//		break
	//	}
	//	Record(columns, values)
	//}

	// to struct.
	//rows, err := stmt.Query_(nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//db := struct {
	//	CompanyId string `db:"name:company_id"`
	//	Code      string `db:"name:code"`
	//}{}
	//
	//for rows.Next_() {
	//	err = rows.Scan(&db)
	//	if err != nil {
	//		panic(err)
	//	}
	//	Log.Infoln("record => ", db)
	//}

	//	 to cursor

}

func Record(columns []string, values []driver.Value) {
	for i, c := range values {
		fmt.Printf("%-25s: %v\n", columns[i], c)
	}
	fmt.Println()
}

func cursor() {
	//var cursor go_ora.RefCursor
	//
	//conn, _ := db.Conn(context.Background())
	//conn.exe
	//_, err = conn.Exec(cmdText, sql.Out{Dest: &cursor})
	//
	//defer cursor.Close()
	//rows, _ := cursor.Query()
	//
	//// check for error
	//
	//var (
	//	var_1 int64
	//	var_2 string
	//)
	//for rows.Next_() {
	//	_ = rows.Scan(&var_1, &var_2)
	//	// check for error
	//	fmt.Println(var_1, var_2)
	//}
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
