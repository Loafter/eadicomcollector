package eadicomcollector

import "errors"

import "os/exec"
import (
	"log"
	"os"
)

type FolderCompressor struct {
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func checkParam(Ctool string, SrcF string, DstF string) error {
	if len(Ctool) == 0 {
		return errors.New("error: compress tool don't set")
	}
	if len(SrcF) == 0 {
		return errors.New("error: target folder don't set")
	}
	if len(DstF) == 0 {
		return errors.New("error: output don't set")
	}
	return nil
}
func (FolderCompressor) CompressFolder(Ctool string, PrmC string, PrmT string, SrcF string, DstF string) error {
	if exists(DstF){
		log.Println("info: archive already refresh")
		os.Remove(DstF)
	}

	if err := checkParam(Ctool, SrcF, DstF); err != nil {
		return err
	}

	if out, err := exec.Command(Ctool, PrmC, PrmT, DstF, SrcF).Output(); err != nil {
		log.Printf("error: %s\n", out)
		return err
	} else {
		log.Printf("success: %s\n", out)
	}
	return nil
}
