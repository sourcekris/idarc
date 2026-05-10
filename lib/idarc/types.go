package idarc

type ArchiveType int

const (
	ARC ArchiveType = 1
	ZIP ArchiveType = 2
	ZOO ArchiveType = 3
	LZH ArchiveType = 4
	DWC ArchiveType = 5
	MDCD ArchiveType = 6
	LBR ArchiveType = 7
	ARJ ArchiveType = 8
	HYP ArchiveType = 9
	UC2 ArchiveType = 10
	HAP ArchiveType = 11
	HA ArchiveType = 12
	HPack ArchiveType = 13
	SQZ ArchiveType = 14
	RAR ArchiveType = 15
	PAK ArchiveType = 16
	ARCPlus ArchiveType = 17
	LIM ArchiveType = 18
	BSN ArchiveType = 19
	PUT ArchiveType = 20
	SQWEZ ArchiveType = 21
	CrushZIP ArchiveType = 22
	CrushARJ ArchiveType = 23
	CrushLZH ArchiveType = 24
	CrushZOO ArchiveType = 25
	CrushHA ArchiveType = 26
	LZExe ArchiveType = 27
	PKLite ArchiveType = 28
	Diet ArchiveType = 29
	TinyProg ArchiveType = 30
	GIF ArchiveType = 31
	JFIF ArchiveType = 32
	HSI ArchiveType = 33
	AIN ArchiveType = 34
	AINEXE ArchiveType = 35
	SAR ArchiveType = 36
	BS2 ArchiveType = 37
	GZIP ArchiveType = 38
	ACB ArchiveType = 39
	MAR ArchiveType = 40
	CPZ ArchiveType = 41
	JRC ArchiveType = 42
	JArcs ArchiveType = 43
	Quantum ArchiveType = 44
	Sof ArchiveType = 45
	CrushUncomp ArchiveType = 46
	ARX ArchiveType = 47
	UCEXE ArchiveType = 48
	WWPack ArchiveType = 49
	QuArk ArchiveType = 50
	YAC ArchiveType = 51
	X1 ArchiveType = 52
	Codec ArchiveType = 53
	AMGC ArchiveType = 54
	NuLIB ArchiveType = 55
	PAKLeo ArchiveType = 56
	TGZ ArchiveType = 57
	WWPackData ArchiveType = 58
	ChArc ArchiveType = 59
	PSA ArchiveType = 60
	ZAR ArchiveType = 61
	LHark ArchiveType = 62
	CrossePAC ArchiveType = 63
	Freeze ArchiveType = 64
	KBoom ArchiveType = 65
	NSQ ArchiveType = 66
	DPA ArchiveType = 67
	TTComp ArchiveType = 68
	WIC ArchiveType = 69
	RKive ArchiveType = 70
	JAR ArchiveType = 71
	ESP ArchiveType = 72
	ZPack ArchiveType = 73
	DRY ArchiveType = 74
	OWS ArchiveType = 75
	SKY ArchiveType = 76
	ARI ArchiveType = 77
	UFA ArchiveType = 78
	CAB ArchiveType = 79
	FOXSQZ ArchiveType = 80
	AR7 ArchiveType = 81
	TSComp ArchiveType = 82
	PPMZ ArchiveType = 83
	MSCompress ArchiveType = 84
	MP3 ArchiveType = 85
	ZET ArchiveType = 86
	XPackData ArchiveType = 87
	XPackDiskimage ArchiveType = 88
	ARQ ArchiveType = 89
	ACE ArchiveType = 90
	Squash ArchiveType = 91
	Terse ArchiveType = 92
	XPackSingleData ArchiveType = 93
	Stuffit ArchiveType = 94
	PUCrunch ArchiveType = 95
	BZip ArchiveType = 96
	UHarc ArchiveType = 97
	ABComp ArchiveType = 98
	CMP ArchiveType = 99
	BZip2 ArchiveType = 100
	LZO ArchiveType = 101
	Szip ArchiveType = 102
	Splint ArchiveType = 103
	TAR ArchiveType = 104
	InstallShield ArchiveType = 105
	CARComp ArchiveType = 106
	LZS ArchiveType = 107
	BOA ArchiveType = 108
	InstallShieldZ ArchiveType = 109
	ARG ArchiveType = 110
	Gather ArchiveType = 111
	PackMagic ArchiveType = 112
	BTS ArchiveType = 113
	ELI ArchiveType = 114
	QFC ArchiveType = 115
	ProPack ArchiveType = 116
	MSXiE ArchiveType = 117
	RAX ArchiveType = 118
	_777 ArchiveType = 119
	LZS221 ArchiveType = 120
	HPA ArchiveType = 121
	Arhangel ArchiveType = 122
	EXP1 ArchiveType = 123
	IMP ArchiveType = 124
	BMF ArchiveType = 125
	NRV ArchiveType = 126
	oPAQue ArchiveType = 127
	Sq ArchiveType = 128
	Par ArchiveType = 129
	HIT ArchiveType = 130
	SBX ArchiveType = 131
	NaShrink ArchiveType = 132
	Disintegrator ArchiveType = 133
	ASD ArchiveType = 134
	InstallShieldCAB ArchiveType = 135
	TOP4 ArchiveType = 136
	BatComp ArchiveType = 137
	BlakHole ArchiveType = 138
	BIX ArchiveType = 139
	ChiefLZA ArchiveType = 140
	Blink ArchiveType = 141
	CAR ArchiveType = 142
	SARJ ArchiveType = 143
	CompackSfx ArchiveType = 144
	LogitechCompress ArchiveType = 145
	ARSSfx ArchiveType = 146
	AKT ArchiveType = 147
	Flash ArchiveType = 148
	PC3270 ArchiveType = 149
	NPack ArchiveType = 150
	PFT ArchiveType = 151
	Xtreme ArchiveType = 152
	SemOne ArchiveType = 153
	AKT32 ArchiveType = 154
	InstallIt ArchiveType = 155
	PPMD ArchiveType = 156
	Swag ArchiveType = 157
	FIZ ArchiveType = 158
	BA ArchiveType = 159
	XPA32 ArchiveType = 160
	RK ArchiveType = 161
	RPM ArchiveType = 162
	DeepFreezer ArchiveType = 163
	ZZip ArchiveType = 164
	DC ArchiveType = 165
	TPac ArchiveType = 166
	Ai ArchiveType = 167
	Ybs ArchiveType = 168
	Ai32 ArchiveType = 169
	SBC ArchiveType = 170
	DitPack ArchiveType = 171
	DMS ArchiveType = 172
	EPC ArchiveType = 173
	VSARC ArchiveType = 174
	PDZ ArchiveType = 175
	PackageForTheWeb ArchiveType = 176
	NullSoftInstaller ArchiveType = 177
	WiseInstaller ArchiveType = 178
	DZip ArchiveType = 179
	_7z ArchiveType = 180
	ReDuq ArchiveType = 181
	aPackage ArchiveType = 182
	WinImage ArchiveType = 183
	GCA ArchiveType = 184
	PPMN ArchiveType = 185
	SAPCAR ArchiveType = 186
	Compressia ArchiveType = 187
	UHBC ArchiveType = 188
	PKZip6BZip2 ArchiveType = 189
	MeltingPotArchiver ArchiveType = 190
	MARUtilityFile ArchiveType = 191
	MARUtilityFolder ArchiveType = 192

	Invalid      ArchiveType = 251
	FileNotFound ArchiveType = 255
)

var PackerNames = map[ArchiveType]string{
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
	45:  "Sof",
	46:  "Crush/uncomp.",
	47:  "ARX",
	48:  "UCEXE",
	49:  "WWPack",
	50:  "QuArk",
	51:  "YAC",
	52:  "X1",
	53:  "Codec",
	54:  "AMGC",
	55:  "NuLIB",
	56:  "PAKLeo (PLL)",
	57:  "TGZ",
	58:  "WWPack data file",
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
	85:  "MP3 (Marco Czudej)",
	86:  "ZET",
	87:  "XPack Data",
	88:  "XPack Diskimage",
	89:  "ARQ",
	90:  "ACE",
	91:  "Squash",
	92:  "Terse",
	93:  "XPack single data",
	94:  "Stuffit (Mac)",
	95:  "PUCrunch",
	96:  "BZip",
	97:  "UHarc",
	98:  "ABComp",
	99:  "CMP (André Olejko)",
	100: "BZip2",
	101: "LZO",
	102: "szip",
	103: "Splint",
	104: "TAR",
	105: "InstallShield",
	106: "CARComp",
	107: "LZS",
	108: "BOA",
	109: "InstallShield Z",
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
	128: "Sq (Mike Albert)",
	129: "Par",
	130: "HIT (Bogdan Ureche)",
	131: "SBX",
	132: "NaShrink",
	133: "Disintegrator",
	134: "ASD",
	135: "InstallShield CAB",
	136: "TOP4",
	137: "BatComp (4DOS)",
	138: "BlakHole",
	139: "BIX (Igor Pavlov)",
	140: "ChiefLZA",
	141: "Blink (D.T.S.)",
	142: "CAR (MylesHi!)",
	143: "SARJ",
	144: "Compack Sfx",
	145: "Logitech Compress",
	146: "ARS-Sfx",
	147: "AKT",
	148: "Flash",
	149: "PC/3270",
	150: "NPack",
	151: "PFT",
	152: "Xtreme",
	153: "SemOne",
	154: "AKT32",
	155: "InstallIt",
	156: "PPMD",
	157: "Swag",
	158: "FIZ",
	159: "BA (M. Lundqvist)",
	160: "XPA32 (J. Tseng)",
	161: "RK (M.Taylor)",
	162: "RPM",
	163: "DeepFreezer",
	164: "ZZip (Damien Debin)",
	165: "DC (Edgar Binder)",
	166: "TPac (Tim Gordon)",
	167: "Ai (E.Ilya)",
	168: "Ybs (Vadim Yoockin)",
	169: "Ai32 (E.Ilya)",
	170: "SBC (Sami Mäkinen)",
	171: "DitPack",
	172: "DMS",
	173: "EPC",
	174: "VSARC",
	175: "PDZ",
	176: "Package for the Web",
	177: "NullSoft Installer",
	178: "Wise Installer",
	179: "DZip (Nolan Pflug)",
	180: "7z",
	181: "ReDuq (J. Mintjes)",
	182: "aPackage",
	183: "WinImage",
	184: "GCA",
	185: "PPMN (Max Smirnov)",
	186: "SAPCAR",
	187: "Compressia",
	188: "UHBC",
	189: "PKZip6/BZip2",
	190: "MeltingPot Archiver",
	191: "MAR Archive / MAR Utility (file)",
	192: "MAR Archive / MAR Utility (folder)",
	251: "unknown",
	255: "File not found",
}
