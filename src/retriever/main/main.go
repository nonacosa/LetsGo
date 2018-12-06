package main

import (
	"fmt"
	"retriever"
	"retriever/httpRetriever"
)

func download(r retriever.Retriever) string {
	return r.Get("http://www.gitrue.com")
}

func main() {
	var r  retriever.Retriever
	r = httpRetriever.HttpRetriever{"apple"}

	fmt.Println(download(r))
}
