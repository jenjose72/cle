package main

import (
	"io"
	"net/http"
	"os"
)

func upload(w http.ResponseWriter, r *http.Request) {
	file, _, _ := r.FormFile("file")
	defer file.Close()

	out, _ := os.Create("uploaded.txt")
	defer out.Close()

	io.Copy(out, file)
	w.Write([]byte("File uploaded"))
}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}