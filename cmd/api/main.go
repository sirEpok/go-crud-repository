package main

import (
	"flag"
	"log"
	"testProj/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string //= "configs/api.toml"//Перенес в инит флаг
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "Запустить из командной строки с флагом path указав путь")
}

func main() {
	flag.Parse()
	log.Println("Работает")

	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Не смог считать конфиг файл, но запустится на стандартном порту:", err)
	}
	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}