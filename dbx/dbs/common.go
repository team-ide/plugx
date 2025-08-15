package main

import (
	"database/sql"
	"encoding/json"
	"github.com/team-ide/plugx/dbx"
	"time"
)

type Config struct {
	DriverName     string `json:"driverName,omitempty"`
	DataSourceName string `json:"dataSourceName,omitempty"`
	// 单位 毫秒 设置连接可以重用的最长时间
	ConnMaxLifetime *int64 `json:"connMaxLifetime,omitempty"`
	// 单位 毫秒 设置连接可以空闲的最长时间
	ConnMaxIdleTime *int64 `json:"connMaxIdleTime,omitempty"`
	// 设置空闲状态下的最大连接数
	MaxIdleConns *int `json:"maxIdleConns,omitempty"`
	// 设置数据库的最大打开连接数
	MaxOpenConns *int `json:"maxOpenConns,omitempty"`
}

type GreeterPlugin struct{}

func (this_ *GreeterPlugin) IsDbPlugin() {}

func (this_ *GreeterPlugin) Open(opts string) (db *sql.DB, err error) {
	c := &Config{}
	err = json.Unmarshal([]byte(opts), c)
	if err != nil {
		return
	}
	db, err = sql.Open(c.DriverName, c.DataSourceName)
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
	return
}

func GetInterface() dbx.Plugin {
	return &GreeterPlugin{}
}
