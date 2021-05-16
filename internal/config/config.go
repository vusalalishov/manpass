// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"sync"
)

type Config interface {
	Get(key string) string
}

type viperConfig struct {
	v *viper.Viper
}

func (c *viperConfig) Get(key string) string {
	panic("Not implemented!")
}
var (
	once sync.Once
)
var v *viper.Viper

func provideConfig() Config {
	once.Do(func() {
		v = viper.New()
	})
	return &viperConfig{v: v}
}

func InjectConfig() Config {
	panic(wire.Build(provideConfig))
}