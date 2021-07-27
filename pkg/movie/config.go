package movie

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SomeSharedEnvVar string `envconfig:"SOME_SHARED_ENV_VAR"`
	SomeFxnEnvVar    int    `envconfig:"SOME_FXN_ENV_VAR"`

	ProcessMovieQueueUrl string `envconfig:"PROCESS_MOVIE_QUEUE_URL"`
}

func NewConfig() *Config {
	c := &Config{}
	envconfig.MustProcess("", c)
	return c
}
