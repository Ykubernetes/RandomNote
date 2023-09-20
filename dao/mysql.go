package dao

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

type Conf struct {
	DBmySQL     DBmySQL
	Application Application
	GinServer   GinServer
}

type DBmySQL struct {
	Username string
	Password string
	Host     string
	Dbname   string
	Port     int
}

type Application struct {
	Timeout   string
	Charset   string
	ParseTime bool `yaml:"parseTime"`
	Loc       string
}

type GinServer struct {
	Host string
	Port string
}

func GetConf() Conf {
	var conf Conf
	yamlFile, err := os.ReadFile("./conf/conf.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conf
}

var conf = GetConf()

func GetGinSerConf() (g GinServer) {
	return conf.GinServer
}

func InitMySQLByFile() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s&timeout=%s", conf.DBmySQL.Username, conf.DBmySQL.Password, conf.DBmySQL.Host, conf.DBmySQL.Port, conf.DBmySQL.Dbname, conf.Application.Charset, conf.Application.ParseTime, conf.Application.Loc, conf.Application.Timeout)
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func InitMySQLByViper() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString("MYSQL_USER"), viper.GetString("MYSQL_PASSWORD"), viper.GetString("MYSQL_HOST"), viper.GetString("MYSQL_PORT"), viper.GetString("MYSQL_DBNAME"))
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
