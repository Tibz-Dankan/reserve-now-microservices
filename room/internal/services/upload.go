package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type Upload struct {
	StorageBucketBucket string `json:"storageBucket"`
	BaseURL             string `json:"baseURL"`
	DownloadURL         string `json:"url"`
	FilePath            string `json:"path"`
}

func (upload *Upload) initStorageBucket() (*storage.BucketHandle, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	currentDirPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fmt.Println("currentDir ===> ", currentDirPath)

	upload.BaseURL = "https://firebasestorage.googleapis.com/v0/b/"
	// storageBucket := os.Getenv("STORAGE_BUCKET")
	storageBucket := "reserve-now-677ca.appspot.com" //To be removed
	upload.StorageBucketBucket = storageBucket

	configStorage := &firebase.Config{
		StorageBucket: storageBucket,
	}
	opt := option.WithCredentialsFile(currentDirPath + "/internal/config/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), configStorage, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}

	return bucket, nil

}

// upload it to firebase storage cloud store
func (upload *Upload) Add(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	filePath := upload.FilePath
	if filePath == "" {
		return "", errors.New("no file path provided")
	}

	bucket, err := upload.initStorageBucket()
	if err != nil {
		return "", err
	}

	wc := bucket.Object(filePath).NewWriter(context.Background())
	_, err = io.Copy(wc, file)
	if err != nil {
		return "", err
	}

	err = wc.Close()
	if err != nil {
		return "", err
	}

	url, err := upload.getDownloadURL()
	if err != nil {
		return "", err
	}

	return url, nil
}

func (upload *Upload) transformFilePath() (string, error) {
	path := upload.FilePath

	if path == "" {
		return "", errors.New("no file path provided")
	}

	path = strings.ReplaceAll(path, "/", "%")
	path = strings.ReplaceAll(path, " ", "%")

	return path, nil
}

func (upload *Upload) getDownloadURL() (string, error) {

	transformedFilePath, err := upload.transformFilePath()
	if err != nil {
		return "", err
	}
	fmt.Println("transformed file path ===> ", transformedFilePath)

	FIREBASE_STORAGE_URL := upload.BaseURL + upload.StorageBucketBucket + "/0/" + transformedFilePath

	req, err := http.NewRequest(http.MethodGet, FIREBASE_STORAGE_URL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	type Response struct {
		Name               string `json:"name"`
		Bucket             string `json:"bucket"`
		Generation         string `json:"generation"`
		Metageneration     string `json:"metageneration"`
		ContentType        string `json:"contentType"`
		TimeCreated        string `json:"timeCreated"`
		Updated            string `json:"updated"`
		StorageClass       string `json:"storageClass"`
		Size               string `json:"size"`
		Md5Hash            string `json:"md5Hash"`
		ContentEncoding    string `json:"contentEncoding"`
		ContentDisposition string `json:"contentDisposition"`
		Crc32c             string `json:"crc32c"`
		Etag               string `json:"etag"`
		DownloadTokens     string `json:"downloadTokens"`
	}
	// type ResponseError struct {
	// 	StatusCode int `json:"statusCode"`
	// }

	if res.StatusCode != http.StatusOK {
		_, err := io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		return "", errors.New("request to firebase storage failed")
	}

	fmt.Printf("firebase-storage-service: status code: %d\n", res.StatusCode)
	rBody, _ := io.ReadAll(res.Body)
	fmt.Printf("firebase-storage-service: response body: %s\n", rBody)

	response := Response{}
	json.NewDecoder(strings.NewReader(string(rBody))).Decode(&response)

	fmt.Printf("response.DownloadTokens : %s\n", response.DownloadTokens)

	downloadURL := FIREBASE_STORAGE_URL + "?alt=media&token=" + response.DownloadTokens
	upload.DownloadURL = downloadURL

	fmt.Printf("downloadURL : %s\n", downloadURL)

	return downloadURL, nil
}

// func (upload *Upload) buildFileDownloadURL(downloadToken string) (string, error) {
// // {"url": "https://firebasestorage.googleapis.com/v0/b/owino-dd2f6.appspot.com/o/prod%2Fchat%2F1695688317408_Screenshot%20from%202023-09-25%2001-08-47.png?alt=media&token=3e6a8289-08b6-4605-a16a-bf5241d36037",
// // "path": "chat/1695688317408_Screenshot from 2023-09-25 01-08-47.png", "type": "image"}
// }
