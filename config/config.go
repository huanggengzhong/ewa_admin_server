package config

type Configuration struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	MySQL MySQL `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

//注意：配置结构体中 mapstructure 标签需对应 config.ymal 中的配置名称， viper 会根据标签 value 值把 config.yaml 的数据赋予给结构体
