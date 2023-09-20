package conf

import (
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {
	// init viper
	if _, err := os.Stat(".env"); os.IsExist(err) {
		err := viper.BindEnv("MYSQL_HOST")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_PORT")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_USER")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_PASSWORD")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_DBNAME")
		if err != nil {
			return
		}
		err = viper.BindEnv("GINS_PORT")
		if err != nil {
			return
		}
	} else {
		viper.SetConfigFile(".env")
		viper.SetDefault("MYSQL_HOST", "localhost")
		viper.SetDefault("MYSQL_PORT", "3306")
		viper.SetDefault("MYSQL_USER", "root")
		viper.SetDefault("MYSQL_PASSWORD", "password")
		viper.SetDefault("MYSQL_DBNAME", "db1")
		viper.SetDefault("GINS_PORT", "8080")
		if err := viper.ReadInConfig(); err != nil {
			panic("Error reading config file: " + err.Error())
		}
	}
}
