package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const fmtstr = `<style>#updates::before{content:"See %d new updates";}</style>`

func main() {
	http.HandleFunc("/", index)
	log.Println("Serving on :8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	_, err := w.Write([]byte(`<html>
	<head><title>Noscript updates</title></head>
	<body><a href="/"><span id="updates"></span></a>`))
	if err != nil {
		return
	}
	flusher.Flush()
	updates := 0
	for {
		time.Sleep(time.Second)
		updates += 1
		_, err := w.Write([]byte(fmt.Sprintf(fmtstr, updates)))
		if err != nil {
			return
		}
		flusher.Flush()
	}
}
