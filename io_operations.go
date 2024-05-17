package main

import (
	"github.com/spf13/viper"
)

func (a *App) ReadConfig() map[string]any {
	config := viper.New()
	config.AddConfigPath(a.InitialConfigRead())
	config.SetConfigName("config")
	config.SetConfigType("toml")
	config.ReadInConfig()
	return config.AllSettings()
}

func (a *App) WriteOutputFormats(format string, state bool){
	config := viper.New()
	config.AddConfigPath(a.InitialConfigRead())
	config.SetConfigName("config")
	config.SetConfigType("toml")
	config.ReadInConfig()
	config.Set("saved_state." + format, state)
	config.WriteConfig()
}

func (a *App) WriteOutputQualities(format string, value int8){
	config := viper.New()
	config.AddConfigPath(a.InitialConfigRead())
	config.SetConfigName("config")
	config.SetConfigType("toml")
	config.ReadInConfig()
	config.Set("saved_state." + format, value)
	config.WriteConfig()
}

func (a *App) WriteOutputCodeOptions(propertie string, value bool){
	config := viper.New()
	config.AddConfigPath(a.InitialConfigRead())
	config.SetConfigName("config")
	config.SetConfigType("toml")
	config.ReadInConfig()
	config.Set("saved_state." + propertie, value)
	config.WriteConfig()
}

func (a *App) WriteIndentingAmount(value int8){
	config := viper.New()
	config.AddConfigPath(a.InitialConfigRead())
	config.SetConfigName("config")
	config.SetConfigType("toml")
	config.ReadInConfig()
	config.Set("saved_state.indenting_amount", value)
	config.WriteConfig()
}