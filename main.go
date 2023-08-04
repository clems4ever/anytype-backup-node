package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/anyproto/any-sync-tools/any-sync-network/cmd"
	"gopkg.in/yaml.v2"
)

func readConfig(filePath string, cfgPtr any) {
	configFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	err = yaml.NewDecoder(configFile).Decode(cfgPtr)
	if err != nil {
		log.Fatal(err)
	}
}

func writeConfig(filePath string, cfg any) {
	configFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	err = yaml.NewEncoder(configFile).Encode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	configurationsDir := "configurations"
	coordinatorAddr := "any-coordinator:4830"
	fileNodeAddr := "any-filenode:4730"
	syncNodeAddr := "any-syncnode:4430"

	cmd.CreateConfig(configurationsDir, true)

	coordinatorConfigPath := filepath.Join(configurationsDir, "coordinator.yml")
	syncNodeConfigPath := filepath.Join(configurationsDir, "sync_1.yml")
	fileNodeConfigPath := filepath.Join(configurationsDir, "file_1.yml")
	heartConfigPath := filepath.Join(configurationsDir, "heart.yml")

	coordinatorConfig := cmd.CoordinatorNodeConfig{}
	readConfig(coordinatorConfigPath, &coordinatorConfig)
	syncNodeConfig := cmd.SyncNodeConfig{}
	readConfig(syncNodeConfigPath, &syncNodeConfig)
	fileNodeConfig := cmd.FileNodeConfig{}
	readConfig(fileNodeConfigPath, &fileNodeConfig)
	heartConfig := cmd.HeartConfig{}
	readConfig(heartConfigPath, &heartConfig)

	coordinatorConfig.Yamux.ListenAddrs[0] = coordinatorAddr
	coordinatorConfig.Network.Nodes[0].Addresses[0] = coordinatorAddr
	coordinatorConfig.Network.Nodes[1].Addresses[0] = syncNodeAddr
	coordinatorConfig.Network.Nodes[2].Addresses[0] = fileNodeAddr

	syncNodeConfig.Yamux.ListenAddrs[0] = syncNodeAddr
	syncNodeConfig.Network.Nodes[0].Addresses[0] = coordinatorAddr
	syncNodeConfig.Network.Nodes[1].Addresses[0] = syncNodeAddr
	syncNodeConfig.Network.Nodes[2].Addresses[0] = fileNodeAddr

	fileNodeConfig.Yamux.ListenAddrs[0] = fileNodeAddr
	fileNodeConfig.Network.Nodes[0].Addresses[0] = coordinatorAddr
	fileNodeConfig.Network.Nodes[1].Addresses[0] = syncNodeAddr
	fileNodeConfig.Network.Nodes[2].Addresses[0] = fileNodeAddr

	heartConfig.Nodes[0].Addresses[0] = coordinatorAddr
	heartConfig.Nodes[1].Addresses[0] = syncNodeAddr
	heartConfig.Nodes[2].Addresses[0] = fileNodeAddr

	writeConfig(coordinatorConfigPath, coordinatorConfig)
	writeConfig(syncNodeConfigPath, syncNodeConfig)
	writeConfig(fileNodeConfigPath, fileNodeConfig)
	writeConfig(heartConfigPath, heartConfig)

}
