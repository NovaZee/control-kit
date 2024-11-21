package config

type CMD struct {
	EtcdPoints []string
	ExternalIp string
	Metrics    bool
	ReportTime int64
	PName      string
}

type Config struct {
	Etcd *Etcd `mapstructure:"etcd" json:"etcd" yaml:"etcd"`

	ExternalIp string `mapstructure:"externalIp" json:"externalIp" yaml:"externalIp"`
	Metrics    bool   `mapstructure:"metrics" json:"metrics" yaml:"metrics"`
	ReportTime int64  `mapstructure:"reportTime" json:"reportTime" yaml:"reportTime"`
	PName      string `mapstructure:"programName" json:"programName" yaml:"programName"`
}

type Etcd struct {
	Endpoints   []string `json:"endpoints"`
	DialTimeout int64    `json:"dialTimeout"`
}

func BuildConfig(cmd CMD) *Config {
	return &Config{
		Etcd: &Etcd{
			Endpoints:   cmd.EtcdPoints,
			DialTimeout: 5,
		},
		ExternalIp: cmd.ExternalIp,
		Metrics:    cmd.Metrics,
		ReportTime: cmd.ReportTime,
		PName:      cmd.PName,
	}
}
