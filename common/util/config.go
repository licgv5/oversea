package common.util
import (
    "github.com/BurntSushi/toml"
    "strings"
    "fmt"
)

type tomlConfig struct {
    Servers server
}

type server struct {
    Port string
}

var (
    cfg * tomlConfig
)

func Config(fileName string) *tomlConfig {
    path, _ := os.Getwd()
    configPath := path + fileName
    fmt.Printf("parse toml file once. filePath: %s\n", filePath)
    if _ , err := toml.DecodeFile(configPath, &cfg); err != nil {
        return nil
    }
    return cfg
}
