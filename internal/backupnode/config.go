package backupnode

import (
	"log"
	"os"

	_ "embed"
)

//go:embed config.yml
var configFile string

func Init() {
	err := os.WriteFile("config.yml", []byte(configFile), 0600)
	if err != nil {
		log.Fatal(err)
	}
}
