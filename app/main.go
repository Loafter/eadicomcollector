package main

import (
	"log"
	"eadicomcollector/lib"
)

func main() {
	EaSrvCmp := eadicomcollector.EaFolderCompressorSrv{}
	log.Println("gui located add http://localhost:9980/index.html")
	EaSrvCmp.Start(9980)
}
