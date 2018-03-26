package main

// Simplest disk cache storing strings in /tmp/dcache
// should be compatible with dcache.js

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var STORE_DIR = "/tmp/dcache"

func p(fs string, args ...interface{}) {
	fmt.Printf(fs+"\n", args...)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	err := os.MkdirAll(STORE_DIR, os.ModePerm)
	panicOn(err)
}

func hash(s string) string {
	arr := md5.Sum([]byte(s))
	return hex.EncodeToString(arr[:])
}

func SetString(k string, v string) string {
	f := filepath.Join(STORE_DIR, "string_"+hash(k))
	err := ioutil.WriteFile(f, []byte(v), os.ModePerm)
	panicOn(err)
	return f
}

func GetString(k string) string {
	f := filepath.Join(STORE_DIR, "string_"+hash(k))
	b, err := ioutil.ReadFile(f)
	if os.IsNotExist(err) {
		return ""
	} else if err != nil {
		panicOn(err)
	}
	return string(b)
}

func Set(k string, v interface{}) string {
	f := filepath.Join(STORE_DIR, "json_"+hash(k))
	b, err := json.Marshal(v)
	panicOn(err)
	err = ioutil.WriteFile(f, b, os.ModePerm)
	panicOn(err)
	return f
}

func Get(k string) interface{} {
	f := filepath.Join(STORE_DIR, "json_"+hash(k))
	b, err := ioutil.ReadFile(f)
	if os.IsNotExist(err) {
		return ""
	} else if err != nil {
		panicOn(err)
	}
	var v interface{}
	err = json.Unmarshal(b, &v)
	return v
}

func main() {

	p("Started")

	url := "allegro.pl/dozownik-na-plyn-do-naczyn-pojemnik-na-gabke-wenko-i6924188275.html"

	vvv := map[string]int{"first": 1, "second": 2}

	p1 := Set(url, vvv)
	p("p=%s", p1)

	html := Get(url)
	p("html=%v", html)
}
