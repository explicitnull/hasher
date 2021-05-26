package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

const (
	salt = "C6qpl4nCgYhg08vTXaQs"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(bytes)))

	var total, matched int

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		total++

		open := salt + record[2] + "000"

		h := sha256.New()
		h.Write([]byte(open))
		hash := h.Sum(nil)

		hashStr := hex.EncodeToString(hash)
		if hashStr == record[0] {
			matched++
		}
	}

	log.Printf("tolal: %d, matched: %d", total, matched)
}
