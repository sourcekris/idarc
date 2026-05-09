package idarc

import (
	"encoding/binary"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// IsExe checks if the file is an EXE or COM file.
func IsExe(path string) (bool, int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, 0, err
	}
	defer f.Close()

	header := make([]byte, 2)
	_, err = f.Read(header)
	if err != nil {
		if err == io.EOF {
			return false, 0, nil
		}
		return false, 0, err
	}

	if (header[0] == 'M' && header[1] == 'Z') || (header[0] == 'Z' && header[1] == 'M') {
		size, err := ExeSize(path)
		if err != nil {
			return true, 0, nil // Assume it's an EXE but size couldn't be determined
		}
		if size > 0 {
			return true, size, nil
		}
	}

	ext := filepath.Ext(path)
	if strings.EqualFold(ext, ".com") {
		return true, 0, nil
	}

	return false, 0, nil
}

// ExeSize calculates the size of the EXE from its header.
func ExeSize(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return -1, err
	}
	defer f.Close()

	header := make([]byte, 6)
	_, err = f.ReadAt(header, 0)
	if err != nil {
		return -1, err
	}

	// Offset 2: bytes in last 512-byte page
	// Offset 4: number of 512-byte pages
	lastPageSize := int64(binary.LittleEndian.Uint16(header[2:4]))
	pages := int64(binary.LittleEndian.Uint16(header[4:6]))

	if pages == 0 {
		return 0, nil
	}

	// Standard formula: (pages-1)*512 + lastPageSize
	// The Pascal code used ((pages-1) Mod 512) shl 9 + lastPageSize
	// shl 9 is * 512. Mod 512 is redundant for uint16 if pages < 512, 
	// but might be used if pages >= 512? No, pages is the total count.
	// Actually, some EXE headers might have weird values.
	
	size := (pages-1)*512 + lastPageSize
	return size, nil
}
