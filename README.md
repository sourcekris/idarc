# Über IDArc

### What is this?

A compressed file (archive) identification tool from the 90's and early 2000's era
which can identify ~190 different archive types. Some of which remain unidentifiable 
by tools like `file` and `trid` even in 2022. Probably because those file formats 
are no longer relevant.

This repo contains:

- The original release package. 
    - A self extracting DOS archive executable for the IDArc 2.15.01 English distribution 
      containing:
        1. The original DOS release (DOS executable)
        2. Pascal source code (unmodified, original)
        3. IDArc documentation.
- Modified Pascal source code.
    - Tested to build on FreePascal on Linux in 2022.
- Makefile
    - Simplify the Pascal build, and release process.

### Building

This tool surprisingly builds with just minor modifications from the original Pascal
source code with [Free Pascal](https://www.freepascal.org/). Those modifications are already
in the repo *.pas files. So on Debian, Ubuntu derived platforms you can just:

```
$ sudo apt install fpc git
$ git clone https://github.com/sourcekris/idarc
$ make
```

A binary release for Linux amd64 is found in the releases section of this repo.

### Use

The resulting binary (`idarc`) is simple to use, provide an archive as the first
argument, it will produce a return code in line with the following table:

| Return Code | Archiver Type |
| ----------- | ------------- |
| 1   | ARC |
| 2   | ZIP |
| 3   | ZOO |
| 4   | LZH |
| 5   | DWC |
| 6   | MDCD |
| 7   | LBR |
| 8   | ARJ |
| 9   | HYP |
| 10  | UC2 |
| 11  | HAP |
| 12  | HA |
| 13  | HPack |
| 14  | SQZ (Squeeze It) |
| 15  | RAR |
| 16  | PAK |
| 17  | ARC+ |
| 18  | LIM |
| 19  | BSN (BSA/PTS-DOS) |
| 20  | PUT |
| 21  | SQWEZ |
| 22  | Crush/ZIP |
| 23  | Crush/ARJ |
| 24  | Crush/LZH |
| 25  | Crush/ZOO |
| 26  | Crush/HA |
| 27  | LZExe |
| 28  | PKLite |
| 29  | Diet |
| 30  | TinyProg |
| 31  | GIF |
| 32  | JPG (JFIF) |
| 33  | JPG (HSI) |
| 34  | AIN |
| 35  | AINEXE |
| 36  | SAR |
| 37  | BS2/BSArc |
| 38  | GZIP/Comp 4.3 |
| 39  | ACB |
| 40  | MAR |
| 41  | CPZ (CPShrink) |
| 42  | JRC |
| 43  | JArcs |
| 44  | Quantum |
| 45  | ReSOF |
| 46  | Crush/uncomp. |
| 47  | ARX |
| 48  | UCEXE |
| 49  | WWPack |
| 50  | QuArk |
| 51  | YAC |
| 52  | X1 |
| 53  | Codec |
| 54  | AMGC |
| 55  | NuLIB |
| 56  | PAKLeo (PLL) |
| 57  | TGZ |
| 58  | WWPack data file |
| 59  | ChArc |
| 60  | PSA |
| 61  | ZAR |
| 62  | LHark |
| 63  | CrossePAC |
| 64  | Freeze |
| 65  | KBoom |
| 66  | NSQ |
| 67  | DPA |
| 68  | TTComp |
| 69  | WIC (Fake!) |
| 70  | RKive |
| 71  | JAR |
| 72  | ESP |
| 73  | ZPack |
| 74  | DRY |
| 75  | OWS (Fake!) |
| 76  | SKY |
| 77  | ARI |
| 78  | UFA |
| 79  | Microsoft CAB |
| 80  | FOXSQZ |
| 81  | AR7 |
| 82  | TSComp |
| 83  | PPMZ |
| 84  | MS Compress |
| 85  | MP3 (Marco Czudej) |
| 86  | ZET |
| 87  | XPack Data |
| 88  | XPack Diskimage |
| 89  | ARQ |
| 90  | ACE |
| 91  | Squash |
| 92  | Terse |
| 93  | XPack single data |
| 94  | Stuffit (Mac) |
| 95  | PUCrunch |
| 96  | BZip |
| 97  | UHarc |
| 98  | ABComp |
| 99  | CMP (Andr� Olejko) |
| 100 | BZip2 |
| 101 | LZO |
| 102 | szip |
| 103 | Splint |
| 104 | TAR |
| 105 | InstallShield |
| 106 | CARComp |
| 107 | LZS |
| 108 | BOA |
| 109 | InstallShield Z |
| 110 | ARG |
| 111 | Gather |
| 112 | Pack Magic |
| 113 | BTS |
| 114 | ELI 5750 |
| 115 | QFC |
| 116 | PRO-PACK |
| 117 | MSXiE |
| 118 | RAX |
| 119 | 777 |
| 120 | LZS221 |
| 121 | HPA |
| 122 | Arhangel |
| 123 | EXP1 |
| 124 | IMP |
| 125 | BMF |
| 126 | NRV |
| 127 | oPAQue |
| 128 | Squish (Mike Albert) |
| 129 | Par |
| 130 | HIT (Bogdan Ureche) |
| 131 | SBX |
| 132 | NaShrink |
| 133 | Disintegrator |
| 134 | ASD |
| 135 | InstallShield CAB |
| 136 | TOP4 |
| 137 | BatComp (4DOS) |
| 138 | BlakHole |
| 139 | BIX (Igor Pavlov) |
| 140 | ChiefLZA |
| 141 | Blink (D.T.S.) |
| 142 | CAR (MylesHi!) |
| 143 | SARJ |
| 144 | Compack Sfx |
| 145 | Logitech Compress |
| 146 | ARS-Sfx |
| 147 | AKT |
| 148 | Flash |
| 149 | PC/3270 |
| 150 | NPack |
| 151 | PFT |
| 152 | Xtreme |
| 153 | SemOne |
| 154 | AKT32 |
| 155 | InstallIt |
| 156 | PPMD |
| 157 | Swag |
| 158 | FIZ |
| 159 | BA (M. Lundqvist) |
| 160 | XPA32 (J. Tseng) |
| 161 | RK (M.Taylor) |
| 162 | RPM |
| 163 | DeepFreezer |
| 164 | ZZip (Damien Debin) |
| 165 | DC (Edgar Binder) |
| 166 | TPac (Tim Gordon) |
| 167 | Ai (E.Ilya) |
| 168 | Ybs (Vadim Yoockin) |
| 169 | Ai32 (E.Ilya) |
| 170 | SBC (Sami M�kinen) |
| 171 | DitPack |
| 172 | DMS |
| 173 | EPC |
| 174 | VSARC |
| 175 | PDZ |
| 176 | Package for the Web |
| 177 | NullSoft Installer |
| 178 | Wise Installer |
| 179 | DZip (Nolan Pflug) |
| 180 | 7z |
| 181 | ReDuq (J. Mintjes) |
| 182 | aPackage |
| 183 | WinImage |
| 184 | GCA |
| 185 | PPMN (Max Smirnov) |
| 186 | SAPCAR |
| 187 | Compressia |
| 188 | UHBC |
| 189 | PKZip6/BZip2 |

### Original License

> IDArc ist ein einfaches Tool zum Identifizieren von komprimierten Dateien 189
> verschiedener Formate. Da es keine große programmiertechnische Leistung dar-
> stellt und ich auch keinerlei Weiterentwicklung oder sonstige Eigenschaften
> garantiere, außer dass es Platz auf der Festplatte belegt, gebe ich es hiermit
> als
> 			    FREEWARE

> frei. Das Copyright verbleibt jedoch bei mir und das Programm darf weder
> verändert noch disassembliert o.�. werden. Eine Weitergabe ist jederzeit
> erlaubt, wenn sie _entgeltlos_ erfolgt (von Kopiergebühren der Shareware-
> Händler mal abgesehen). Der Autor übernimmt KEINERLEI Haftung!

> Das Programm verwendet die beiliegende Pascal-Unit IDPACKER.PAS. Wen es inter-
> essiert, der kann darin einige Infos zum Identifizieren von Archiven finden.

### Original Author

Jürgen Peters

See `idarc.txt` for original documentation.
