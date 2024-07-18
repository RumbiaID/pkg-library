package filevalidation

import (
	"bufio"
	"encoding/base64"
	"errors"
	"github.com/gabriel-vasile/mimetype"
	"io"
	"os"
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
