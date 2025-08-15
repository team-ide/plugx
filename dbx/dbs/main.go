package main

import (
	"fmt"
	"os"
	"path"
	"plugin"
	"reflect"
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

	findDbPlugin, err := p.Lookup("DbPlugin")
	if err != nil {
		panic(err)
	}
	fmt.Println("findDbPlugin:", findDbPlugin)
	fmt.Println("findDbPlugin Type:", reflect.TypeOf(findDbPlugin))
	dbPlugin, ok := findDbPlugin.(**GreeterPlugin)
	fmt.Println("to plugin ok:", ok)
	fmt.Println("to plugin dbPlugin:", dbPlugin)
	open_, err := p.Lookup("Open")
	if err != nil {
		panic(err)
	}
	openFunc := open_.(func(Config) (string, error))
	close_, err := p.Lookup("Close")
	if err != nil {
		panic(err)
	}
	closeFunc := close_.(func(string) error)
	dbC := Config{
		DriverName: "sqlite3",
	}
	openKey, err := openFunc(dbC)
	if err != nil {
		panic(err)
	}
	fmt.Println(openKey)
	defer func() { _ = closeFunc(openKey) }()
}
