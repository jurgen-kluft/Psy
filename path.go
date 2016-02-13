package main

import (
	"fmt"
	"path"
)

// Path is the base structure holding the full URL
type Path struct {
	url string
}

// FilePath contains the main URL and the filename (slice)
type FilePath struct {
	DirPath
	filename string
}

// DirPath represents a part of a FilePath
type DirPath struct {
	DevicePath
	directory string
}

// AliasPath is aliasing an actual physical 'device:\path\'
// Example: {APPDIR}config.ini, expands to 'C:\Programs\App\config.ini'
type AliasPath struct {
	dirpath DirPath
	alias   string // The name of the alias (all uppercase), used as {NAME}
}

// AliasPathRegistry contains all the aliases and is used by some functions to
// resolve the alias in a path
type AliasPathRegistry struct {
	aliases map[string]string
}

// DevicePath is a slice of the main URL
type DevicePath struct {
	Path
	device string
}

// WalkPath is used when browsing for a directory
type WalkPath struct {
	DirPath
}

func sanitizePath(pathstr string) string {
	pathstr = path.Clean(pathstr)
	return pathstr
}

// Return the slice that represents the filename
func sliceFilename(pathstr string) string {
	// Search backwards until encountering '\' or start of string
	end := len(pathstr)
	pos := end - 1
	for pos >= 0 {
		c := pathstr[pos]
		if c == '\\' {
			return pathstr[pos+1 : end]
		}
		pos--
	}
	return pathstr[pos:end]
}

// NewFilePath constructs a new filepath from a string
func NewFilePath(pathstr string) (fp FilePath) {
	str := sanitizePath(pathstr)
	fp = FilePath{}
	fp.url = str
	return
}

// NewDirPath constructs a new dirpath from a string
func NewDirPath(pathstr string) (dp DirPath) {
	str := sanitizePath(pathstr)
	dp = DirPath{}
	dp.url = str
	return
}

// NewWalkPath constructs a new WalkPath from a string and directory-level
func NewWalkPath(pathstr string, level int) (wp WalkPath) {
	return
}

//

// Root returns true if we are at the root?
// Note: In the case of an alias in the path 'root' means relative to the alias
func (d *DirPath) Root() bool {
	return false
}

// ChDir changes the directory part of the path
func (d *DirPath) ChDir(dp DirPath) {
	return
}

// MkAlias makes the current WalkPath.dirpath relative and returns the absolute
// part as an alias
func MkAlias(wp WalkPath) (abs DirPath, rel DirPath) {
	return
}

// Up moves the path up to the parent directory
func (w *WalkPath) Up() {
	if w.DirPath.Up() {

	}
}

// Up moves the path up to the parent directory
func (d *DirPath) Up() bool {
	// Search backwards until encountering '\' or start of string
	end := len(d.directory)
	if end == 0 {
		return false
	}
	pos := end
	for pos >= 0 {
		c := d.directory[pos]
		if c == '\\' {
			d.directory = d.directory[0 : pos-1]
			return true
		}
		pos--
	}
	d.directory = d.url[0:0]
	return true
}

// Sub moves the path further into a sub-directory
func (w *WalkPath) Sub(folder string) {
	return
}

// ChExt changes the file extension
func (f *FilePath) ChExt(ext string) {

}

// Print prints the path to stdout
func (p *Path) Print() {
	fmt.Println(p.url)
}

// HasSameDevice checks if they are the same device
func HasSameDevice(d1 DevicePath, d2 DevicePath) bool {
	return false
}

// HasSameDirectory checks if they both are identical directory paths
func HasSameDirectory(d1 DirPath, d2 DirPath) bool {
	return false
}

// HasSameFilename checks if both paths have the same filename
func HasSameFilename(f1 FilePath, f2 FilePath) bool {
	return false
}

// HasSameExtension checks if both filepaths have the same extension
func HasSameExtension(f1 FilePath, f2 FilePath) bool {
	return false
}
