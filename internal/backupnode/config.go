package backupnode

import (
	"log"
	"os"

	_ "embed"
)

//go:embed config.yml
var configFile string

//go:embed Dockerfile
var dockerfile string

//go:embed docker-compose.yml
var dockerComposeFile string

func dumpFile(path string, content string) {
	err := os.WriteFile(path, []byte(dockerfile), 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func Init() {
	err := os.WriteFile("config.yml", []byte(configFile), 0600)
	if err != nil {
		log.Fatal(err)
	}

	dumpFile("Dockerfile", dockerfile)
	dumpFile("docker-compose.yml", dockerComposeFile)
}
