package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var db = database{"shoes": 50, "socks": 5}

func main() {
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var mu sync.Mutex

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "params err: %s\n", err)
	}
	mu.Lock()
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s is exist\n", item)
	} else {
		db[item] = dollars(p)
	}
	mu.Unlock()
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s isn't exist\n", item)
	} else {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "params err: %s\n", err)
	}
	mu.Lock()
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s isn't exist\n", item)
	} else {
		db[item] = dollars(p)
	}
	mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s isn't exist\n", item)
	} else {
		mu.Lock()
		delete(db, item)
		mu.Unlock()
	}
}
