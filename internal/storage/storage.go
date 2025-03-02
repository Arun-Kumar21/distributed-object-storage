package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Arun-Kumar21/distributed-object-storage/pkg/config"
	"github.com/Arun-Kumar21/distributed-object-storage/pkg/security"
)

const StorageDir = "storage"

// Save chunk data to storage dir 
func SaveChunk(hash string, data []byte) error {
	key := config.GetEncryptionKey()

	encryptedData, err := security.EncyptData(data, key)
	if err != nil {
		return fmt.Errorf("encryption failed: %w", err)
	}

	err = os.MkdirAll(StorageDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create storage directory: %w", err)
	}

	filepath := filepath.Join(StorageDir, hash)

	err = os.WriteFile(filepath, encryptedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write chunk: %w",  err)
	}

	return nil
}


// Retrieve Chunk data from storage dir
func RetrieveChunk(hash string) ([]byte, error) {
	key := config.GetEncryptionKey()

	filepath := filepath.Join(StorageDir, hash)

	encryptedData, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read chunk: %w", err)
	}

	data, err := security.DecryptData(encryptedData, key)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}
	return data, nil
}

