package config

import (
	"github.com/jinzhu/configor"
	log "github.com/inconshreveable/log15"
	"os"
)


var Config struct {
	Port 	uint 	`default:"7000" env:"LISTEN_PORT"`
	DB 	struct	{
		Name 		string 		`default:"qor_example"`
		Adapter		string		`default:"mysql"`
		User		string
		Password 	string 
		Host		string	`default:"localhost:3000"`
	}
}

func init() {
        env := os.Getenv("GOLANG_ENV"); 
        os.Setenv("CONFIGOR_ENV", os.Getenv("GOLANG_ENV"))
        if env == "" {
                env = "dev"
                os.Setenv("CONFIGOR_ENV", env)
        }
        log.Info("Loading Configuration For :", "env", env)
	if err := configor.Load(&Config, "config/config.yml"); err != nil {
		panic(err)
	}
}
