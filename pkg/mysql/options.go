package mysql

import "fmt"

type dsn struct {
	host        string
	port        string
	password    string
	username    string
	database    string
	charset     string
	maxIdleCons int
	maxOpenCons int
}

type Option func(*dsn)

func Host(host string) Option {
	return func(c *dsn) {
		c.host = host
	}
}

func Port(port string) Option {
	return func(c *dsn) {
		c.port = port
	}
}

func UserName(username string) Option {
	return func(c *dsn) {
		c.username = username
	}
}

func Password(pwd string) Option {
	return func(c *dsn) {
		c.password = pwd
	}
}

func Database(db string) Option {
	return func(c *dsn) {
		c.database = db
	}
}
func Charset(charset string) Option {
	return func(c *dsn) {
		c.charset = charset
	}
}

func MaxIdleCons(maxIdleCons int) Option {
	return func(c *dsn) {
		c.maxIdleCons = maxIdleCons
	}
}

func MaxOpenCons(maxOpenCons int) Option {
	return func(c *dsn) {
		c.maxOpenCons = maxOpenCons
	}
}

func (c *dsn) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", c.username, c.password, c.host, c.port, c.database, c.charset)
}