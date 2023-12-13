package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func HASH_CalculateChecksum(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	hash.Write(data)
	checksum := hex.EncodeToString(hash.Sum(nil))

	fmt.Printf("Checksum for %s is %s\n", filePath, checksum)
	return checksum, nil
}

func HASH_CalculateChecksumForDirectory(dirPath string) (string, error) {
	var dirChecksum string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		checkSum, err := HASH_CalculateChecksum(path)
		if err != nil {
			return err
		}

		dirChecksum += checkSum
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	hash := sha256.New()
	hash.Write([]byte(dirChecksum))
	return hex.EncodeToString(hash.Sum(nil)), nil
}
