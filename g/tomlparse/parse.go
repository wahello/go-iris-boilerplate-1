package tomlparse

import (
	"github.com/kataras/golog"
	"github.com/pelletier/go-toml"
)

func Config(name ...string) *toml.Tree {
	group := "app"
	if len(name) > 0 {
		group = name[0]
	}
	config, err := toml.LoadFile("./resource/config/"+group+".toml")
	if err != nil {
		golog.Fatal("TomlError err : "+err.Error())
		return nil
	}
	return config
}