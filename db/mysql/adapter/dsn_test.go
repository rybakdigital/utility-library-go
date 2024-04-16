package adapter

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNewDsn(t *testing.T) {
	u := "john"
	p := "secret"
	d := "foo"
	dsn := NewDsn(u, p, d)
	assert.Equal(t, dsn.user, u)
	assert.Equal(t, dsn.password, p)
	assert.Equal(t, dsn.database, d)
	assert.Equal(t, dsn.host, DefaultHost)
	assert.Equal(t, dsn.port, DefaultPort)
}

func TestSetters(t *testing.T) {
	u := "john"
	nu := "marry"
	p := "secret"
	np := "hush-hush"
	d := "foo"
	nd := "bar"
	host := "mysql"
	port := "3000"
	dsn := NewDsn(u, p, d)
	dsn.SetUser(nu).SetPassword(np).SetDatabase(nd).SetHost(host).SetPort(port)
	assert.Equal(t, dsn.user, nu)
	assert.Equal(t, dsn.password, np)
	assert.Equal(t, dsn.database, nd)
	assert.Equal(t, dsn.host, host)
	assert.Equal(t, dsn.port, port)
}

func TestDsn(t *testing.T) {
	u := "john"
	p := "secret"
	d := "foo"
	dsn := NewDsn(u, p, d)
	expected := "john:secret@tcp(localhost:3306)/foo?charset=utf8mb4&parseTime=True&loc=Local"
	assert.Equal(t, dsn.Dsn(), expected)
}
