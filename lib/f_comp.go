package eadicomcollector

import "errors"

import "os/exec"
import (
	"log"
	"os"
)

type ZipCompressor struct {
	ctool string
	prmC  string
	prmT  string
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func (zc ZipCompressor) checkZTool() bool {
	_, err := exec.Command(zc.ctool).Output()
	return err == nil
}


func CreateCompressor() (zc *ZipCompressor,err error){
	zc=new(ZipCompressor)
	zc.ctool="7z.exe"
	zc.prmC="a"
	zc.prmT="-tzip"
	if !zc.checkZTool(){
		return nil,errors.New("error: cant' create compressor")
	}
	return zc,nil

}
func (zc ZipCompressor) CompressFolder(SrcF string, DstF string) error {
	if exists(DstF) {
		log.Println("info: archive already refresh")
		os.Remove(DstF)
	}
	if out, err := exec.Command(zc.ctool, zc.prmC, zc.prmT,DstF,SrcF ).Output(); err != nil {
		log.Printf("error: %s\n", out)
		return err
	} else {
		log.Printf("success: %s\n", out)
	}
	return nil
}
