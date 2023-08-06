package backupnode

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/anyproto/any-sync-tools/any-sync-network/cmd"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HostIP    string `yaml:"host_ip"`
	ConfigDir string `yaml:"config_dir"`

	MinioUser     string `yaml:"minio_user"`
	MinioPassword string `yaml:"minio_password"`

	MongoUser     string `yaml:"mongo_user"`
	MongoPassword string `yaml:"mongo_password"`
}

func readConfig(filePath string, cfgPtr any) {
	configFile, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
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

func GenerateNetworkConfig(configFilePath string) {
	b, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	configurationsDir := cfg.ConfigDir

	err = os.RemoveAll(configurationsDir)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
	err = os.Mkdir(configurationsDir, 0700)
	if err != nil {
		log.Fatal(err)
	}

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

	coordinatorAddr := fmt.Sprintf("%s:4830", cfg.HostIP)
	fileNodeAddr := fmt.Sprintf("%s:4730", cfg.HostIP)
	syncNodeAddr := fmt.Sprintf("%s:4430", cfg.HostIP)
	coordinatorConfig.Yamux.ListenAddrs[0] = "0.0.0.0:4830"
	coordinatorConfig.Network.Nodes[0].Addresses[0] = coordinatorAddr
	coordinatorConfig.Network.Nodes[1].Addresses[0] = syncNodeAddr
	coordinatorConfig.Network.Nodes[2].Addresses[0] = fileNodeAddr
	coordinatorConfig.Mongo.Connect = fmt.Sprintf("mongodb://%s:%s@mongo:27017", cfg.MongoUser, cfg.MongoPassword)

	syncNodeConfig.Yamux.ListenAddrs[0] = "0.0.0.0:4430"
	syncNodeConfig.Network.Nodes[0].Addresses[0] = coordinatorAddr
	syncNodeConfig.Network.Nodes[1].Addresses[0] = syncNodeAddr
	syncNodeConfig.Network.Nodes[2].Addresses[0] = fileNodeAddr

	fileNodeConfig.Yamux.ListenAddrs[0] = "0.0.0.0:4730"
	fileNodeConfig.Network.Nodes[0].Addresses[0] = coordinatorAddr
	fileNodeConfig.Network.Nodes[1].Addresses[0] = syncNodeAddr
	fileNodeConfig.Network.Nodes[2].Addresses[0] = fileNodeAddr
	fileNodeConfig.S3Store.Endpoint = "http://minio:9000/"
	fileNodeConfig.S3Store.Credentials.AccessKey = cfg.MinioUser
	fileNodeConfig.S3Store.Credentials.SecretKey = cfg.MinioPassword
	fileNodeConfig.S3Store.ForcePathStyle = true
	fileNodeConfig.Redis.URL = "redis://redis:6379/?dial_timeout=3&db=1&read_timeout=6s&max_retries=2"

	heartConfig.Nodes[0].Addresses[0] = fmt.Sprintf("%s:4830", cfg.HostIP)
	heartConfig.Nodes[1].Addresses[0] = fmt.Sprintf("%s:4430", cfg.HostIP)
	heartConfig.Nodes[2].Addresses[0] = fmt.Sprintf("%s:4730", cfg.HostIP)

	writeConfig(coordinatorConfigPath, coordinatorConfig)
	writeConfig(syncNodeConfigPath, syncNodeConfig)
	writeConfig(fileNodeConfigPath, fileNodeConfig)
	writeConfig(heartConfigPath, heartConfig)
}
