package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	getSite()
}
func getSite() {
	resp, err := http.Get(os.Args[1])
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	outputFile, err := os.Create("output.html")
	if err != nil {
		log.Fatalln("Error occurred while creating new file:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error occurred while reading site:", err)
	}
	_, err = outputFile.WriteString(string(body))
	if err != nil {
		log.Println("Error occurred while writing to file:", err)
	}

}
