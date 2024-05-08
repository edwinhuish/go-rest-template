package config

// Configuration for logging
type LoggerConfig struct {
	ConsoleEnabled bool   `json:"console_enabled" mapstructure:"console_enabled"`
	ConsoleLevel   string `json:"console_level" mapstructure:"console_level"`
	ConsoleJson    bool   `json:"console_json" mapstructure:"console_json"`

	FileEnabled bool   `json:"file_enabled" mapstructure:"file_enabled"`
	FileLevel   string `json:"file_level" mapstructure:"file_level"`
	FileJson    bool   `json:"file_json" mapstructure:"file_json"`

	// Directory to log to to when filelogging is enabled
	Directory string `json:"directory" mapstructure:"directory"`
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string `json:"filename" mapstructure:"filename"`
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int `json:"max_size" mapstructure:"max_size"`
	// MaxBackups the max number of rolled files to keep
	MaxBackups int `json:"max_backups" mapstructure:"max_backups"`
	// MaxAge the max age in days to keep a logfile
	MaxAge int `json:"max_age" mapstructure:"max_age"`
}
