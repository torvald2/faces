package config

import (
	"os"
	"reflect"
	"sync"
)

// TODO  - marshaller with validator
type appConfig struct {
	DBConnectionString string `env:"DB_CONNECTION_STRING"`
	TCP_Port           string `env:"APP_PORT"`
	LogLevel           string `env:"LOG_LEVEL"`
	Tolerance          string `env:"TOLERANCE"`
	RecognizionMethod  string `env:"RECOGNIZING_METHOD"`
	Shops              string `env:"ACTIVE_SHOPS"`
	Emails             string `env:"SEND_EMAILS"`
	IsDev              string `env:"IS_DEV"`
}

var conf appConfig
var once sync.Once

func GetConfil() *appConfig {
	once.Do(setUp)
	return &conf
}

func setUp() {
	v := reflect.ValueOf(&conf).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		varname, _ := f.Tag.Lookup("env")
		env, ok := os.LookupEnv(varname)
		if ok {
			v.Field(i).SetString(env)
		}

	}
}
