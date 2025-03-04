package test

import (
	"fmt"
	"log"
	"os"

	"github.com/Arun-Kumar21/distributed-object-storage/internal/chunking"
	"github.com/Arun-Kumar21/distributed-object-storage/internal/storage"
)
func TestStorage(){
	testfile := "test_file.txt"
	err := os.WriteFile(testfile, []byte("This is test file to verify chunking."), 0644)

	if err != nil {
		log.Fatalf("Chunking failed: %v", err)
	}

	chunks, err := chunking.SplitFileIntoChunks(testfile)
	if err != nil {
		log.Fatalf("Chunking failed: %v", err)
	}


	// Storing all chunks
	for _, chunk := range chunks {
		err := storage.SaveChunk(chunk.Hash, chunk.Data)
		if err != nil {
			log.Fatalf("Failed to store chunk: %v", err)
		}
		fmt.Printf("Saved chunk: %s\n", chunk.Hash)	
	}
	
	// Retrieving chunks
	for _, chunk := range chunks {
		data, err := storage.RetrieveChunk(chunk.Hash)
		if err != nil {
			log.Fatalf("Failed to retrieve chunk: %v", err)
		}
		
		if string(data) == string(chunk.Data) {
			fmt.Printf("Chunk verified: %s ✅", chunk.Hash)
		} else {
			fmt.Printf("Chunk mismatch: %s ❌", chunk.Hash)
		}
	}

}
