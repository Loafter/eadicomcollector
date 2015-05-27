package main

import "log"
import "testing"

func TestFc(t *testing.T) {
	fc := FolderCompressor{}
	if err := fc.CompressFolder("", "", "", ""); err == nil {
		t.Errorf("error: wrong param test failed ")
	}

	if err := fc.CompressFolder("7z.exe", "a -tzip", "\"C:\\Temp\\\"", "\"C:\\Temp\\temp.zip\""); err != nil {
		t.Errorf("error: normal compress failed ")
		log.Println(err)
	}
}
