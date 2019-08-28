package xmysql

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/jinzhu/gorm"
	"github.com/xiaodingchen/golibs/db"
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
	config.Config = &db.Config{}
	config.Config.DSN = dsn
	config.Config.Driver = "mysql"

	return config
}

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

func Test_Client(t *testing.T) {
	config := fconfig()
	client, err := NewClient(config)
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Fatal("mysql client nil")
	}

	t.Log("name", config.Name)
	client2, err := NewClient(config)
	if !(client == client2) {
		t.Fatal("db manager fatal")
	}

	t.Log("mysql client pass")
}

func Test_Create_Table(t *testing.T) {
	config := fconfig()
	client, err := NewClient(config)
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Fatal("mysql client nil")
	}

	client.LogMode(true)

	err = client.CreateTable(User{}).Error
	if err != nil {
		t.Fatal("create table err", err)
	}

	t.Log("mysql client create table pass")
}

func Test_Insert(t *testing.T) {
	config := fconfig()
	client, err := NewClient(config)
	if err != nil {
		t.Fatal(err)
	}

	if client == nil {
		t.Fatal("mysql client nil")
	}

	client.LogMode(true)

	now := time.Now()
	no := uuid.New().String()
	email := fmt.Sprintf("%d@qq.com", time.Now().UnixNano())
	user := User{
		Name: "Jinzhu",
		Age: sql.NullInt64{
			Int64: 18,
		},
		Birthday:     &now,
		MemberNumber: &no,
		Email:        email,
	}

	err = client.Create(&user).Error
	if err != nil {
		t.Fatal("create table record err", err)
	}

	t.Log("mysql client create table record pass")
}

func Test_Select(t *testing.T) {
	multipleClients()
	client1, err := Select("master")
	if err != nil {
		t.Fatal("mysql client err", err)
	}
	client2, err := Select("slave1")
	if err != nil {
		t.Fatal("mysql client err", err)
	}
	client3, err := Select("slave2")
	if err != nil {
		t.Fatal("mysql client err", err)
	}

	if client1 == nil || client2 == nil || client3 == nil {
		t.Fatal("db manager fatal")
	}

	if client1 == client2 || client1 == client3 || client2 == client3 {
		t.Fatal("db manager fatal")
	}

	// time.Sleep(3 * time.Minute)

	t.Log("mysql client select pass")
}

func Test_Clear(t *testing.T) {
	multipleClients()
	err := DefaultManager.Clear()
	if err != nil {
		t.Fatal("mysql client clear err", err)
	}

	client1, err := Select("master")
	if err != nil {
		t.Log("mysql client err", err)
	}

	client2, err := Select("slave1")
	if err != nil {
		t.Log("mysql client err", err)
	}

	client3, err := Select("slave2")
	if err != nil {
		t.Log("mysql client err", err)
	}

	if client1 != nil || client2 != nil || client3 != nil {
		t.Fatal("db manager clear fatal")
	}

	t.Log("mysql client clear pass")
}

func multipleClients() {
	cfgs := []*Config{}
	config1 := fconfig()
	config1.Name = "master"
	cfgs = append(cfgs, config1)
	config2 := fconfig()
	config2.Name = "slave1"
	cfgs = append(cfgs, config2)
	config3 := fconfig()
	config3.Name = "slave2"
	cfgs = append(cfgs, config3)
	NewMultipleClients(cfgs)
}
