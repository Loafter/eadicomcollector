package main

import (
	"log"
	"eadicomcollector/lib"
	"bufio"
	"os"
)

func main() {


	EaSrvCmp := eadicomcollector.EaFolderCompressorSrv{}
	EaSrvCmp.Spath=readAllPath()
	log.Println("gui located add http://localhost:9980/index.html")
	EaSrvCmp.Start(9980)
}
func readAllPath()(sc []string) {
	file, err := os.Open("scan.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sc=append(sc,scanner.Text())
	}
	return
}
