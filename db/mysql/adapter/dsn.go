package adapter

import (
	"fmt"
)

const (
	DefaultPort = "3306"
	DefaultHost = "localhost"
)

type Dsn struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func NewDsn(user string, password string, database string) *Dsn {
	return &Dsn{
		user:     user,
		password: password,
		database: database,
		host:     DefaultHost,
		port:     DefaultPort,
	}
}

func (d *Dsn) SetUser(user string) *Dsn {
	d.user = user

	return d
}

func (d *Dsn) SetPassword(password string) *Dsn {
	d.password = password

	return d
}

func (d *Dsn) SetDatabase(database string) *Dsn {
	d.database = database

	return d
}

func (d *Dsn) SetHost(host string) *Dsn {
	d.host = host

	return d
}

func (d *Dsn) SetPort(port string) *Dsn {
	d.port = port

	return d
}

func (d *Dsn) Dsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.user,
		d.password,
		d.host,
		d.port,
		d.database,
	)
}
