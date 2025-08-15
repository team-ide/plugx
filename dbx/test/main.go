package main

import (
	"encoding/json"
	"fmt"
	"github.com/team-ide/plugx/dbx"
	"os"
	"path"
	"plugin"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println("pwd:", pwd)
	f := path.Join(pwd, "../../plugins/dbx_sqlite.so")
	fmt.Println("plugin:", f)
	p, err := plugin.Open(f)
	if err != nil {
		panic(err)
	}

	sym, err := p.Lookup("GetInterface")
	if err != nil {
		panic(err)
	}
	dbPlugin := sym.(func() dbx.Plugin)()
	conf := map[string]any{
		"driverName":     "sqlite3",
		"dataSourceName": "./dbx_sqlite.db",
	}
	confB, _ := json.Marshal(conf)
	db, err := dbPlugin.Open(string(confB))
	if err != nil {
		panic(err)
	}
	defer func() { _ = db.Close() }()

	if err = db.Ping(); err != nil {
		panic(err)
	}
	_, err = db.Exec("create table tb_xxx(id int not null  primary key,name varchar(100) not null)")
	if err != nil {
		panic(err)
	}
	if _, err = db.Exec("insert into tb_xxx(id,name) values(?,?)", 1, "张三"); err != nil {
		panic(err)
	}
}
