FPC=fpc
PFLAGS=-dENGLISH

build: comp.pas idpacker.pas idarc.pas
	$(FPC) $(PFLAGS) idarc.pas

release: clean build
	zip idarc.zip idarc comp.pas idpacker.pas idarc.pas idarc.txt idar215e.exe Makefile

clean:
	rm -f *.o *.ppu idarc