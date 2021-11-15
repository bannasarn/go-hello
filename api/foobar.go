package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/bannasarn/go-hello/foo"
)

type foobarResp struct {
	Param  int    `json:"param"`
	Foobar string `json:"foobar"`
}

func ApiFooBar() {
	// Hello world, the web server

	foobarHandler := func(w http.ResponseWriter, req *http.Request) {

		// Get Param
		param := strings.TrimPrefix(req.RequestURI, "/foobar/")

		// Convert Param String -> Int
		n, err := strconv.Atoi(param)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err.Error())
			return
		}

		// Assign value to response struct
		var resp foobarResp
		resp.Param = n
		resp.Foobar = foo.String(n)

		// Create JSON response header
		w.Header().Set("Context-Type", "application/json")
		json.NewEncoder(w).Encode(&resp)
	}

	http.HandleFunc("/foobar/", foobarHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
