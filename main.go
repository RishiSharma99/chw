package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/reujab/wallpaper"
)

var url string = "https://source.unsplash.com/featured/1360x768"

func makeReq(url string) []byte {
	// res, _ := http.Head(url)
	// maps := res.Header
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return data

}

func downloadImage(data []byte) string {
	log.Println("Downloading Image")

	home := os.Getenv("HOME")
	addrs := home + "/.image"
	f, err := os.Create(addrs)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	n2, err := f.Write(data)
	log.Println("Downloaded Image of Size :", n2)
	return addrs
}

func main() {
	wallpaper.SetFromFile(downloadImage(makeReq(url)))
}
