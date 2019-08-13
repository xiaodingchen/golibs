package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Read 读取配置文件
func Read(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(fmt.Sprintf("Fatal error config path: %s, err: %v", path, err))
	}
	ff := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.Index(f.Name(), ".") == 0 {
			continue
		}

		ext := f.Name()[(strings.LastIndex(f.Name(), ".") + 1):]
		if !isAllowConfigFileExt(ext) {
			continue
		}

		ff = append(ff, f.Name())
	}
	if len(ff) < 1 {
		panic(fmt.Sprintf("Fatal error config path: %s, err: %v", path, "no such config file."))
	}

	viper.AddConfigPath(path)
	for i, cf := range ff {
		// viper.SetConfigFile(cf)
		name := cf[0:strings.Index(cf, ".")]
		viper.SetConfigName(name)
		if i == 0 {
			err = viper.ReadInConfig()
		} else {
			err = viper.MergeInConfig()
		}
		if err != nil {
			panic(fmt.Sprintf("Fatal error config file:%s, err: %v", filepath.Join(path, cf), err))
		}
	}

	log.Printf("config files : %v", ff)
}

// Watch 监控配置文件
func Watch(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(fmt.Sprintf("Fatal error watch config path: %s, err: %v", path, err))
	}
	defer watcher.Close()
	err = watcher.Add(path)
	if err != nil {
		panic(fmt.Sprintf("Fatal error watch config path: %s, err: %v", path, err))
	}

	for {
		select {
		case event := <-watcher.Events:
			// log.Printf("config is change :%s \n", event.String())
			if (event.Op&fsnotify.Write == fsnotify.Write) ||
				(event.Op&fsnotify.Create == fsnotify.Create) ||
				(event.Op&fsnotify.Remove == fsnotify.Remove) ||
				(event.Op&fsnotify.Rename == fsnotify.Rename) {
				Read(path)
				log.Printf("config is change :%s \n", event.String())
			}
		case err = <-watcher.Errors:
			panic(fmt.Sprintf("Fatal error watch config path: %s, err: %v", path, err))
		}
	}
}

func isAllowConfigFileExt(ext string) bool {
	allowExt := []string{"toml", "json", "yml", "properties"}
	for _, v := range allowExt {
		if v == strings.ToLower(ext) {
			return true
		}
	}

	return false
}
