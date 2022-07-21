package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

type (
	// Config -.
	Config struct {
		APP     App     `yaml:"app"`
		HTTP    HTTP    `yaml:"http"`
		Redis   Redis   `yaml:"redis"`
		RPC     RPC     `yaml:"rpc"`
		Log     Log     `yaml:"log"`
		TopList TopList `yaml:"toplist,omitempty"`
		Mysql   Mysql   `yaml:"mysql"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name"    env:"APP_NAME"`
		Version string `yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `yaml:"port" default:"8080"`
	}

	// RPC -.
	RPC struct {
		Network string `yaml:"network" default:"tcp"`
		Host    string `yaml:"host" default:""`
		Port    int    `yaml:"port" default:"8081"`
	}

	// Log -.
	Log struct {
		Level string `yaml:"level" env:"LOG_LEVEL"`
	}

	// Redis -.
	Redis struct {
		Host     string `yaml:"host" env:"REDIS_HOST" env-require:"true"`
		Port     string `yaml:"port" env:"REDIS_PORT" env-require:"true"`
		Password string `yaml:"password" env:"REDIS_PASSWORD" env-require:"true"`
		DB       int    `yaml:"db" env:"REDIS_DB" env-require:"true"`
	}

	// Mysql -.
	Mysql struct {
		Host     string `yaml:"host" env:"MYSQL_HOST" env-require:"true"`
		Port     string `yaml:"port" env:"MYSQL_PORT" env-require:"true"`
		Password string `yaml:"password" env:"MYSQL_PASSWORD" env-require:"true"`
		DB       string    `yaml:"db" env:"MYSQL_DB" env-require:"true"`
		User       string    `yaml:"user" env:"MYSQL_USER" env-require:"true"`
	}

	// TopList -.
	TopList struct {
		Host   string `yaml:"host"`
		Key    string `yaml:"key" env:"TOPLIST_API_KEY"`
		Secret string `yaml:"secret" env:"TOPLIST_API_SECRET"`
	}
)

func NewConfig() (*Config, error) {
	var cfg Config

	if err := readFile("./config/config.yaml", &cfg); err != nil {
		return nil, err
	}

	readEnv(&cfg)

	return &cfg, nil
}

// ReadFile 文件读取
func readFile(fp string, cfg *Config) (err error) {
	var bs []byte
	if bs, err = os.ReadFile(fp); err != nil {
		return err
	} else if err = yaml.Unmarshal(bs, cfg); err != nil {
		return err
	}
	return nil
}

func readEnv(cfg *Config) {
	if appName := os.Getenv("APP_NAME"); appName != "" {
		cfg.APP.Name = appName
	}

	if appVersion := os.Getenv("APP_VERSION"); appVersion != "" {
		cfg.APP.Version = appVersion
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		cfg.Log.Level = logLevel
	}

	if host := os.Getenv("REDIS_HOST"); host != "" {
		cfg.Redis.Host = host
	}

	if port := os.Getenv("REDIS_PORT"); port != "" {
		cfg.Redis.Port = port
	}

	if pwd := os.Getenv("REDIS_PASSWORD"); pwd != "" {
		cfg.Redis.Password = pwd
	}

	if db, err := strconv.Atoi(os.Getenv("REDIS_DB")); err == nil && db > -1 {
		cfg.Redis.DB = db
	}

	if key := os.Getenv("TOPLIST_API_KEY"); key != "" {
		cfg.TopList.Key = key
	}

	if secret := os.Getenv("TOPLIST_API_SECRET"); secret != "" {
		cfg.TopList.Secret = secret
	}

	// mysql
	if mysqlHost := os.Getenv("MYSQL_HOST"); mysqlHost != "" {
		cfg.Mysql.Host = mysqlHost
	}

	if mysqlPort := os.Getenv("MYSQL_PORT"); mysqlPort != "" {
		cfg.Mysql.Port = mysqlPort
	}

	if mysqlPassWord := os.Getenv("MYSQL_PASSWORD"); mysqlPassWord != "" {
		cfg.Mysql.Password = mysqlPassWord
	}

	if mysqlDb := os.Getenv("MYSQL_DB"); mysqlDb != "" {
		cfg.Mysql.DB = mysqlDb
	}

	if mysqlUser := os.Getenv("MYSQL_USER"); mysqlUser != "" {
		cfg.Mysql.User = mysqlUser
	}

}
