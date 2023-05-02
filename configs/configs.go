package configs

import (
	"embed"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

//go:embed *
var configsFS embed.FS

var appFile = "app.yaml"

var app = &App{}

func init() {
	file, _ := configsFS.Open(appFile)
	v := viper.New()
	v.SetConfigType("yaml")
	_ = v.ReadConfig(file)

	dir, err := os.Getwd()
	if err == nil {
		appConfig := dir + string(os.PathSeparator) + "configs" + string(os.PathSeparator) + appFile
		v.SetConfigFile(appConfig)
		_ = v.MergeInConfig()
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
			_ = v.Unmarshal(app)
			log.Println("config change", app, app.Website.Meta)
		})
	}
	_ = v.Unmarshal(app)
}

func GetApp() *App {
	return app
}
