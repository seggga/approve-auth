package configs

import (
	"flag"
	"log"
	"os"

	"github.com/seggga/approve-auth/internal/entity"
	"gopkg.in/yaml.v3"
)

// Config represents configuration for the application
type Config struct {
	JWT   JWT                        `yaml:"jwt"`
	Mongo Mongo                      `yaml:"mongo"`
	Data  map[string]entity.UserOpts `yaml:"users"`
}

// JWT contains settings for JWT
type JWT struct {
	Secret string `yaml:"secret"`
}

// Mongo represents configuration data for establishing connection
type Mongo struct {
	DSN string `yaml:"connection-string"`
}

// Read parses yaml file to get application Config
func Read() *Config {

	path := flag.String("c", "./configs/config.yaml", "set path to config yaml-file")
	flag.Parse()

	// cfgPath := "./configs/users.yaml"
	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("cannot open %s config file: %v", *path, err)
	}
	defer f.Close()

	cfg := &Config{}
	d := yaml.NewDecoder(f)
	if err := d.Decode(cfg); err != nil {
		log.Fatalf("cannot parse %s to users struct: %v", *path, err)
	}
	return cfg
}
