package config

type Configuration struct {
	App           App           `mapstructure:"app" json:"app" yaml:"app"`
	Log           Log           `mapstructure:"log" json:"log" yaml:"log"`
	Database      Database      `mapstructure:"database" json:"database" yaml:"database"`
	Redis         Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch" json:"elasticsearch" yaml:"elasticsearch"`
}
