package conf

import (
	"github.com/kelseyhightower/envconfig"
	"toufu3.jp/Idp/repository"
)

var Env = getEnv()

func getEnv() *repository.Env {
	var env repository.Env
	envconfig.Process("", &env)
	return &env
}