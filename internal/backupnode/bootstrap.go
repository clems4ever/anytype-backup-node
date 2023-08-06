package backupnode

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gopkg.in/yaml.v3"
)

func dockerCompose(args ...string) {
	cmd := exec.Command("docker-compose", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateConfig(configPath string) {
	GenerateNetworkConfig(configPath)

	b, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	envFileContent := ""
	envFileContent += fmt.Sprintf("MINIO_ROOT_USER=%s", cfg.MinioUser) + "\n"
	envFileContent += fmt.Sprintf("MINIO_ROOT_PASSWORD=%s", cfg.MinioPassword) + "\n"
	envFileContent += fmt.Sprintf("MONGO_INITDB_ROOT_USERNAME=%s", cfg.MongoUser) + "\n"
	envFileContent += fmt.Sprintf("MONGO_INITDB_ROOT_PASSWORD=%s", cfg.MongoPassword) + "\n"
	dumpFile("backupnode.env", envFileContent)
}

func Bootstrap(ctx context.Context, configPath string) {
	GenerateConfig(configPath)

	dockerCompose("up", "--build", "-d")

	fmt.Println("Waiting for the infrastructure to boot...")
	time.Sleep(time.Second)

	Initialize(ctx, configPath)
}

func Initialize(ctx context.Context, configPath string) {
	b, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	endpoint := fmt.Sprintf("%s:9000", cfg.HostIP)
	accessKeyID := cfg.MinioUser
	secretAccessKey := cfg.MinioPassword

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := "any-sync-files"
	location := "eu-central-1"

	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	fmt.Println("Trying to create minio bucket...")
	err = minioClient.MakeBucket(newCtx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(newCtx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created bucket %s in minio\n", bucketName)
	}
}
