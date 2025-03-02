package chunking

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// 4MB chunk size
const ChunkSize = 4 * 1024 * 1024

type Chunk struct {
	Hash string
	Data []byte
}


func SplitFileIntoChunks (filepath string) ([]Chunk, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	var chunks []Chunk
	buffer := make([]byte, ChunkSize)

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF{
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		if n==0 {
			break
		}

		// Hash for chunks
		hash := sha256.Sum256(buffer[:n])
		hashStr := hex.EncodeToString(hash[:])

		chunks = append(chunks, Chunk{
			Hash: hashStr,
			Data: buffer[:n],
		})
	}
	return chunks, nil
}