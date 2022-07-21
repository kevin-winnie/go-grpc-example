package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-grpc/config"
	"sync"
)

const (
	_defaultMaxOpenConn = 100
	_defaultMaxIdleConn = 2
	_defaultCharset     = "utf8mb4"
)

type server map[string]*sql.DB
type ServerMap struct {
	sync.RWMutex

	collect server
}

func Init() *ServerMap {
	return &ServerMap{
		collect: make(map[string]*sql.DB, 0),
	}
}

// New -.
func (serverMap *ServerMap) New(name string,opts config.Mysql) (*sql.DB, error) {

	serverMap.Lock()
	defer serverMap.Unlock()

	if server, ok := serverMap.collect[name]; ok {
		return server, nil
	}


	c := &dsn{
		maxIdleCons: _defaultMaxIdleConn,
		maxOpenCons: _defaultMaxOpenConn,
		charset:     _defaultCharset,
		host : opts.Host,
		port : opts.Port,
		password : opts.Password,
		database:opts.DB,
		username: opts.User,
	}


	dsn1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", c.username, c.password, c.host, c.port, c.database, c.charset)

	db, err := sql.Open("mysql", dsn1)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	serverMap.collect[name] = db

	return serverMap.collect[name], nil
}

func (serverMap *ServerMap) Check(name string) bool {
	if _, ok := serverMap.collect[name]; ok {
		return true
	}
	return false
}

// Use -.
func (serverMap *ServerMap) Use(name string) *sql.DB {
	return serverMap.collect[name]
}

// Close -.
func (serverMap *ServerMap) Close() error {
	for _, db := range serverMap.collect {
		if err := db.Close(); err != nil {
			return err
		}
	}
	return nil
}
