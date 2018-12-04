package app

import (
	"encoding/json"
	"flag"
	"os"

	log "github.com/panjiang/golog"
)

// Config Global static config
var Config *config

type config struct {
	SiteName string       `json:"sitename"`
	Release  bool         `json:"release"`
	Log      *log.Config  `json:"log"`
	Bind     string       `json:"bind"`
	BaseURL  string       `json:"baseUrl"`
	Redis    *RedisConfig `json:"redis"`
	Mysql    *MysqlConfig `json:"mysql"`
}

func (c *config) parse(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		return err
	}

	log.Infof("Log config (%s)", c.Log.DebugString())
	return log.ParseConfig(c.Log)
}

// ParseConfig parse the config with json file from args
// (default 'config.json')
func ParseConfig() error {
	filename := "config.json"
	if len(flag.Args()) > 0 {
		filename = flag.Args()[0]
	}

	log.Infof("Parsing config file '%s'", filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return err
	}

	Config = new(config)
	err := Config.parse(filename)
	if err != nil {
		return err
	}
	return nil
}

// ParseTo 解析配置到实例
func ParseTo(filename string, obj interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
