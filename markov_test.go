package main

import (
	"testing"
)

func TestImportFileNewline(t *testing.T) {
	ImportFile([]byte(`Hello\\nWorld`))
	DataDict.RLock()

	val, ok := DataDict.Map[[2]string{"hello", "\n"}]
	DataDict.RUnlock()
	if !ok {
		t.Fatal("Map missing newline with hello as prefix")
	}

	if len(val) != 1 {
		t.Fatal("Map value has wrong length")
	}

	if val[0] != "world" {
		t.Fatal("Map value isn't world")
	}
}