package configs

import "github.com/spf13/viper"

type Config struct {
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`
	DBName  string `mapstructure:"DB_NAME"`
	AppPort string `mapstructure:"APP_PORT"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
