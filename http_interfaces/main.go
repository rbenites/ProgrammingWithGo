package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	//os.Stdout implements writer interface, resp.Body implements reader interface
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)

	//reader interface - bs= byte slice that "makes" a slice type byte with 99999 empty spaces
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
}
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("wrote ", len(bs), " bytes")
	return len(bs), nil
}
