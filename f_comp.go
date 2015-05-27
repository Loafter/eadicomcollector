package main

import "errors"

import "os/exec"
import "log"

type FolderCompressor struct {
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
func (FolderCompressor) CompressFolder(Ctool string, Prm string, SrcF string, DstF string) error {
	if err := checkParam(Ctool, SrcF, DstF); err != nil {
		return err
	}

	if out, err := exec.Command(Ctool, Prm, DstF, SrcF).Output(); err != nil {
		log.Printf("error: %s\n", out)
		return err
	} else {
		log.Println(out)
	}
	return nil
}
