package main

import (
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
	getInterface, ok := sym.(func() dbx.Plugin)
	fmt.Println("to getInterface ok:", ok)
	fmt.Println("to getInterface res:", getInterface)
	dbPlugin := getInterface()

	conf := `
{
"DriverName":"sqlite3",
"DataSourceName":"./dbx_sqlite.db",
}
`
	openKey, err := dbPlugin.Open(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(openKey)
	defer func() { _ = dbPlugin.Close(openKey) }()

	db, err := dbPlugin.GetDB(openKey)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("create table tb_xxx(id int not null  primary key,name varchar(100) not null)")
	if err != nil {
		panic(err)
	}
}
