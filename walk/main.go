package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type Entry interface {
	Description()
}

type Directory struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Children []Entry `json:"children,omitempty"`
}

func (d *Directory) Description() {
	for _, entry := range d.Children {
		entry.Description()
	}
}

func (d *Directory) Add(entries []Entry) {
	d.Children = append(d.Children, entries...)
}

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (f *File) Description() {
	fmt.Println(f.Path)
}

func Walk(root string, ignores []string) {
	dir := &Directory{Name: "pebbles", Path: "pebbles"}
	entries, _ := walk(dir.Path, ignores)
	dir.Add(entries)
	//dir.Description()
	b, err := json.Marshal(dir)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func walk(path string, ignores []string) ([]Entry, error) {
	entries := []Entry{}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)

	for _, name := range names {
		if isIgnore(ignores, name) {
			continue
		}

		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			fmt.Println("Error")
		} else {
			if fileInfo.IsDir() {
				dir := &Directory{Name: name, Path: filename}
				_entries, _ := walk(filename, ignores)
				dir.Add(_entries)
				entries = append(entries, dir)
			} else {
				file := &File{Name: name, Path: filename}
				entries = append(entries, file)
			}
		}
	}

	return entries, nil
}

func isIgnore(ignores []string, file string) bool {
	for _, ign := range ignores {
		if ign == file {
			return true
		}
	}
	return false
}

func main() {
	Walk("pebbles", []string{".git"})
}
