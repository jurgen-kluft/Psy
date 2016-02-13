package main

// Sync photos

// Poller: Timer / Trigger / UDP

// URL -> Content Reader
// URL -> Stats Reader/Writer (mod, time, size)

// Xtra : filepath, title, description, keywords, image dimension
// Cache: filepath, mod, size, hash
// Hasher: filepath, hash
// Evaluator: filepath, result (Same/Different)

// NanoGen: Generating nanoJS HTML that can be added to a Hugo based site

import (
	"io"
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

// file://path/filename.ext
// gdrive://path/filename.ext
// s3://path/filename.ext
func openURL(url string) (r io.Reader) {

	return
}

type item struct {
	filename string
}

type itemStats struct {
	mod  time.Time
	size int64
}

type itemHash struct {
	hash [32]byte
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

func readXtra(filename string) (x itemXtra) {

	return
}

func readStats(filename string) (s itemStats) {

	return
}

func calcHash(filename string) (h itemHash) {

	return
}

func evalState(filename string, srcpath string) (s itemState) {
	return
}

func aUseCase() {
	fp1 := NewFilePath("D:/folder/filename.ext")
	dp1 := NewDirPath("D:/folder1/folder2")
	fp1.Print()
	dp1.Print()
	fp1.Up()
	var wp1 WalkPath
	wp1 = WalkPath{DirPath: dp1}
	wp1.Print()
}

func main() {
	aUseCase()
}
