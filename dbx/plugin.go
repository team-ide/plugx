package dbx

import (
	"context"
	"database/sql"
	"errors"
)

type Plugin interface {
	Open(c string) (openKey string, err error)
	Close(openKey string) (err error)
	IsDbPlugin()
}

// Database 定义统一的数据库接口
type Database interface {
	// 基础操作
	Connect(ctx context.Context, dsn string) error
	Close() error
	Ping(ctx context.Context) error

	// 查询操作
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row

	// 执行操作
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// 事务
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// 数据库类型
	Type() string
}

// DatabaseConstructor 插件构造函数类型
type DatabaseConstructor func() Database

var plugins = make(map[string]DatabaseConstructor)

// Register 注册数据库插件
func Register(name string, constructor DatabaseConstructor) {
	plugins[name] = constructor
}

// New 创建指定类型的数据库实例
func New(name string) (Database, error) {
	constructor, ok := plugins[name]
	if !ok {
		return nil, ErrPluginNotRegistered
	}
	return constructor(), nil
}

// SupportedDrivers 返回已注册的驱动列表
func SupportedDrivers() []string {
	drivers := make([]string, 0, len(plugins))
	for name := range plugins {
		drivers = append(drivers, name)
	}
	return drivers
}

var (
	ErrPluginNotRegistered = errors.New("plugin not registered")
)
