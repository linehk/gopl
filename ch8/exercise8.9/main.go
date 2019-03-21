package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type dir struct {
	id   int // 用来指示是哪个目录
	size int64
}

func main() {
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	info := make(chan dir)
	var n sync.WaitGroup
	for id, root := range roots {
		n.Add(1)
		go walkDir(root, &n, id, info)
	}
	go func() {
		n.Wait()
		close(info)
	}()

	tick := time.Tick(500 * time.Millisecond)
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case dir, ok := <-info:
			if !ok {
				break loop
			}
			nfiles[dir.id]++
			nbytes[dir.id] += dir.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(roots, nfiles, nbytes) // final totals
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for id, root := range roots {
		fmt.Printf("%d files %.1f GB in %s\n",
			nfiles[id], float64(nbytes[id])/1e9, root)
	}
}

func walkDir(d string, n *sync.WaitGroup, root int, info chan<- dir) {
	defer n.Done()
	for _, entry := range dirents(d) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(d, entry.Name())
			go walkDir(subdir, n, root, info)
		} else {
			info <- dir{root, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
