package idarc

import "bytes"

type Signature struct {
	Pattern []byte
	Offset  int // 1-indexed, 0 means anywhere
	Type    ArchiveType
}

var signatures = []Signature{
	{Pattern: []byte("HLSQZ"), Offset: 1, Type: SQZ},
	{Pattern: []byte("SQWEZ"), Offset: 1, Type: SQWEZ},
	{Pattern: []byte("HPAK"), Offset: 1, Type: HPack},
	{Pattern: []byte("\x13HF"), Offset: 1, Type: HAP},
	{Pattern: []byte("\xdc\xa7\xc4\xfd"), Offset: 21, Type: ZOO},
	{Pattern: []byte("HA"), Offset: 1, Type: HA},
	{Pattern: []byte("MDmd"), Offset: 1, Type: MDCD},
	{Pattern: []byte("LM\x1a\x08\x00"), Offset: 1, Type: LIM},
	{Pattern: []byte("LM\x1a\x07\x00"), Offset: 1, Type: LIM},
	{Pattern: []byte("LH5"), Offset: 4, Type: SAR},
	{Pattern: []byte("\xd4\x03SB \x00"), Offset: 1, Type: BS2},
	{Pattern: []byte("-ah"), Offset: 3, Type: MAR},
	{Pattern: []byte("JRchive"), Offset: 1, Type: JRC},
	{Pattern: []byte("JARCS"), Offset: 1, Type: JAR},
	{Pattern: []byte("DS\x00"), Offset: 1, Type: Quantum},
	{Pattern: []byte("PK\x03\x06"), Offset: 1, Type: Sof},
	{Pattern: []byte("7\x04"), Offset: 1, Type: QuArk},
	{Pattern: []byte("YC"), Offset: 15, Type: YAC},
	{Pattern: []byte("X1"), Offset: 1, Type: X1},
	{Pattern: []byte("XhDr"), Offset: 1, Type: X1},
	{Pattern: []byte("\xad6\""), Offset: 1, Type: AMGC},
	{Pattern: []byte("N\xfcF\xbcl\xe1"), Offset: 1, Type: NuLIB},
	{Pattern: []byte("LEOLZW"), Offset: 1, Type: PAKLeo},
	{Pattern: []byte("\x1f\x8b\x08"), Offset: 1, Type: TGZ},
	{Pattern: []byte("SChF"), Offset: 1, Type: ChArc},
	{Pattern: []byte("PSA"), Offset: 1, Type: PSA},
	{Pattern: []byte("DSIGDCC"), Offset: 1, Type: PAK},
	{Pattern: []byte("\x1f\x9f\x4a\x10\x0a"), Offset: 1, Type: Freeze},
	{Pattern: []byte("\xbcMP\xbe"), Offset: 1, Type: KBoom},
	{Pattern: []byte("\x76\xff"), Offset: 1, Type: NSQ},
	{Pattern: []byte("Dirk Paehl"), Offset: 1, Type: DPA},
	{Pattern: []byte("ESP"), Offset: 1, Type: ESP},
	{Pattern: []byte("\x01ZPK\x01"), Offset: 1, Type: ZPack},
	{Pattern: []byte("\xbc\x40"), Offset: 1, Type: SKY},
	{Pattern: []byte("UFA"), Offset: 1, Type: UFA},
	{Pattern: []byte("-H2O"), Offset: 1, Type: DRY},
	{Pattern: []byte("MSCF"), Offset: 1, Type: CAB},
	{Pattern: []byte("FOXSQZ"), Offset: 1, Type: FOXSQZ},
	{Pattern: []byte(",AR7"), Offset: 1, Type: AR7},
	{Pattern: []byte("PPMZ"), Offset: 1, Type: PPMZ},
	{Pattern: []byte("\x88\xf0\x27"), Offset: 5, Type: MSCompress},
	{Pattern: []byte("MP3\x1a"), Offset: 1, Type: MP3},
	{Pattern: []byte("OZ\xbe"), Offset: 1, Type: ZET},
	{Pattern: []byte("\x65\x5d\x13\x8c\x08\x01\x03\x00"), Offset: 1, Type: TSComp},
	{Pattern: []byte("gW\x04\x01"), Offset: 1, Type: ARQ},
	{Pattern: []byte("OctSqu"), Offset: 4, Type: Squash},
	{Pattern: []byte("\x05\x01\x01\x00"), Offset: 1, Type: Terse},
	{Pattern: []byte("SIT!"), Offset: 1, Type: Stuffit},
	{Pattern: []byte("\x01\x08\x0b\x08\xef\x00\x9e\x32\x30\x36\x31"), Offset: 1, Type: PUCrunch},
	{Pattern: []byte("UHA"), Offset: 1, Type: UHarc},
	{Pattern: []byte("\x02AB"), Offset: 1, Type: ABComp},
	{Pattern: []byte("\x03AB2"), Offset: 1, Type: ABComp},
	{Pattern: []byte("CO\x00"), Offset: 1, Type: CMP},
	{Pattern: []byte("\x8dLZO"), Offset: 1, Type: LZO},
	{Pattern: []byte("\x93\xb9\x06"), Offset: 1, Type: Splint},
	{Pattern: []byte("\x13\x5d\x65\x8c"), Offset: 1, Type: InstallShieldZ},
	{Pattern: []byte("GTH"), Offset: 2, Type: Gather},
	{Pattern: []byte("BOA"), Offset: 1, Type: BOA},
	{Pattern: []byte("ULEB\x0a"), Offset: 1, Type: RAX},
	{Pattern: []byte("ULEB\x00"), Offset: 1, Type: Xtreme},
	{Pattern: []byte("\x40\xbe\x01\x00"), Offset: 1, Type: PackMagic},
	{Pattern: []byte("Ora "), Offset: 1, Type: ELI},
}

func (d *Detector) matchSignatures() ArchiveType {
	for _, sig := range signatures {
		if sig.Offset > 0 {
			start := sig.Offset - 1
			end := start + len(sig.Pattern)
			if end <= len(d.IDStr) {
				if bytes.Equal(d.IDStr[start:end], sig.Pattern) {
					return sig.Type
				}
			}
		} else {
			if bytes.Contains(d.IDStr, sig.Pattern) {
				return sig.Type
			}
		}
	}
	return Invalid
}
