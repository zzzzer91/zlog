package zlog

// LoggerConfig 定义了日志的配置
type LoggerConfig struct {
	Level       int  `yaml:"level"` //  日志等级，error=2，warn=3，info=4，debug=5，见 `logrus.Level`
	TerminalOut bool `yaml:"terminalOut"`
	FileOut     bool `yaml:"fileOut"`
	FileConfig  struct {
		Path       string `yaml:"path"`
		MaxSize    int    `yaml:"maxSize"`
		MaxBackups int    `yaml:"maxBackups"`
		MaxAge     int    `yaml:"maxAge"`
	} `yaml:"fileConfig"`
}
