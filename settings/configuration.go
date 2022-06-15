package settings

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Listen string `envconfig:"LISTEN" required:"true"`
}

func (c *Configuration) Load() error {
	return envconfig.Process("", c)
}
