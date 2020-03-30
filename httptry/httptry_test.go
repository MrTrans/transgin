package httptry_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	http.HandleFunc("/", func(respw http.ResponseWriter, rsq *http.Request) {
		respw.Write([]byte("hello trans!"))
	})
	fmt.Printf("hello http is listening...")
	if err := http.ListenAndServe(":8086", nil); err != nil {
		panic(err)
	}
}

func TestDir(t *testing.T) {
	http.Handle("/chloe", http.NotFoundHandler())
	http.Handle("/miji", http.RedirectHandler("www.mijisou.com", 301))
	if err := http.ListenAndServe(":8086", nil); err != nil {
		panic(err)
	}
}

func TestMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/chloe", func(respw http.ResponseWriter, rsq *http.Request) {
		io.WriteString(respw, "hello chloe!")
	})
	mux.HandleFunc("/trans", func(respw http.ResponseWriter, rsq *http.Request) {
		io.WriteString(respw, rsq.URL.Path)
	})
	if err := http.ListenAndServe(":8086", mux); err != nil {
		panic(err)
	}
}

func TestGet(t *testing.T) {
	resp, err := http.Get("http://localhost:8086/chloe")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	t.Log(string(body))
}
