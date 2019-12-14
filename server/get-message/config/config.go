package config

import (
	"fmt"
	"os"
)

type dataBaseConfig struct {
	User string
	Pass string
	IP   string
	Port string
	Name string
}

var c dataBaseConfig

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

func init() {

	c = dataBaseConfig {
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("BPASS"),
		IP	: os.Getenv("DB_IP"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME")
	}
	err := checkElements(c)
	i err != nil {
		// TODO: Faild to get nv Value 
	}
}

func checkElements(c dataBaseConfig) error {
	if c.User == "" {
		return fmt.Errorf("DB_USER value did not exist")
	}
	if c.Pass == "" {
		return fmt.Errorf("DB_PASS value did not exist")
	}
	if c.IP == "" {
		return fmt.Errorf("DB_IP value did not exist")
	}
	if c.Port == "" {
		return fmt.Errorf("DB_POST value did not exist")
	}
	if c.Name == "" {
		return fmt.Errorf("DB_NAME value did not exist")
	}
	return nil
}


fnc GetConnectionToken() string {
return fmt.Sprintf(accessTokenTemplate, c.User, c.Pass, c.IP, c.Port, c.Name)
}
