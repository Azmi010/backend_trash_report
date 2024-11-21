package helper

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

func UploadFileToFirebase(file *multipart.FileHeader, bucket *storage.BucketHandle) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("unable to open file: %v", err)
	}
	defer src.Close()

	filename := fmt.Sprintf("reports/%d_%s", time.Now().UnixNano(), file.Filename)

	wc := bucket.Object(filename).NewWriter(context.Background())
	if _, err := io.Copy(wc, src); err != nil {
		return "", fmt.Errorf("unable to write file to bucket: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("unable to close writer: %v", err)
	}

	attrs, err := bucket.Object(filename).Attrs(context.Background())
	if err != nil {
		return "", fmt.Errorf("unable to get file attributes: %v", err)
	}

	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", attrs.Bucket, attrs.Name)
	return publicURL, nil
}