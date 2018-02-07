package main

import "log"
import "testing"
import "eadicomcollector/lib"
func TestFc(t *testing.T) {
	fc,err := eadicomcollector.CreateCompressor()
	if err!=nil{
		t.Error(err)
		return
	}

	if err := fc.CompressFolder( "C:\\Temp\\", "C:\\Temp\\temp.zip\\"); err != nil {
		t.Errorf("error: normal compress failed ")
		log.Println(err)
		return
	}
}
