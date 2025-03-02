package main 

import (
	"fmt"
	"log"
	"os"

	"github.com/Arun-Kumar21/distributed-object-storage/internal/chunking"
)

func main() {
	testfile := "test_file.txt"
	err := os.WriteFile(testfile, []byte("This is test file to verify chunking."), 0644)

	if err != nil {
		log.Fatalf("Chunking failed: %v", err)
	}

	chunks, err := chunking.SplitFileIntoChunks(testfile)
	if err != nil {
		log.Fatalf("Chunking failed: %v", err)
	}

	for i, chunk:= range chunks {
		fmt.Printf("Chunk %d: Hash: %s, Size: %d bytes\n", i+1, chunk.Hash, len(chunk.Data))
	}

	_= os.Remove(testfile)
}