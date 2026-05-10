package idarc

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Add the complex checks to the Detect method
func (d *Detector) DetectComplex() DetectionResult {
	// AINEXE
	// Pascal: BlockRead 34 bytes, check Pos('AIN') == 33
	if d.Size >= 34 {
		// We already read up to 255 bytes
		if len(d.IDStr) >= 34 {
			if d.CheckAnySign(32, []byte("AIN")) {
				return DetectionResult{Type: AINEXE}
			}
		}
	}

	// UC2
	if bytes.HasPrefix(d.IDStr, []byte("UC2")) || bytes.HasPrefix(d.IDStr, []byte("UC2SFX Header")) {
		return DetectionResult{Type: UC2}
	}

	// UCEXE
	if d.Size >= 32 {
		if d.CheckAnySign(28, []byte("UC2X")) {
			return DetectionResult{Type: UCEXE}
		}
	}

	// ACB
	if len(d.IDStr) >= 4 {
		if (d.IDStr[1] >= 0x80 && d.IDStr[1] <= 0x84) && d.IDStr[3] == 0 {
			return DetectionResult{Type: ACB}
		}
	}

	// CPZ
	if len(d.IDStr) >= 4 {
		if d.IDStr[1] == 0 && d.IDStr[2] == 0 && d.IDStr[3] == 0 && strings.EqualFold(filepath.Ext(d.Path), ".cpz") {
			return DetectionResult{Type: CPZ}
		}
	}
	
	// AIN
	if bytes.HasPrefix(d.IDStr, []byte{33, 18}) || bytes.HasPrefix(d.IDStr, []byte{33, 17}) {
		return DetectionResult{Type: AIN}
	}

	// LZH / LHark / CAR
	if d.Pos("-lh") == 3 || d.Pos("-lz") == 3 {
		if d.Pos("-lh7-") == 3 {
			return DetectionResult{Type: LHark}
		}
		if strings.EqualFold(filepath.Ext(d.Path), ".car") {
			return DetectionResult{Type: CAR}
		}
		return DetectionResult{Type: LZH}
	}

	// BZip2
	if len(d.IDStr) >= 5 && bytes.HasPrefix(d.IDStr, []byte("BZh")) {
		if d.IDStr[3] >= '0' && d.IDStr[3] <= '9' {
			return DetectionResult{Type: BZip2}
		}
	}

	// BZip
	if len(d.IDStr) >= 5 && bytes.HasPrefix(d.IDStr, []byte("BZ")) {
		if d.IDStr[2] >= '0' && d.IDStr[2] <= '9' && d.IDStr[3] >= '0' && d.IDStr[3] <= '9' {
			return DetectionResult{Type: BZip}
		}
	}

	// TAR
	if strings.Contains(strings.ToUpper(filepath.Base(d.Path)), ".TAR") {
		return DetectionResult{Type: TAR}
	}

	// 7z
	if bytes.HasPrefix(d.IDStr, []byte("7z\xbc\xaf\x27\x1c")) || bytes.HasPrefix(d.IDStr, []byte("7z\xbc\xaf")) {
		return DetectionResult{Type: _7z}
	}

	// MAR Utility
	if bytes.HasPrefix(d.IDStr, []byte("MAR\x80file")) {
		return DetectionResult{Type: MARUtilityFile}
	}
	if bytes.HasPrefix(d.IDStr, []byte("MAR\x80fold")) {
		return DetectionResult{Type: MARUtilityFolder}
	}

	// MeltingPotArchiver
	if bytes.HasPrefix(d.IDStr, []byte("MAr0")) {
		return DetectionResult{Type: MeltingPotArchiver}
	}

	// 777
	if bytes.HasPrefix(d.IDStr, []byte("777")) {
		return DetectionResult{Type: _777}
	}

	// Arhangel LG
	if bytes.HasPrefix(d.IDStr, []byte("LG")) {
		return DetectionResult{Type: Arhangel} // Note: LgType is 122 which is Arhangel
	}

	// BIX
	if bytes.HasPrefix(d.IDStr, []byte("BIX0")) {
		return DetectionResult{Type: BIX}
	}

	// BLINK
	if bytes.HasPrefix(d.IDStr, []byte("Blink")) {
		return DetectionResult{Type: Blink}
	}

	// CAR (MylesHi!) is already handled in LZH logic, but wait, CaCType = 106 is also .CAR extension
	if strings.Contains(strings.ToUpper(filepath.Base(d.Path)), ".CAR") && d.Pos("-lh") != 3 {
		return DetectionResult{Type: CARComp}
	}

	// LBR
	if strings.Contains(strings.ToUpper(filepath.Base(d.Path)), ".LBR") {
		return DetectionResult{Type: LBR}
	}

	// IMP
	if bytes.HasPrefix(d.IDStr, []byte("IMP\n")) {
		return DetectionResult{Type: IMP}
	}

	// ChiefLZA
	if d.Pos("ChfLZ") == 4 {
		return DetectionResult{Type: ChiefLZA}
	}

	// PFT
	if bytes.HasPrefix(d.IDStr, []byte{0, 0x50, 0, 0x14}) {
		return DetectionResult{Type: PFT}
	}

	// PUT
	if d.Pos("-lZ") == 3 && len(d.IDStr) > 6 && d.IDStr[6] == '-' {
		return DetectionResult{Type: PUT}
	}

	// XPACK (XPA32)
	if bytes.HasPrefix(d.IDStr, []byte("xpa\x00\x01")) {
		return DetectionResult{Type: XPA32}
	}

	// XPACK Data
	if bytes.HasPrefix(d.IDStr, []byte("xpa")) && !bytes.HasPrefix(d.IDStr, []byte("xpa\x00\x01")) {
		return DetectionResult{Type: XPackData}
	}

	// YZ1 (DeepFreezer)
	if len(d.IDStr) >= 5 && bytes.HasPrefix(d.IDStr, []byte("yz0")) {
		if d.IDStr[3] >= '1' && d.IDStr[3] <= '9' && d.IDStr[4] >= '0' && d.IDStr[4] <= '9' {
			return DetectionResult{Type: DeepFreezer}
		}
	}

	// LZOP
	if bytes.HasPrefix(d.IDStr, []byte("\x89LZO")) {
		return DetectionResult{Type: LZO}
	}

	// HAP
	if bytes.HasPrefix(d.IDStr, []byte("\x91\x33HF")) {
		return DetectionResult{Type: HAP}
	}

	// MSXiE
	if len(d.IDStr) >= 4 && bytes.HasPrefix(d.IDStr, []byte("MS")) {
		if d.IDStr[2] >= 0 && d.IDStr[2] <= 15 && d.IDStr[3] >= 0 && d.IDStr[3] <= 9 {
			return DetectionResult{Type: MSXiE}
		}
	}

	// NSK
	if bytes.HasPrefix(d.IDStr, []byte("NSK")) {
		return DetectionResult{Type: NaShrink} // NskType = 132
	}

	// ASD
	if bytes.HasPrefix(d.IDStr, []byte("ASD")) {
		return DetectionResult{Type: ASD}
	}

	// JARCS (JArcs)
	if bytes.HasPrefix(d.IDStr, []byte("JARCS")) {
		return DetectionResult{Type: JArcs}
	}

	// JRchive
	if bytes.HasPrefix(d.IDStr, []byte("JRchive")) {
		return DetectionResult{Type: JRC}
	}
	
	// ARJSoftwareJAR
	if d.Pos("\x1aJar\x1b\x00") == 15 {
		return DetectionResult{Type: JAR}
	}

	// ZET (4F 5A DD is "OZ" + #$DD)
	if bytes.HasPrefix(d.IDStr, []byte{0x4F, 0x5A, 0xDD}) {
		return DetectionResult{Type: ZET}
	}

	// GCA
	if bytes.HasPrefix(d.IDStr, []byte("GCAX")) {
		return DetectionResult{Type: GCA}
	}

	// Crush / uncomp
	p := d.Pos("CRUSH$")
	q := d.Pos(".CRU")
	r := d.Pos(".cru")
	o := d.Pos("CRUSH")
	if p == 35 || (q >= 24 && q <= 39) || (r >= 24 && r <= 31) || (r >= 82 && r <= 89) || o == 1 {
		if o == 1 {
			return DetectionResult{Type: CrushUncomp}
		} else if bytes.HasPrefix(d.IDStr, []byte("PK\x03\x04")) {
			return DetectionResult{Type: CrushZIP}
		} else if bytes.HasPrefix(d.IDStr, []byte("\x60\xea")) || d.Pos("\x60\xea") == 3 || d.Pos(".ARJ") == 41 {
			return DetectionResult{Type: CrushARJ}
		} else if d.Pos("-lh") == 3 {
			return DetectionResult{Type: CrushLZH}
		} else if bytes.HasPrefix(d.IDStr, []byte("ZOO")) {
			return DetectionResult{Type: CrushZOO}
		} else if bytes.HasPrefix(d.IDStr, []byte("HA")) {
			return DetectionResult{Type: CrushHA}
		}
	}

	// ArcMethod (ARC, PAK, HYP, ARCPlus)
	arcMethodRes := d.ArcMethod()
	if arcMethodRes != Invalid {
		return DetectionResult{Type: arcMethodRes}
	}

	return DetectionResult{Type: Invalid}
}

func (d *Detector) ArcMethod() ArchiveType {
	f, err := os.Open(d.Path)
	if err != nil {
		return Invalid
	}
	defer f.Close()

	if d.Offset > 0 {
		f.Seek(d.Offset, io.SeekStart)
	}

	// ARCHeader: Marker(1), Method(1), Name(13), Size(4), Stamp(4), CRC(2), Length(4) = 29 bytes
	header := make([]byte, 29)
	for {
		_, err := io.ReadFull(f, header)
		if err != nil {
			break
		}

		if header[0] == 0x1A { // ARCId
			method := header[1]
			if method >= 0x0A { // PAKId
				if header[13] == 0x14 && header[14] == 0x15 && header[15] == 0x13 { // Name[11..13] (1-indexed) -> header[12..14]
					return PAK
				}
				if method >= 0x48 { // HYPId
					return HYP
				}
				if method == 0x14 { // ARPId
					return ARCPlus
				}
				return PAK // Fallback if Method >= PAKId and not caught above
			}
			return ARC
		}
		break
	}

	return Invalid
}
