package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
	"fmt"
	"os/user"
)

var (
	db = initDB()
)

func initDB() *sql.DB {
	db, err := sql.Open("mysql", readDbConf())
	checkErr(err)
	return db
}

type DbConf struct {
	UserName     string
	Passwd       string
	IpPort       string
	DatabaseName string
}

func readDbConf() (mysqlConfig string) {
	mysqlConfig = "userName:passwd@tcp(ipPort)/databaseName?charset=utf8"
	var dbConf DbConf

	u, err := user.Current()
	checkErr(err)

	f, err := os.Open(fmt.Sprintf("%s/.tinyUrl/db.toml", u.HomeDir))
	checkErr(err)

	if _, err = toml.DecodeReader(f, &dbConf); err != nil {
		checkErr(err)
	}

	mysqlConfig = strings.Replace(mysqlConfig, "userName", dbConf.UserName, -1)
	mysqlConfig = strings.Replace(mysqlConfig, "passwd", dbConf.Passwd, -1)
	mysqlConfig = strings.Replace(mysqlConfig, "ipPort", dbConf.IpPort, -1)
	mysqlConfig = strings.Replace(mysqlConfig, "databaseName", dbConf.DatabaseName, -1)
	return
}
