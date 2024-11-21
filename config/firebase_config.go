package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var FirebaseStorage *storage.Client

func InitializeFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase: %v", err)
	}
	FirebaseApp = app

	storageClient, err := storage.NewClient(context.Background(), opt)
	if err != nil {
		log.Fatalf("error initializing Google Cloud Storage: %v", err)
	}
	FirebaseStorage = storageClient
}

func GetStorageBucket() *storage.BucketHandle {
	bucketName := os.Getenv("FIREBASE_STORAGE_BUCKET")
	return FirebaseStorage.Bucket(bucketName)
}