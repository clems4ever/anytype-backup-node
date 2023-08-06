package backupnode

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

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

func Bootstrap(configPath string) {
	GenerateConfig(configPath)

	dockerCompose("up", "--build", "-d")
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
	accessKeyID := "miniorootuser"
	secretAccessKey := "miniorootpassword"

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
		log.Printf("Successfully created %s\n", bucketName)
	}
}
