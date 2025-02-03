package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConf
	Server ServerConf
	Log    LogConf
	DB     DBConf
}

type AppConf struct {
	Mode    string
	Name    string
	Version string
}

type ServerConf struct {
	Run  string
	Port string
}

type LogConf struct {
	Level    string
	SavePath string
	FileName string
	FileExt  string
}

type DBConf struct {
	Type     string
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

// New returns a new configuration that have been read
func New() *Config {
	return load()
}

// load will read the configuration use Viper
func load() *Config {
	var conf *Config
	vp := viper.New()

	// read in enviroment variables that match
	vp.AutomaticEnv()
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// read config file
	vp.AddConfigPath(".")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	// auto search and read config file
	if err := vp.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// set custom development environment variables
	// default environment variables is config.yml
	//
	// custom environment variables file name setting rule:
	// config.{ENVIRONMENT}.yml
	// eg. config.dev.yml | config.beta.yml
	//
	// set "-env" paramter to use the custom environment:
	// eg. go run main.go -env dev
	for i := 0; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-env":
			env := os.Args[i+1]
			vp.SetConfigName("config." + env)

			if err := vp.MergeInConfig(); err != nil {
				panic(fmt.Errorf("merge config error: %s", err))
			}
		}
	}

	vp.Unmarshal(&conf)
	return conf
}
