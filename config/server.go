package config

type Config struct {
	Etcd        *Etcd `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	RefreshTime int   `mapstructure:"refreshTime" json:"refreshTime" yaml:"refreshTime"`
	Port        int   `mapstructure:"port" json:"port" yaml:"port"`
}

type Etcd struct {
	Endpoints   []string `json:"endpoints"`
	DialTimeout int      `json:"dialTimeout"`
}

func NewDefaultConf() *Config {
	return &Config{
		Etcd: &Etcd{
			Endpoints:   []string{},
			DialTimeout: 5,
		},
		RefreshTime: 5,
		Port:        8080,
	}
}

func LoadConfig(path string) (*Config, error) {

	return NewDefaultConf(), nil
}
