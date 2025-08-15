package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/team-ide/plugx/dbx"
	"time"
)

type Config struct {
	DriverName     string
	DataSourceName string
	// 单位 毫秒 设置连接可以重用的最长时间
	ConnMaxLifetime *int64
	// 单位 毫秒 设置连接可以空闲的最长时间
	ConnMaxIdleTime *int64
	// 设置空闲状态下的最大连接数
	MaxIdleConns *int
	// 设置数据库的最大打开连接数
	MaxOpenConns *int
}

var (
	openConnector = map[string]*DbConnector{}
)

type DbConnector struct {
	Config *Config
	*sql.DB
}

type GreeterPlugin struct{}

func (this_ *GreeterPlugin) IsDbPlugin() {}

func (this_ *GreeterPlugin) Open(opts string) (openKey string, err error) {
	c := &Config{}
	err = json.Unmarshal([]byte(opts), c)
	if err != nil {
		return
	}
	db, err := sql.Open(c.DriverName, c.DataSourceName)
	if err != nil {
		return
	}
	if c.ConnMaxLifetime != nil {
		db.SetConnMaxLifetime(time.Millisecond * time.Duration(*c.ConnMaxLifetime))
	}
	if c.ConnMaxIdleTime != nil {
		db.SetConnMaxIdleTime(time.Millisecond * time.Duration(*c.ConnMaxIdleTime))
	}
	if c.MaxIdleConns != nil {
		db.SetMaxIdleConns(*c.MaxIdleConns)
	}
	if c.MaxOpenConns != nil {
		db.SetMaxOpenConns(*c.MaxOpenConns)
	}
	openKey = fmt.Sprintf("%s_%d", c.DriverName, time.Now().Unix())
	find := openConnector[openKey]
	if find != nil {
		err = errors.New("连接 Key " + openKey + " 已存在")
		return
	}
	find = &DbConnector{}
	find.Config = c
	find.DB = db
	openConnector[openKey] = find
	return
}

func (this_ *GreeterPlugin) Close(openKey string) (err error) {
	find := openConnector[openKey]
	if find == nil {
		return
	}
	delete(openConnector, openKey)
	err = find.DB.Close()
	return
}

func GetInterface() dbx.Plugin {
	return &GreeterPlugin{}
}
