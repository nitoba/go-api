package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var (
	config *conf
	logger *Logger
)

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
	logger := GetLogger("configs")
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	if len(path) == 0 {
		logger.Warn("Running in test environment...")
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(path[0])
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Error to reading configs: %v", err)
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Errorf("Error to reading configs: %v", err)
		panic(err)
	}

	config.TokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	return config, err
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetConfig() *conf {
	return config
}
