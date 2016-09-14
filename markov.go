package main

import (
	"encoding/gob"
	"math/rand"
	"os"
	"strings"
	"time"
)

var END = "@END@"

func generateMarkovResponse(inputText string) string {
	seed := strings.Fields(preprocessText(inputText))[0]

	if _, ok := DataDict[seed]; !ok { // If key is not in DataDict
		return "idk how to understand"
	}
	var currWord string
	var sentence string
	current := 0
	currWord = seed
	for {
		total := 0

		keys := []string{}
		cw := DataDict[currWord]
		for k, v := range cw {
			total += v
			keys = append(keys, k)
		}

		rand.Seed(time.Now().Unix())
		threshold := rand.Intn(total)

		for i := 1; i < total; i++ {
			if current > threshold {
				if keys[i-1] != END {
					currWord = keys[i-1]
					sentence += " " + currWord
				} else {
					return sentence
				}
			}

			current += DataDict[currWord][keys[i]]
		}
	}

	return sentence
}

func trainMessage(msg string) {
	words := strings.Fields(preprocessText(msg)) // Split by whitespace to get individual words
	for i, word := range words {
		if _, ok := DataDict[word]; !ok { // If key is not in DataDict
			DataDict[word] = make(map[string]int)
		}
		if i == len(words)-1 {
			DataDict[word][END] += 1
		} else {
			DataDict[word][words[i+1]] += 1
		}
		// DataDict[word][words[i + 1]]
	}
}

func loadDataset(path string) (DataMapType, error) {
	res := make(DataMapType)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return res, nil

	}

	dataFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dataFile.Close()

	dataDecoder := gob.NewDecoder(dataFile)

	err = dataDecoder.Decode(&res)
	return res, err
}

func saveDataset(path string) error {
	dataFile, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	dataEncoder := gob.NewEncoder(dataFile)
	err = dataEncoder.Encode(DataDict)
	if err != nil {
		return err
	}

	return nil
}

func importFile(fp string) {
	// TODO
}

func preprocessText(text string) string {
	// TODO: Add more here to clean up punctuation, etc.
	return strings.ToLower(text)
}

func postprocessText(text string) string {
	return text
}
