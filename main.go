package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Archive type constants
const (
	ARCType      = 1
	ZIPType      = 2
	ZOOType      = 3
	LZHType      = 4
	DWCType      = 5
	MDType       = 6
	LBRType      = 7
	ARJType      = 8
	HYPType      = 9
	UC2Type      = 10
	HAPType      = 11
	HAType       = 12
	HPKType      = 13
	SQZType      = 14
	RARType      = 15
	PAKType      = 16
	ARCPlusType  = 17
	LIMType      = 18
	BSNType      = 19
	PUTType      = 20
	SQWEZType    = 21
	CruPType     = 22
	CruJType     = 23
	CruLType     = 24
	CruZType     = 25
	CruHType     = 26
	LZEXEType    = 27
	PKLiteType   = 28
	DietType     = 29
	TinyProgType = 30
	GIFType      = 31
	JFIFType     = 32
	JHSIType     = 33
	AINType      = 34
	AINEXEType   = 35
	SARType      = 36
	BS2Type      = 37
	GZIPType     = 38
	ACBType      = 39
	MARType      = 40
	CPZType      = 41
	JRCType      = 42
	JARType      = 43
	QType        = 44
	SofType      = 45
	CruType      = 46
	ARXType      = 47
	UCEXEType    = 48
	WWPType      = 49
	QARKType     = 50
	YACType      = 51
	X1Type       = 52
	CDCType      = 53
	AMGType      = 54
	NLIType      = 55
	PLLType      = 56
	TGZType      = 57
	WWDType      = 58
	CHZType      = 59
	PSAType      = 60
	ZARType      = 61
	LHKType      = 62
	PACType      = 63
	XFType       = 64
	KBOType      = 65
	NSQType      = 66
	DPAType      = 67
	TTCType      = 68
	WICType      = 69
	RKVType      = 70
	JRType       = 71
	ESPType      = 72
	ZPKType      = 73
	DRYType      = 74
	OWSType      = 75
	SkyType      = 76
	ARIType      = 77
	UfaType      = 78
	CABType      = 79
	FSqzType     = 80
	AR7Type      = 81
	TSCType      = 82
	PPMZType     = 83
	ExpType      = 84
	MP3Type      = 85
	ZetType      = 86
	XpaType      = 87
	XdiType      = 88
	ArqType      = 89
	AceType      = 90
	ArhType      = 91
	TerType      = 92
	XpdType      = 93
	SitType      = 94
	PucType      = 95
	BZipType     = 96
	UhaType      = 97
	AbcType      = 98
	CmpType      = 99
	BZip2Type    = 100
	LzoType      = 101
	SzipType     = 102
	SplType      = 103
	TarType      = 104
	IShType      = 105
	CaCType      = 106
	LzsType      = 107
	BoaType      = 108
	IShZType     = 109
	ArgType      = 110
	GthType      = 111
	PckType      = 112
	BtsType      = 113
	EliType      = 114
	QfcType      = 115
	RncType      = 116
	XieType      = 117
	RaxType      = 118
	_777Type     = 119
	StacType     = 120
	HpaType      = 121
	LgType       = 122
	Exp1Type     = 123
	ImpType      = 124
	BmfType      = 125
	NrvType      = 126
	PddType      = 127
	SqType       = 128
	ParType      = 129
	HitType      = 130
	SbxType      = 131
	NskType      = 132
	DstType      = 133
	AsdType      = 134
	IscType      = 135
	T4Type       = 136
	BtmType      = 137
	BhType       = 138
	BixType      = 139
	LzaType      = 140
	BliType      = 141
	CarType      = 142
	SArjType     = 143
	CpkType      = 144
	LgCType      = 145
	ArsType      = 146
	AktType      = 147
	FlhType      = 148
	PC3Type      = 149
	NpaType      = 150
	PftType      = 151
	XTType       = 152
	SemType      = 153
	A32Type      = 154
	IiType       = 155
	PpmType      = 156
	SwgType      = 157
	FizType      = 158
	BaType       = 159
	Xpa32Type    = 160
	RKType       = 161
	RpmType      = 162
	DfType       = 163
	ZZType       = 164
	DCType       = 165
	TpcType      = 166
	AiType       = 167
	YbsType      = 168
	Ai32Type     = 169
	SbcType      = 170
	DitType      = 171
	DmsType      = 172
	EpcType      = 173
	VsaType      = 174
	PdzType      = 175
	PfwType      = 176
	NullType     = 177
	WiseType     = 178
	DZType       = 179
	_7zType      = 180
	RdqType      = 181
	ApkType      = 182
	ImaType      = 183
	GcaType      = 184
	PmnType      = 185
	SapType      = 186
	CpaType      = 187
	UhbType      = 188
	PKBZType     = 189

	Invalid      = 251
	FileNotFound = 255

	nIDPacker = 189 + 2
)

// Flags for archive properties
var (
	multiVolume = false
	avProtected = false
	idStr       = ""
)

// PackerNames maps archive type codes to names
var packerNames = map[int]string{
	1:   "ARC",
	2:   "ZIP",
	3:   "ZOO",
	4:   "LZH",
	5:   "DWC",
	6:   "MDCD",
	7:   "LBR",
	8:   "ARJ",
	9:   "HYP",
	10:  "UC2",
	11:  "HAP",
	12:  "HA",
	13:  "HPack",
	14:  "SQZ (Squeeze It)",
	15:  "RAR",
	16:  "PAK",
	17:  "ARC+",
	18:  "LIM",
	19:  "BSN (BSA/PTS-DOS)",
	20:  "PUT",
	21:  "SQWEZ",
	22:  "Crush/ZIP",
	23:  "Crush/ARJ",
	24:  "Crush/LZH",
	25:  "Crush/ZOO",
	26:  "Crush/HA",
	27:  "LZExe",
	28:  "PKLite",
	29:  "Diet",
	30:  "TinyProg",
	31:  "GIF",
	32:  "JPG (JFIF)",
	33:  "JPG (HSI)",
	34:  "AIN",
	35:  "AINEXE",
	36:  "SAR",
	37:  "BS2/BSArc",
	38:  "GZIP/Comp 4.3",
	39:  "ACB",
	40:  "MAR",
	41:  "CPZ (CPShrink)",
	42:  "JRC",
	43:  "JArcs",
	44:  "Quantum",
	45:  "ReSOF",
	46:  "Crush/uncompressed",
	47:  "ARX",
	48:  "UCEXE",
	49:  "WWPack",
	50:  "QuArk",
	51:  "YAC",
	52:  "X1",
	53:  "Codec",
	54:  "AMGC",
	55:  "NuLIB",
	56:  "PAKLeo",
	57:  "TGZ",
	58:  "WWPack-Data",
	59:  "ChArc",
	60:  "PSA",
	61:  "ZAR",
	62:  "LHark",
	63:  "CrossePAC",
	64:  "Freeze",
	65:  "KBoom",
	66:  "NSQ",
	67:  "DPA",
	68:  "TTComp",
	69:  "WIC (Fake!)",
	70:  "RKive",
	71:  "JAR",
	72:  "ESP",
	73:  "ZPack",
	74:  "DRY",
	75:  "OWS (Fake!)",
	76:  "SKY",
	77:  "ARI",
	78:  "UFA",
	79:  "Microsoft CAB",
	80:  "FOXSQZ",
	81:  "AR7",
	82:  "TSComp",
	83:  "PPMZ",
	84:  "MS Compress",
	85:  "MP3 (M.Czudej)",
	86:  "ZET",
	87:  "XPack Data",
	88:  "Xpk DiskImg",
	89:  "ARQ",
	90:  "ACE",
	91:  "Squash",
	92:  "Terse",
	93:  "Xpk SData",
	94:  "Stuffit (Mac)",
	95:  "PUCrunch",
	96:  "BZip",
	97:  "UHarc",
	98:  "ABComp",
	99:  "CMP",
	100: "BZip2",
	101: "LZOP",
	102: "szip",
	103: "Splint",
	104: "TAR",
	105: "InstallShield",
	106: "CARComp",
	107: "LZS",
	108: "BOA",
	109: "InstallSh. Z",
	110: "ARG",
	111: "Gather",
	112: "Pack Magic",
	113: "BTS",
	114: "ELI 5750",
	115: "QFC",
	116: "PRO-PACK",
	117: "MSXiE",
	118: "RAX",
	119: "777",
	120: "LZS221",
	121: "HPA",
	122: "Arhangel",
	123: "EXP1",
	124: "IMP",
	125: "BMF",
	126: "NRV",
	127: "oPAQue",
	128: "Squish",
	129: "Par",
	130: "HIT (B. Ureche)",
	131: "SBX",
	132: "NaShrinK",
	133: "Disintegrator",
	134: "ASD (T. Svensson)",
	135: "InStallSh. CAB",
	136: "TOP4",
	137: "BatComp (4DOS)",
	138: "BlakHole",
	139: "BIX (I. Pavlov)",
	140: "ChiefLZA",
	141: "Blink (D.T.S.)",
	142: "CAR (MylesHi!)",
	143: "SARJ",
	144: "CompackSfx",
	145: "LGExpand",
	146: "ARS-Sfx",
	147: "Akt",
	148: "Flash",
	149: "PC/3270",
	150: "NPack",
	151: "PFT",
	152: "Xtreme",
	153: "SemOne",
	154: "Akt32",
	155: "InstallIt",
	156: "PPMD",
	157: "Swag",
	158: "FIZ",
	159: "BA",
	160: "XPA32",
	161: "RK",
	162: "RPM",
	163: "DeepFreezer",
	164: "ZZip",
	165: "DC (E.Binder)",
	166: "TPac (Tim Gordon)",
	167: "Ai (E.Ilya)",
	168: "Ybs",
	169: "Ai32 (E.Ilya)",
	170: "SBC (Sami Mäkinen)",
	171: "DitPack",
	172: "DMS",
	173: "EPC",
	174: "VSARC",
	175: "PDZ",
	176: "PfW",
	177: "NullSoft Inst.",
	178: "Wise Installer",
	179: "DZip (N.Pflug)",
	180: "7z",
	181: "ReDuq",
	182: "aPackage",
	183: "WinImage",
	184: "GCA",
	185: "PPMN",
	186: "SAPCAR",
	187: "Compressia",
	188: "UHBC",
	189: "PKZip6/BZip2",
	251: "unknown",
	255: "not found",
}

func readFileHeader(filename string, maxBytes int) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, maxBytes)
	n, err := file.Read(buf)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}

	return buf[:n], nil
}

func identifyArchive(filename string) int {
	buf, err := readFileHeader(filename, 16384)
	if err != nil || len(buf) == 0 {
		return FileNotFound
	}

	// Check magic bytes for various formats
	if len(buf) >= 2 {
		// ZIP
		if buf[0] == 0x50 && buf[1] == 0x4b {
			if len(buf) >= 4 {
				if (buf[2] == 0x03 && buf[3] == 0x04) ||
					(buf[2] == 0x05 && buf[3] == 0x06) ||
					(buf[2] == 0x07 && buf[3] == 0x08) {
					return ZIPType
				}
			}
		}

		// RAR
		if len(buf) >= 7 {
			if buf[0] == 0x52 && buf[1] == 0x61 && buf[2] == 0x72 && 
			   buf[3] == 0x21 && buf[4] == 0x1a && buf[5] == 0x07 {
				if len(buf) >= 13 {
					if buf[6] == 0x00 {
						return RARType
					} else if buf[6] == 0x01 {
						return RARType
					}
				}
			}
		}

		// 7z
		if len(buf) >= 6 {
			if buf[0] == 0x37 && buf[1] == 0x7a && buf[2] == 0xbc && 
			   buf[3] == 0xaf && buf[4] == 0x27 && buf[5] == 0x1c {
				return _7zType
			}
		}

		// GZIP
		if buf[0] == 0x1f && buf[1] == 0x8b {
			return GZIPType
		}

		// BZIP2
		if buf[0] == 0x42 && buf[1] == 0x5a {
			if len(buf) >= 3 && (buf[2] == 0x68 || buf[2] == 0x30) {
				return BZip2Type
			}
		}

		// CAB (Microsoft)
		if buf[0] == 0x4d && buf[1] == 0x53 && len(buf) >= 4 {
			if buf[2] == 0x43 && buf[3] == 0x46 {
				return CABType
			}
		}

		// ARJ
		if buf[0] == 0x60 && buf[1] == 0xea {
			return ARJType
		}

		// ACE
		if len(buf) >= 4 {
			if buf[0] == 0x41 && buf[1] == 0x43 && buf[2] == 0x45 && buf[3] == 0x7a {
				return AceType
			}
		}

		// BZIP
		if buf[0] == 0x42 && buf[1] == 0x5a && len(buf) >= 3 && buf[2] == 0x68 {
			return BZip2Type
		}

		// TAR (check for tar magic)
		if len(buf) >= 512 {
			if len(buf) >= 265 && bytes.Equal(buf[257:262], []byte("ustar")) {
				return TarType
			}
		}

		// LZH
		if len(buf) >= 4 {
			if buf[0] == 0x2d && buf[1] == 0x6c && 
			   (buf[2] == 0x68 || buf[2] == 0x7a) {
				return LZHType
			}
		}

		// ARZ (compressed ARC)
		if buf[0] == 0x1a {
			return AINType
		}

		// ZOO
		if len(buf) >= 4 {
			if buf[0] == 0x5a && buf[1] == 0x4f && 
			   buf[2] == 0x4f && buf[3] == 0x20 {
				return ZOOType
			}
		}

		// GIF
		if len(buf) >= 3 {
			if buf[0] == 0x47 && buf[1] == 0x49 && buf[2] == 0x46 {
				return GIFType
			}
		}

		// JPEG (JFIF)
		if len(buf) >= 3 {
			if buf[0] == 0xff && buf[1] == 0xd8 && buf[2] == 0xff {
				return JFIFType
			}
		}

		// RPM
		if len(buf) >= 4 {
			if buf[0] == 0xed && buf[1] == 0xab && buf[2] == 0xee && buf[3] == 0xdb {
				return RpmType
			}
		}

		// DMS (Amiga)
		if len(buf) >= 4 {
			if buf[0] == 0x44 && buf[1] == 0x4d && buf[2] == 0x53 && buf[3] == 0x21 {
				return DmsType
			}
		}

		// ARC
		if buf[0] == 0x1a {
			if len(buf) >= 2 {
				if buf[1] >= 0x01 && buf[1] <= 0x09 {
					return ARCType
				}
			}
		}
	}

	// Check by file extension as fallback
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".zip":
		return ZIPType
	case ".rar":
		return RARType
	case ".7z":
		return _7zType
	case ".gz":
		return GZIPType
	case ".tar":
		return TarType
	case ".bz2":
		return BZip2Type
	case ".cab":
		return CABType
	case ".arj":
		return ARJType
	case ".ace":
		return AceType
	case ".rpm":
		return RpmType
	case ".exe":
		return Unknown
	default:
		return Invalid
	}
}

const Unknown = 0

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	filename := os.Args[1]

	// Check if file exists
	if _, err := os.Stat(filename); err != nil {
		os.Exit(FileNotFound)
	}

	arcType := identifyArchive(filename)

	// Reset flags
	multiVolume = false
	avProtected = false

	// Build status string
	status := ""
	if multiVolume && avProtected {
		status = "  (multiple volume archive + AV-secured/locked)"
	} else if multiVolume {
		status = "  (multiple volume archive)"
	} else if avProtected {
		status = "  (AV-secured/locked)"
	}

	// Get packer name
	name, ok := packerNames[arcType]
	if !ok {
		name = "unknown"
	}

	fmt.Printf("Archive type = %s%s (%d)\n", name, status, arcType)
	os.Exit(arcType)
}

func printHelp() {
	fmt.Println("IDArc V3.16 (Go Port) - Archive Identification Tool")
	fmt.Println("Identifies", nIDPacker-2, "types of archive files and returns a corresponding exit code:")
	fmt.Println()
	fmt.Println("  0. ---          1. ARC             2. ZIP                  3. ZOO")
	fmt.Println("  4. LZH          5. DWC             6. MDCD                 7. LBR")
	fmt.Println("  8. ARJ          9. HYP            10. UC2                 11. HAP")
	fmt.Println(" 12. HA          13. HPack (HPK)    14. SQZ (Squeeze It)    15. RAR")
	fmt.Println(" 16. PAK         17. ARC+           18. LIM                 19. BSN/BSA")
	fmt.Println(" 20. PUT         21. SQWEZ          22. Crush/ZIP           23. Crush/ARJ")
	fmt.Println(" 24. Crush/LZH   25. Crush/ZOO      26. Crush/HA            27. LZExe")
	fmt.Println(" 28. PKLite      29. Diet           30. TinyProg            31. GIF")
	fmt.Println(" 32. JPG (JFIF)  33. JPG (HSI)      34. AIN                 35. AINEXE")
	fmt.Println(" 36. SAR         37. BS2/BSArc      38. GZIP/Comp 4.3       39. ACB")
	fmt.Println(" 40. MAR         41. CPShrink       42. JRC                 43. JArcs")
	fmt.Println(" 44. Quantum     45. ReSOF          46. Crush/uncompressed  47. ARX")
	fmt.Println(" 48. UCEXE       49. WWPack         50. QuArk               51. YAC")
	fmt.Println(" 52. X1          53. Codec          54. AMGC                55. NuLIB")
	fmt.Println(" 56. PAKLeo      57. TGZ            58. WWPack-Data         59. ChArc")
	fmt.Println(" 60. PSA         61. ZAR            62. LHark               63. CrossePAC")
	fmt.Println(" 64. Freeze      65. KBoom          66. NSQ                 67. DPA")
	fmt.Println(" 68. TTComp      69. WIC (Fake!)    70. RKive               71. JAR")
	fmt.Println(" 72. ESP         73. ZPack          74. DRY                 75. OWS (Fake!)")
	fmt.Println(" 76. SKY         77. ARI            78. UFA                 79. Microsoft CAB")
	fmt.Println(" 80. FOXSQZ      81. AR7            82. TSComp              83. PPMZ")
	fmt.Println(" 84. MS Compress 85. MP3 (M.Czudej) 86. ZET                 87. XPack Data")
	fmt.Println(" 88. Xpk DiskImg 89. ARQ            90. ACE                 91. Squash")
	fmt.Println(" 92. Terse       93. Xpk SData      94. Stuffit (Mac)       95. PUCrunch")
	fmt.Println(" 96. BZip        97. UHarc          98. ABComp              99. CMP")
	fmt.Println("100. BZip2      101. LZOP          102. szip               103. Splint")
	fmt.Println("104. TAR        105. InstallShield 106. CARComp            107. LZS")
	fmt.Println("108. BOA        109. InstallSh. Z  110. ARG                111. Gather")
	fmt.Println("112. Pack Magic 113. BTS           114. ELI 5750           115. QFC")
	fmt.Println("116. PRO-PACK   117. MSXiE         118. RAX                119. 777")
	fmt.Println("120. LZS221     121. HPA           122. Arhangel           123. EXP1")
	fmt.Println("124. IMP        125. BMF           126. NRV                127. oPAQue")
	fmt.Println("128. Squish     129. Par           130. HIT (B. Ureche)    131. SBX")
	fmt.Println("132. NaShrinK   133. Disintegrator 134. ASD (T. Svensson)  135. InStallSh. CAB")
	fmt.Println("136. TOP4       137. BatComp (4DOS)138. BlakHole           139. BIX (I. Pavlov)")
	fmt.Println("140. ChiefLZA   141. Blink (D.T.S.)142. CAR (MylesHi!)     143. SARJ")
	fmt.Println("144. CompackSfx 145. LGExpand      146. ARS-Sfx            147. Akt")
	fmt.Println("148. Flash      149. PC/3270       150. NPack              151. PFT")
	fmt.Println("152. Xtreme     153. SemOne        154. Akt32              155. InstallIt")
	fmt.Println("156. PPMD       157. Swag          158. FIZ                159. BA")
	fmt.Println("160. XPA32      161. RK            162. RPM                163. DeepFreezer")
	fmt.Println("164. ZZip       165. DC (E.Binder) 166. TPac (Tim Gordon)  167. Ai (E.Ilya)")
	fmt.Println("168. Ybs        169. Ai32 (E.Ilya) 170. SBC (Sami Mäkinen) 171. DitPack")
	fmt.Println("172. DMS        173. EPC           174. VSARC              175. PDZ")
	fmt.Println("176. PfW        177. NullSoft Inst.178. Wise Installer     179. DZip (N.Pflug)")
	fmt.Println("180. 7z         181. ReDuq         182. aPackage           183. WinImage")
	fmt.Println("184. GCA        185. PPMN          186. SAPCAR             187. Compressia")
	fmt.Println("188. UHBC       189. PKZip6/BZip2")
	fmt.Println()
	fmt.Println("251. unknown    255. not found")
}
