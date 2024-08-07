package config

type Database struct {
	Host     string
	Port     int
	Name     string
	UserName string
	Password string
}

var RDbSetting *Database
var WDbSetting *Database
