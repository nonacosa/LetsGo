package httpRetriever

import (
	"net/http"
	"net/http/httputil"
)

type HttpRetriever struct {
	Name string
}

func (httpRetriever HttpRetriever) Get(url string) string {
	 if url == "" {
	 	panic("url is empty")
	 }
	 if response, err := http.Get(url); err == nil {
	 	result,err := httputil.DumpResponse(response,true)

	 	if err != nil {
	 		panic(err)
		}

	 	return string(result)
	 } else {
	 	panic(err)
	 }
}

