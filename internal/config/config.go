package config

import "os"

type HTTPServer struct {
	Addr string
}

// env-default:"production"

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}


func MustLoad(){
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	
}