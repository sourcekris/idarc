package idarc

import (
	"bytes"
	"io"
	"os"
)

// CheckAnySign checks for a signature at a specific offset in the file.
func (d *Detector) CheckAnySign(offset int64, pattern []byte) bool {
	if offset+int64(len(pattern)) > d.Size {
		return false
	}

	f, err := os.Open(d.Path)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Seek(offset, io.SeekStart)
	if err != nil {
		return false
	}

	buf := make([]byte, len(pattern))
	_, err = f.Read(buf)
	if err != nil {
		return false
	}

	return bytes.Equal(buf, pattern)
}

// Porting ArjWinSfxPacked
func (d *Detector) ArjWinSfxPacked() bool {
	if d.CheckAnySign(900, []byte("ARJSFX")) {
		// In Pascal, this also reloads IDStr from a different offset
		// For now we just return true
		return true
	}
	return false
}

// Porting ArjDOSSfxPacked
func (d *Detector) ArjDOSSfxPacked() bool {
	offsets := []int64{225, 567, 663, 664}
	for _, off := range offsets {
		if d.CheckAnySign(off, []byte("ARJSFX")) {
			return true
		}
	}
	return false
}
