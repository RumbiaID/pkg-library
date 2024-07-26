package filevalidation

import (
	"bufio"
	"encoding/base64"
	"errors"
	"github.com/gabriel-vasile/mimetype"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)

func ValidateImage(base64String string, maxsize int64) (string, error) {
	// Decode the base64 string to bytes
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return "", err
	}

	// Get the size of the decoded data
	fileSize := int64(len(data))

	// Check if the file size exceeds the maximum allowed size
	if fileSize > maxsize {
		return "", errors.New("file size exceeds the maximum allowed size")
	}

	// Detect mimeType
	mimeType := mimetype.Detect(data)

	// Extension or mimeType checker, if not .png/.jpeg/.webp/.bmp return error
	if mimeType.String() != "image/png" && mimeType.String() != "image/jpeg" && mimeType.String() != "image/vnd.mozilla.apng" && mimeType.String() != "image/webp" && mimeType.String() != "image/bmp" {
		return "", errors.New("extension invalid: " + mimeType.String())
	}

	return mimeType.Extension(), nil
}

func UploadFile64(filePath string, filename string, base64String string) error {
	// Decode base64 string into a byte slice
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}

	// Create the file with the specified path
	dst, err := os.Create(filePath + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Write the decoded data to the file
	if _, err := dst.Write(data); err != nil {
		return err
	}

	return nil
}

func LoadImage64(filePath string, filename string) (string, error) {
	// Open file on disk.
	file, err := os.Open(filePath + filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(file)
	content, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded, nil
}

func ValidateImageReader(reader io.Reader) (string, error) {
	// Read the image data from the io.ReadCloser
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	// Check if the file size exceeds the maximum allowed size
	if len(data) > 5242880 {
		logrus.Printf("exceeding max file size (max byte %d)", 5)
		return "", errors.New("file size exceeds the maximum allowed size")
	}

	// Detect mimeType
	mimetype.SetLimit(0)
	mimeType := mimetype.Detect(data)

	// Extension or mimeType checker, if not .png/.jpeg/.webp/.bmp return error
	if mimeType.String() != "application/pdf" && mimeType.String() != "image/png" && mimeType.String() != "application/msword" && mimeType.String() != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" && mimeType.String() != "application/zip" {
		return "", errors.New("extension or mimeType invalid: " + mimeType.String())
	}

	return mimeType.Extension(), nil
}

func ValidateFile(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close() // Close the file when we're done

	// Get file info to check its size
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	maxSize := os.Getenv("FILE_MAX_SIZE")
	maxSizeInt, _ := strconv.Atoi(maxSize)
	// Check if the file size exceeds the maximum allowed size
	if fileInfo.Size() > int64(maxSizeInt) {
		logrus.Printf("Exceeding max file size (max bytes %d)", maxSizeInt)
		return "", errors.New("file size exceeds the maximum allowed size")
	}

	// Read the image data from the file
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Detect mimeType
	mimetype.SetLimit(0)
	mimeType := mimetype.Detect(data)

	// Extension or mimeType checker, if not .png/.jpeg/.webp/.bmp return error
	if mimeType.String() != "application/pdf" && mimeType.String() != "image/png" && mimeType.String() != "application/msword" && mimeType.String() != "application/vnd.openxmlformats-officedocument.wordprocessingml.document" && mimeType.String() != "application/zip" {
		return "", errors.New("extension or mimeType invalid: " + mimeType.String())
	}

	return mimeType.Extension(), nil
}
