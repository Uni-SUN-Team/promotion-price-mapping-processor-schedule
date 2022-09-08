package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Service struct {
	Name string
	Path string
}

func New(name string, path string) *Service {
	return &Service{
		Name: name,
		Path: path,
	}
}

func (srv *Service) ConfigENV() error {
	viper.SetConfigName(srv.Name)
	viper.AddConfigPath(srv.Path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
