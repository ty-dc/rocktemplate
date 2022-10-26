package types

type ConfigmapConfig struct {
	EnableIPv4 bool `yaml:"enableIPv4"`
	EnableIPv6 bool `yaml:"enableIPv6"`
}
