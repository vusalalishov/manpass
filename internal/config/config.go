package config

import (
	"github.com/spf13/viper"
	"sync"
)

type Config interface {
	Get(string) string
}

type viperConfig struct {
	v *viper.Viper
}

func (c *viperConfig) Get(key string) string {
	return c.v.GetString(key)
}

var (
	once sync.Once
	v *viper.Viper
)

func ProvideConfig() Config {
	once.Do(func() {
		v = viper.New()
		v.AutomaticEnv()
	})
	return &viperConfig{v: v}
}

func InjectConfig() Config {
	return ProvideConfig()
}
