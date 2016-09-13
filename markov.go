package main

import (
	"os"
	"encoding/gob"
)

func generate_response(input_text string) string {
	return input_text
}

func train(train_text string) {
	dataDict[train_text] = []string{"hi", "hi"}
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
	}
}

func import_file(fp string) {
	
}

