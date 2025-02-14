package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config โครงสร้างของ Config
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

var AppConfig Config

// LoadConfig โหลดค่า config จากไฟล์ YAML
func LoadConfig() {
	file, err := os.Open("configs/config.yaml") // อ้างอิง path ให้ถูกต้อง
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
