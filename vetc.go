// Public domain.

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	ret := 0
	b := make([]byte, 200)
	r := regexp.MustCompile(`(?is:copyright.+20\d\d.+license|public domain)`)
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		if info.IsDir() {
			return nil // no action for directory entry
		}
		if matched, err := filepath.Match("*.go", info.Name()); err != nil {
			log.Fatal("fatal:", err) // malformed pattern
		} else if !matched {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			log.Println(err)
			return nil
		}
		defer f.Close()
		n, err := f.Read(b)
		if err != nil {
			log.Println(err)
			return nil
		}
		if !r.Match(b[:n]) {
			if i := bytes.IndexByte(b[:n], '\n'); i >= 0 {
				n = i
			}
			fmt.Printf("%s: %s\n", path, b[:n])
			ret = 1
		}
		return nil
	})
	os.Exit(ret)
}
