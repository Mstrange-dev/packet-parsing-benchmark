package main

import (
	"bufio" // Added for buffering
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Packet struct {
	ID        uint32
	Timestamp int64
}

func main() {
	fmt.Println("Starting...")
	startTime := time.Now()

	file, err := os.Open("sample.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	
	reader := bufio.NewReader(file)
	buf := make([]byte, 12)
	packetCount := 0

	for {
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
			fmt.Printf("Read error: %v\n", err)
			break
		}

		_ = Packet{
			ID:        binary.LittleEndian.Uint32(buf[0:4]),
			Timestamp: int64(binary.LittleEndian.Uint64(buf[4:12])),
		}

		packetCount++
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Printf("\n--- Buffered Parsing Complete ---\n")
	fmt.Printf("Total Packets Read: %d\n", packetCount)
	fmt.Printf("Time Elapsed: %v\n", duration)
}
