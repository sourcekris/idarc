# IDArc Go Port

Go port of the IDArc utility for identifying 189+ different types of compressed and archived files.

## Overview

This is a modern Go implementation of the classic IDArc archive identification tool. The original Turbo Pascal version identified archive formats by their magic bytes and file signatures.

## Building

```bash
go build -o idarc
```

## Usage

```bash
./idarc <archive_file>
```

### Examples

```bash
./idarc archive.zip
./idarc backup.rar
./idarc data.tar.gz
```

## Output

The tool returns:
- Archive type name (e.g., "ZIP", "RAR", "7z")
- Status flags for multiple volume archives and AV-protected files
- Exit code matching the archive type number

```
Archive type = ZIP (2)
Archive type = RAR  (multiple volume archive) (15)
Archive type = 7z (180)
```

## Supported Archive Types

Identifies 189+ formats including:
- **Common**: ZIP, RAR, 7z, TAR, GZIP, BZIP2
- **Legacy**: ARC, ARJ, LZH, ZOO, HAP
- **Specialty**: CAB, RPM, DMS, ACE
- **Executables**: PKLite, Diet, UPX, etc.
- **Images**: GIF, JPEG, WinImage

## Exit Codes

Exit codes correspond to archive type:
- `0` - Unknown/invalid archive
- `1-189` - Specific archive type (see help output)
- `251` - Unknown archive type
- `255` - File not found

Run `idarc --help` or `idarc` with no arguments to see the complete list.

## Original Source

Based on IDArc v3.15.01 by Jürgen Peters. Original source files from the `main` branch.

## Notes

This Go port maintains compatibility with the original's behavior:
- Same archive type numbering system
- Same exit codes for scripting
- Detects multiple volume archives
- Detects AV-protected/locked archives
- Same help output and format

The implementation prioritizes clean, idiomatic Go code while preserving the original's functionality.
