package main

// Sync photos

// Poller: Timer / Trigger / UDP

// URL -> Content Reader
// URL -> Stats Reader/Writer (mod, time, size)

// Collector: path -> []i
// Xtra : filepath, title, description, keywords, image dimension
// Cache: filepath, mod, size, hash
// Hasher: filepath, hash
// Evaluator: filepath, result (Same/Different)

// NanoGen: Generating nanoJS HTML that can be added to a Hugo based site

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type pollerFunction interface {
	call(f func(t time.Time))
}

type pollerInterface interface {
	poll(fun *pollerFunction)
}

type pollerByTime struct {
	pollerInterface
}

func (p *pollerByTime) poll(fun *pollerFunction) {

}

// D:/path/filename.ext
// gdrive:/path/filename.ext
// s3:/path/filename.ext
func openURL(url string) (r io.Reader) {

	return
}

type hash256 struct {
	digest [32]byte
}

type tab struct {
	path  string
	items map[hash256]*syncItem
}

type session struct {
	left  tab
	right tab
}

type itemStats struct {
	mod  time.Time
	size int64
}

type itemHash struct {
	hash hash256
}

type estate int

const (
	same    estate = 1
	changed estate = 2
	missing estate = 3
	removed estate = 4
)

type itemState struct {
	state estate
}

type itemXtra struct {
	title    string
	descr    string
	keywords []string
	width    int32
	height   int32
}

type itemPath struct {
	path string
}

type syncItem struct {
	path  itemPath
	stats itemStats
	hash  itemHash
	state itemState
	xtra  itemXtra
}

func collectItems(t *tab, ignoreDirs []string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			dir := filepath.Base(path)
			for _, d := range ignoreDirs {
				if d == dir {
					return filepath.SkipDir
				}
			}
		} else {
			item := &syncItem{path: itemPath{path: path}}
			digest := sha256.Sum256([]byte(strings.ToLower(path)))
			h := hash256{digest: digest}
			t.items[h] = item
		}
		return nil
	}
}

func collect(t *tab) {
	ignoreDirs := []string{".bzr", ".hg", ".git"}
	filepath.Walk(t.path, collectItems(t, ignoreDirs))
}

func readXtra(filepath string) (x itemXtra) {

	return
}

func readStats(filepath string) (s itemStats) {

	return
}

func calcHash(filepath string) (h itemHash) {

	return
}

func evalState(filepath string) (s itemState) {

	return
}

func aUseCase() {
	fp1 := NewFilePath("D:/folder/filename.a.b.c")
	dp1 := NewDirPath("D:/folder1/folder2/")
	fp1.Print()
	dp1.Print()
	if fp1.Up() == false {
		fmt.Println("failed to move up!")
	}
	fp1.Print()
	if fp1.Root() {
		fmt.Println("at root")
	}

	fp1.Down("newfolder")
	fp1.Print()
	fp1.ChExt(".x")
	fp1.Print()
	fp1.ChExt("ext")
	fp1.Print()

	var wp1 WalkPath
	wp1 = WalkPath{DirPath: fp1.DirPath}
	wp1.Print()
}

func main() {
	aUseCase()
}
