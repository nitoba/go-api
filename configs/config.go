package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var Config *conf

type conf struct {
	// DB configuration
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	// Web server configuration
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`

	// JWT configuration
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth    *jwtauth.JWTAuth
}

func LoadConfig(path ...string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	if len(path) == 0 {
		println("Test environment")
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(path[0])
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}

	Config.TokenAuth = jwtauth.New("HS256", []byte(Config.JWTSecret), nil)

	return Config, err
}
