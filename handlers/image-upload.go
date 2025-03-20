package handlers

import (
	"fmt"
	"mime/multipart"
	"os"

	storage_go "github.com/supabase-community/storage-go"
)

var storageClient *storage_go.Client
var bucketName string = "images" // Name of Supabase storage bucket

func InitSupabase() {
	storageClient = storage_go.NewClient("https://uzxsucssubyqooijxfni.supabase.co/storage/v1", os.Getenv("SUPABASE_SERVICE_ROLE_KEY"), nil)

	// url := os.Getenv("SUPABASE_URL")
	// key := os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

	// client, err := supabase.NewClient(url, key, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to initialize Supabase: %v", err)
	// }

	// sbClient = client
}

// UploadRestaurantImage uploads an image to Supabase Storage
func UploadRestaurantImage(file multipart.File, fileName string) (string, error) {
	filePath := fmt.Sprintf("restaurant/%s", fileName) // Path inside the bucket

	upsert := true            // Will overwrite file if one w/ same name already exists (instead of 409 error)
	imageType := "image/jpeg" // TODO: support for at least .png
	uploadOptions := storage_go.FileOptions{Upsert: &upsert, ContentType: &imageType}
	_, err := storageClient.UploadFile(bucketName, filePath, file, uploadOptions) // not using response for now
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	// Get the public URL of the uploaded image
	publicURL := storageClient.GetPublicUrl(bucketName, filePath)
	//publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", os.Getenv("SUPABASE_URL"), bucketName, filePath)

	return publicURL.SignedURL, nil

	// bucketName := "images"                             // Name of Supabase storage bucket
	// filePath := fmt.Sprintf("restaurant/%s", fileName) // Path inside the bucket

	// // Read file into a buffer
	// var fileBuffer bytes.Buffer
	// _, err := io.Copy(&fileBuffer, file)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to read file: %w", err)
	// }

	// // Upload file to Supabase Storage
	// _, err = sbClient.From(bucketName).Upload(context.Background(), filePath, &fileBuffer, supabase.FileOptions{
	// 	ContentType: "image/png",
	// 	Upsert:      true, // Overwrite existing file if exists
	// })
	// if err != nil {
	// 	return "", fmt.Errorf("failed to upload file: %w", err)
	// }

	// // Get the public URL of the uploaded image
	// publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", os.Getenv("SUPABASE_URL"), bucketName, filePath)

	// return publicURL, nil
}

func DeleteRestaurantImage(name []string) (string, error) {
	name[0] = fmt.Sprintf("restaurant/%s", name[0]) // prepend path inside the bucket
	fmt.Println("Paths to delete:", name)

	resp, err := storageClient.RemoveFile(bucketName, name)
	println(resp[0].Key, resp[0].Error, resp[0].Message)
	if err != nil {
		return "", fmt.Errorf("failed to delete image(s): %w", err)
	}

	// Return how many images we deleted
	return fmt.Sprintf("Deleted %d image(s)", len(name)), nil
}
