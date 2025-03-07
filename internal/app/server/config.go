package server

// Config 구조체
type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Mode string `mapstructure:"mode"`
	} `mapstructure:"server"`
	Database struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
	} `mapstructure:"database"`
}
