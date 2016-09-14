package main

import (
	"os"
	"encoding/gob"
	"strings"
	"fmt"
	"math/rand"
	"time"
)

var END = "@END@"

func generate_response(input_text string) string {
	seed := strings.Fields(preprocess_text(input_text))[0]

	if _, ok := dataDict[seed]; !ok { // If key is not in dataDict
		return "idk how to understand"
	}
	var currWord string
	var sentence string
	current := 0
	currWord = seed
	for true {
		total := 0

		keys := []string{}
		for k, v := range dataDict[currWord] {
			total += v
			keys = append(keys, k)
		}

		rand.Seed(time.Now().Unix())
		threshold := rand.Intn(total)

		for i := 1; i < total; i++ {
			if (current > threshold) {
				if keys[i - 1] != END {
					currWord = keys[i - 1]
					sentence += " " + currWord
				} else {
					return sentence;
				}
			}

	        current += dataDict[currWord][keys[i]]
		}
	}

	return sentence
}

func train(train_text string) {
	words := strings.Fields(preprocess_text(train_text)) // Split by whitespace to get individual words
	for i, word := range words {
		if _, ok := dataDict[word]; !ok { // If key is not in dataDict
			dataDict[word] = make(map[string]int)
		}
		if (i+1 == len(words)) {
			dataDict[word][END] += 1
		} else {
			dataDict[word][words[i + 1]] += 1
		}
		// dataDict[word][words[i + 1]] 
	}

	save_dataset()
}

func load_dataset(fp string) {
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		// If fp doesn't exist, create an empty file there and leave dataDict empty
		dataFile, _ := os.Create(fp)
		dataEncoder := gob.NewEncoder(dataFile)
	 	dataEncoder.Encode(dataDict)
		dataFile.Close()

	} else {
		dataFile, _ := os.Open(fp)
		dataDecoder := gob.NewDecoder(dataFile)
 		_ = dataDecoder.Decode(&dataDict)
 		dataFile.Close()
 		fmt.Println(dataDict)
	}
}

func save_dataset() {
	dataFile, _ := os.OpenFile(filePath, os.O_WRONLY, os.ModeAppend)
	dataEncoder := gob.NewEncoder(dataFile)
	dataEncoder.Encode(dataDict)
	dataFile.Close()

	fmt.Println("Saved dataset")
}

func import_file(fp string) {
	// TODO
}

func preprocess_text(text string) string {
	// TODO: Add more here to clean up punctuation, etc.
	return strings.ToLower(text) 
}

func postprocess_text(text string) string {
	return text
}
