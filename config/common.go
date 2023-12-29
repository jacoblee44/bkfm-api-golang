package config

import (
	"fmt"
	"github.com/gravitational/configure"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Domain   string   `yaml:"domain"`
	Demo     bool     `yaml:"demo"`
	Database Database `yaml:"database"`
	Secret   string   `yaml:"secret"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Database string `yaml:"database"`
	Password string `yaml:"password"`
}

type Cloudflare struct {
	AccountId string           `yaml:"account_id"`
	Email     string           `yaml:"email"`
	Stream    CloudflareStream `yaml:"stream,flow"`
}

type Sentry struct {
	Dsn string `yaml:"dsn"`
}

type CloudflareStream struct {
	Token     string `yaml:"token"`
	Subdomain string `yaml:"subdomain"`
}

type Mail struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Sender   string `yaml:"sender"`
}

type Migration struct {
	Migrate   bool       `yaml:"migrate"`
	Auto      bool       `yaml:"auto"`
	Databases []Database `yaml:"databases,flow"`
}

func Init() Config {
	var config Config
	log.Info("Load config!")
	data, err := os.ReadFile("config.yml")
	err = configure.ParseYAML(data, &config)

	if err != nil {
		fmt.Println(err)
	}
	log.Info("Config loaded!")
	return config
}
