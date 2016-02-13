package main

import (
	"fmt"
	"path"
	"strings"
)

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

// AliasPath is aliasing an actual physical 'device:/path/'
// Example:
//          APPDIR:/config.ini
//          APPDIR=C:/Programs/
//          expands to 'C:/Programs/App/config.ini'
type AliasPath struct {
	dirpath DirPath
	alias   string // The name of the alias (all uppercase)
}

// AliasPathRegistry contains all the aliases and is used by some functions to
// resolve the alias in a path
// Initially populated with all drives:
//    C=C:/
//    D=D:/
//    E=E:/
type AliasPathRegistry struct {
	aliases map[string]string
}

// DevicePath is a slice of the main URL
type DevicePath struct {
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
func slicePath(pathstr string) (device string, directory string, filename string) {
	// Search backwards until encountering '/' or start of string
	end := len(pathstr)
	pos := end - 1
	for pos >= 0 {
		c := pathstr[pos]
		if c == '/' {
			dev := strings.Index(pathstr, ":")
			if dev > 0 {
				device = pathstr[0 : dev+2]
			} else {
				device = pathstr[0:0]
			}
			if (dev + 2) < end {
				directory = pathstr[dev+2 : pos+1]
				if (pos + 1) < end {
					filename = pathstr[pos+1 : end]
				} else {
					filename = pathstr[0:0]
				}
			} else {
				directory = pathstr[0:0]
				filename = pathstr[0:0]
			}
			return
		}
		pos--
	}
	device = pathstr[0:0]
	directory = pathstr[0:0]
	filename = pathstr[pos+1 : end]
	return
}

// NewFilePath constructs a new filepath from a string
func NewFilePath(pathstr string) (fp FilePath) {
	str := sanitizePath(pathstr)
	fp = FilePath{}
	fp.device, fp.directory, fp.filename = slicePath(str)
	return
}

// NewDirPath constructs a new dirpath from a string
func NewDirPath(pathstr string) (dp DirPath) {
	pathstr = sanitizePath(pathstr)
	dp = DirPath{}
	dp.device, dp.directory, _ = slicePath(pathstr)
	return
}

// NewWalkPath constructs a new WalkPath from a string and directory-level
func NewWalkPath(pathstr string, level int) (wp WalkPath) {
	return
}

// Root returns true if we are at the root
// Note: In the case of an alias in the path 'root' means relative to the alias
func (d *DirPath) Root() bool {
	return len(d.directory) == 0
}

// IsRelative returns true if this path is relative
func (d *DirPath) IsRelative() bool {
	return len(d.device) == 0
}

// ChDir changes the directory part of the path
func (d *DirPath) ChDir(dp DirPath) {
	d.directory = dp.directory
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
	pos := end - 1
	pos--
	for pos >= 0 {
		c := d.directory[pos]
		if c == '/' {
			d.directory = d.directory[0 : pos+1]
			return true
		}
		pos--
	}
	d.directory = d.directory[0:0]
	return true
}

// Down moves the path further into a sub-directory
func (d *DirPath) Down(folder string) {
	d.directory = fmt.Sprintf("%s%s/", d.directory, folder)
	return
}

// ChExt changes the file extension
func (f *FilePath) ChExt(ext string) {
	if len(ext) > 0 && ext[0] == '.' {
		// find the first '.' from the left
		pos := strings.LastIndex(f.filename, ".")
		// replace anything on the right with the new extension
		f.filename = fmt.Sprintf("%s%s", f.filename[0:pos], ext)
	} else {
		// find the first '.' from the left
		pos := strings.Index(f.filename, ".")
		// replace anything on the right with the new extension
		f.filename = fmt.Sprintf("%s%s", f.filename[0:pos+1], ext)
	}
}

// Print prints the file-path to stdout
func (f *FilePath) Print() {
	fmt.Printf("%s+%s+%s\n", f.device, f.directory, f.filename)
}

// Print prints the dir-path to stdout
func (d *DirPath) Print() {
	fmt.Printf("%s+%s\n", d.device, d.directory)
}

// Print prints the device-path to stdout
func (d *DevicePath) Print() {
	fmt.Printf("%s\n", d.device)
}

// IsSame returns true if the two file-path's are the same
func IsSame(f1 FilePath, f2 FilePath) bool {
	return strings.EqualFold(f1.device, f2.device) && strings.EqualFold(f1.directory, f2.directory) && strings.EqualFold(f1.filename, f2.filename)
}

// HasSameDevice checks if they are the same device
func HasSameDevice(d1 DevicePath, d2 DevicePath) bool {
	return strings.EqualFold(d1.device, d2.device)
}

// HasSameDirectory checks if they both are identical directory paths
func HasSameDirectory(d1 DirPath, d2 DirPath) bool {
	return strings.EqualFold(d1.directory, d2.directory)
}

// HasSameFilename checks if both paths have the same filename
func HasSameFilename(f1 FilePath, f2 FilePath) bool {
	return strings.EqualFold(f1.filename, f2.filename)
}

// HasSameExtension checks if both filepaths have the same extension
func HasSameExtension(f1 FilePath, f2 FilePath) bool {
	// find the first '.' from the left
	pos1 := strings.Index(f1.filename, ".")
	pos2 := strings.Index(f2.filename, ".")
	if pos1 == -1 || pos2 == -1 {
		return false
	}
	return strings.EqualFold(f1.filename[pos1:], f2.filename[pos2:])
}
