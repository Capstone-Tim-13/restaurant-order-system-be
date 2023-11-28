package config

import (
	"context"
	"fmt"
	"io"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func init() {
	// Inisialisasi SDK Firebase
	opt := option.WithCredentialsFile("helpers/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(fmt.Errorf("error initializing app: %v", err))
	}

	firebaseApp = app
}

func UploadImageToStorage(itemID string, imageFile io.Reader) error {
	ctx := context.Background()

	// Mengambil referensi ke Firebase Storage
	client, err := firebaseApp.Storage(ctx)
	if err != nil {
		return err
	}

	// Set object path di Firebase Storage
	objectPath := fmt.Sprintf("images/%s.jpg", itemID)

	// Membuat objek di Firebase Storage
	bucket, err := client.DefaultBucket()
	if err != nil {
		return err
	}

	wc := bucket.Object(objectPath).NewWriter(ctx)
	if _, err := io.Copy(wc, imageFile); err != nil {
		return err
	}
	defer wc.Close()

	return nil
}
