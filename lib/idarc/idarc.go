package idarc

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type DetectionResult struct {
	Type           ArchiveType
	MultipleVolume bool
	AV             bool
}

type Detector struct {
	Path   string
	IDStr  []byte // First 255 bytes after EXE stub
	Size   int64
	Offset int64
}

func NewDetector(path string) (*Detector, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	d := &Detector{
		Path: path,
		Size: info.Size(),
	}

	isEx, exeSize, err := IsExe(path)
	if err == nil && isEx {
		d.Offset = exeSize
	}

	if d.Size <= d.Offset {
		d.IDStr = []byte{}
		return d, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Seek(d.Offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	bufLen := 255
	if d.Size-d.Offset < int64(bufLen) {
		bufLen = int(d.Size - d.Offset)
	}

	d.IDStr = make([]byte, bufLen)
	_, err = f.Read(d.IDStr)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return d, nil
}

func (d *Detector) Detect() DetectionResult {
	if len(d.IDStr) == 0 {
		return DetectionResult{Type: Invalid}
	}

	// Porting the logic from ArchiveType(ArcName: PathStr): Byte
	
	// Crush detection (Complex logic)
	// TODO: Port Crush logic
	
	// ZIP
	if bytes.HasPrefix(d.IDStr, []byte("PK\x03\x04")) || bytes.HasPrefix(d.IDStr, []byte("PK00PK")) {
		if bytes.Contains(d.IDStr, []byte("BZh")) {
			return DetectionResult{Type: PKZip6BZip2}
		}
		return DetectionResult{Type: ZIP}
	}

	// GZIP
	if bytes.HasPrefix(d.IDStr, []byte("\x1f\x8b\x08\x08")) || bytes.HasPrefix(d.IDStr, []byte("\x1f\x9d\x90")) {
		return DetectionResult{Type: GZIP}
	}

	// ARJ
	if bytes.HasPrefix(d.IDStr, []byte("\x60\xea")) || (len(d.IDStr) >= 5 && bytes.Equal(d.IDStr[2:4], []byte("\x60\xea"))) ||
		d.ArjWinSfxPacked() || d.ArjDOSSfxPacked() {
		res := DetectionResult{Type: ARJ}
		// Volume and AV flags would go here
		if strings.Contains(strings.ToUpper(filepath.Base(d.Path)), ".SRJ") {
			res.Type = SARJ
		}
		return res
	}

	// RAR
	if bytes.HasPrefix(d.IDStr, []byte("RE~\x5e")) || bytes.HasPrefix(d.IDStr, []byte("Rar!")) {
		return DetectionResult{Type: RAR}
	}

	// ACE
	if d.Pos("**ACE**") == 8 {
		return DetectionResult{Type: ACE}
	}

	// Table-driven matching for simple signatures
	if t := d.matchSignatures(); t != Invalid {
		return DetectionResult{Type: t}
	}

	res := d.DetectComplex()
	if res.Type != Invalid {
		return res
	}

	// Default
	return DetectionResult{Type: Invalid}
}

func (d *Detector) Pos(pattern string) int {
	idx := bytes.Index(d.IDStr, []byte(pattern))
	if idx == -1 {
		return 0
	}
	return idx + 1 // 1-indexed like Pascal
}

func (d *Detector) PosBytes(pattern []byte) int {
	idx := bytes.Index(d.IDStr, pattern)
	if idx == -1 {
		return 0
	}
	return idx + 1 // 1-indexed like Pascal
}
