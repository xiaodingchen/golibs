package db

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var fconfig = func() *Config {
	host := "localhost"
	port := "3306"
	user := "root"
	password := "root"
	dbname := "test"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, dbname, charset)
	config := &Config{}
	config.Driver = "mysql"
	config.DSN = dsn

	return config
}

func Test_NewClient(t *testing.T) {
	config := fconfig()
	client, err := NewClient(config)
	if err != nil {
		t.Fatal("db client err", err)
	}

	if client == nil {
		t.Fatal("db client nil")
	}

	t.Log("db client pass")
}
