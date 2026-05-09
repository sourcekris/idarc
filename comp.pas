{ $Id: comp.pas 1.6 2001/08/01 19:00:06 rar_ulm Exp $ }
{ --- COMPILER - Directives ------------------------------------------- }
{ Multi Platform Compiler settings (c) by ROSE SWE, Ralph Roth		}
{ Works for Turbo Pascal, Free Pascal, Virtual Pascal, GPC, TMT etc.	}
{ Contact:   http://come.to/rose_swe, rose_swe@hotmail.com		}
{ --------------------------------------------------------------------- }

(*
$Log: comp.pas $
Revision 1.6  2001/08/01 19:00:06  rar_ulm
Tried to fixed VP 2.1 bug with PE-Head (AlignRec etc.)

Revision 1.5  2001/04/07 13:38:38  rar_ulm
VP 2.1, Send to Mathias

Revision 1.4  2001/03/19 08:16:26  rar
Stack checking under Linux removed

Revision 1.3  2001/03/18 12:37:28  rar
TMT Pascal 3.50 support added

Revision 1.2  2001/03/11 11:49:31  rose300
First RCS version: Common generic file for TP6, FPC & GNU Pascal
*)

{ --------------------------------------------------------------------- }
{ FPC 1.00 Support added 17-Aug-2000, rar                               }
{ Aenderung am 08.01.95 - G+ Option, Stack off!                         }
{ Aenderungen am 12.10.92 fuer TP 6.0 - CoProz                          }
{ Aenderung am 1.7.91 fuer TP 6.0                                       }
{ 28/5/90 - COMP.PAS - For Turbo Pascal v5.0                            }
{ --------------------------------------------------------------------- }

{$IFDEF __TMT__}                (* fix 17.03.2001 for TMT 3.50d *)
{$DEFINE FPC}
{$ENDIF}

{$IFDEF VIRTUALPASCAL}          (* fix 29.03.2001 for Virtual Pascal 2.1 *)
{$DEFINE FPC}
{&G5+}
{&SmartLink+}
{&USE32+}               { 16/32 bit for VPC }
{&AlignRec-}            { 01.08.2001, rar }
{$ENDIF}

{$IFNDEF GNU_PASCAL}
{$R-}   { Array Range Checking }
{$IFNDEF LINUX}
{$S-}   { Stack Checking on, sonst eventuell Absturz }
{$ENDIF}
{$I-}   { I/O Checking }
{$V-}   { Variablen Checking }

{$IFNDEF FPC}
{$B-}   { Boolean Checking }
{$F+}   { Far Calls On, Errorhandling, Overlay Manager }
{$O+}   { Overlay allowed }
{.N-}   { CoProzessor }
{$D-}   { Debug Info }
{$L-}   { Local Symbols }
{.E-}   { 8087 Emulation falls kein Coprozessor vorhanden, verbraucht }
        { aber ernorm Speicherplatz (bigger EXE file) }
{$A+}   { Alignment checkin' - schneller fuer alle 16Bit-Rechner}

{$G+}   { 80286 code }

{$ENDIF}
{$ENDIF}        { nothing for GNU PASCAL! }

{ Ersparen bis zu 30% Platz und bringen einen Geschwindigskeits-
  zuwachs von etwa 100% ! }

