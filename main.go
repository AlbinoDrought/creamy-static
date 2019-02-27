package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("FILE_PATH")
	usingDefaultPath := len(path) == 0
	if usingDefaultPath {
		path = "./static"
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if usingDefaultPath {
			// create the folder
			os.Mkdir(path, 0777)
		} else {
			log.Fatalf("selected non-default path does not exist: %+v", path)
		}
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}

	watermark := os.Getenv("WATERMARK")

	fs := http.FileServer(http.Dir(path))

	// if there's no watermark set, don't use the watermarkedFileServer
	if len(watermark) > 0 {
		fs = watermarkedFileServer{
			Handler:   fs,
			watermark: watermark,
		}
	}

	http.Handle("/", fs)

	log.Printf("serving directory %v on %v", path, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
