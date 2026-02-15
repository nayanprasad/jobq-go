package appConfig

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port int      `yaml:"port"`
	DB   DBConfig `yaml:"db"`
}

type DBConfig struct {
	DSN string `yaml:"dsn,omitempty"`
}

// func DefaultAppConfig() *AppConfig {
// 	return &AppConfig{
// 		Server: ServerConfig{
// 			Port: 8080,
// 			DB: DBConfig{
// 				DSN: "default vaoue",
// 			},
// 		},
// 	}
// }
