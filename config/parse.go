package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func init() {
	var err error
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if !FileExists(configFileName) {
		configFileName = filepath.Join(rootDir, configFileName)
	}
	Cfg, err = ParseConfig(configFileName, rootDir)
	if err != nil {
		log.Fatal(err)
	}
}

func ParseConfig(cfgPath string, rootDir string) (*Configs, error) {
	fd, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	config := &Configs{}
	config.RuntimeParam.RootDir = rootDir
	err = yaml.NewDecoder(fd).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// FileExists FileExists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
