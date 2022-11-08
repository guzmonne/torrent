package main

import (
	"log"
	"os"

	"github.com/guzmonne/torrent/pkg/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
