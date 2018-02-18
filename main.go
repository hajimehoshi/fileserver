// Copyright 2018 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"net/http"
	"strconv"
)

var (
	port = flag.Int("port", 8000, "port number")
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	// Don't use http.ServeFile due to directory traversal attack.
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":" + strconv.Itoa(*port), nil); err != nil {
		panic(err)
	}
}
