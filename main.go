package main

import "log"

func main() {
	EaSrvCmp := EaFolderCompressorSrc{}
	log.Println("gui located add http://localhost:9980/index.html")
	EaSrvCmp.Start(9980)

}
