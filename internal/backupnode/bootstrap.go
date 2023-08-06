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

func dockerCompose(args []string, env []string) {
	cmd := exec.Command("docker-compose", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Bootstrap(ctx context.Context, configPath string) {
	GenerateConfig(configPath)

	b, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	envs := os.Environ()
	envs = append(envs, fmt.Sprintf("MINIO_ROOT_USER=%s", cfg.MinioUser))
	envs = append(envs, fmt.Sprintf("MINIO_ROOT_PASSWORD=%s", cfg.MinioPassword))
	envs = append(envs, fmt.Sprintf("MONGO_ROOT_USERNAME=%s", cfg.MongoUser))
	envs = append(envs, fmt.Sprintf("MONGO_ROOT_PASSWORD=%s", cfg.MongoPassword))

	dockerCompose([]string{"up", "--build", "-d"}, envs)

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

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created bucket %s in minio\n", bucketName)
	}
}
