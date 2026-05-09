(* Beispielprogramm fuer die Anwendung der Unit IDPACKER.PAS *)
(* Sample program for the use of IDPACKER.PAS                *)

(* fixed and ported 8-Aug-2001 by Ralph Roth, ROSE SWE to FPC 1.0.4/32 bit *)
(* doesn't compile under VPC (uses "registers")                            *)

// Modified in 2022 to build on modern Linux with Free Pascal.
{$IFDEF WOLNA}
{$N+,E+}
{$ENDIF}
{$IFNDEF DEBUG}
{$I comp.pas}
{$D-,L-}
{$ENDIF}

{&R c:\tp\icons\tools.ico}		{ resource for vpc, rar }

USES DOS,IDPacker{$IFDEF LONGNAME},LFN{$ENDIF};

CONST Version='2.15.01';
      Jahr='2022';
      {$IFDEF FPC}      Ausgabeumgeleitet = FALSE;      {$ENDIF}

VAR   ArcTyp: Byte;
      Archiv,s: String;
      EL: String[3];

PROCEDURE Hilfe;
 BEGIN
  WriteLn('IDArc V',Version,' (C) ',Jahr,' Jürgen Peters');
  {$IFNDEF ENGLISH}
  WriteLn('Syntax: IDArc archivdateiname');
  WriteLn('Identifiziert ',nIDPacker-2,' Archivdateitypen und gibt entsprechenden ERRORLEVEL aus:');
  {$ELSE}
  WriteLn('Syntax: IDArc archive_name');
  WriteLn('Identifies ',nIDPacker-2,' types of archive files and returns a corresponding ERRORLEVEL:');
  {$ENDIF}
  WriteLn('  0. ---          1. ARC             2. ZIP                  3. ZOO');
  WriteLn('  4. LZH          5. DWC             6. MDCD                 7. LBR');
  WriteLn('  8. ARJ          9. HYP            10. UC2                 11. HAP');
  WriteLn(' 12. HA          13. HPack (HPK)    14. SQZ (Squeeze It)    15. RAR');
  WriteLn(' 16. PAK         17. ARC+           18. LIM                 19. BSN/BSA');
  WriteLn(' 20. PUT         21. SQWEZ          22. Crush/ZIP           23. Crush/ARJ');
  WriteLn(' 24. Crush/LZH   25. Crush/ZOO      26. Crush/HA            27. LZExe');
  WriteLn(' 28. PKLite      29. Diet           30. TinyProg            31. GIF');
  WriteLn(' 32. JPG (JFIF)  33. JPG (HSI)      34. AIN                 35. AINEXE');
  WriteLn(' 36. SAR         37. BS2/BSArc      38. GZIP/Comp 4.3       39. ACB');
  WriteLn(' 40. MAR         41. CPShrink       42. JRC                 43. JArcs');
  {$IFNDEF ENGLISH}
  WriteLn(' 44. Quantum     45. ReSOF          46. Crush/ungepackt     47. ARX');
  {$ELSE}
  WriteLn(' 44. Quantum     45. ReSOF          46. Crush/uncompressed  47. ARX');
  {$ENDIF}
  WriteLn(' 48. UCEXE       49. WWPack         50. QuArk               51. YAC');
  WriteLn(' 52. X1          53. Codec          54. AMGC                55. NuLIB');
  WriteLn(' 56. PAKLeo      57. TGZ            58. WWPack-Data         59. ChArc');
  WriteLn(' 60. PSA         61. ZAR            62. LHark               63. CrossePAC');
  WriteLn(' 64. Freeze      65. KBoom          66. NSQ                 67. DPA');
  WriteLn(' 68. TTComp      69. WIC (Fake!)    70. RKive               71. JAR');
  WriteLn(' 72. ESP         73. ZPack          74. DRY                 75. OWS (Fake!)');
  WriteLn(' 76. SKY         77. ARI            78. UFA                 79. Microsoft CAB');
  WriteLn(' 80. FOXSQZ      81. AR7            82. TSComp              83. PPMZ');
  WriteLn(' 84. MS Compress 85. MP3 (M.Czudej) 86. ZET                 87. XPack Data');
  WriteLn(' 88. Xpk DiskImg 89. ARQ            90. ACE                 91. Squash');
  WriteLn(' 92. Terse       93. Xpk SData      94. Stuffit (Mac)       95. PUCrunch');
  WriteLn(' 96. BZip        97. UHarc          98. ABComp              99. CMP');
  WriteLn('100. BZip2      101. LZOP          102. szip               103. Splint');
  WriteLn('104. TAR        105. InstallShield 106. CARComp            107. LZS');
  WriteLn('108. BOA        109. InstallSh. Z  110. ARG                111. Gather');
  WriteLn('112. Pack Magic 113. BTS           114. ELI 5750           115. QFC');
  WriteLn('116. PRO-PACK   117. MSXiE         118. RAX                119. 777');
  WriteLn('120. LZS221     121. HPA           122. Arhangel           123. EXP1');
  WriteLn('124. IMP        125. BMF           126. NRV                127. oPAQue');
  WriteLn('128. Squish     129. Par           130. HIT (B. Ureche)    131. SBX');
  WriteLn('132. NaShrinK   133. Disintegrator 134. ASD (T. Svensson)  135. InStallSh. CAB');
  WriteLn('136. TOP4       137. BatComp (4DOS)138. BlakHole           139. BIX (I. Pavlov)');
  WriteLn('140. ChiefLZA   141. Blink (D.T.S.)142. CAR (MylesHi!)     143. SARJ');
  WriteLn('144. CompackSfx 145. LGExpand      146. ARS-Sfx            147. Akt');
  WriteLn('148. Flash      149. PC/3270       150. NPack              151. PFT');
  WriteLn('152. Xtreme     153. SemOne        154. Akt32              155. InstallIt');
  WriteLn('156. PPMD       157. Swag          158. FIZ                159. BA');
  WriteLn('160. XPA32      161. RK            162. RPM                163. DeepFreezer');
  WriteLn('164. ZZip       165. DC (E.Binder) 166. TPac (Tim Gordon)  167. Ai (E.Ilya)');
  WriteLn('168. Ybs        169. Ai32 (E.Ilya) 170. SBC (Sami M�kinen) 171. DitPack');
  WriteLn('172. DMS        173. EPC           174. VSARC              175. PDZ');
  WriteLn('176. PfW        177. NullSoft Inst.178. Wise Installer     179. DZip (N.Pflug)');
  WriteLn('180. 7z         181. ReDuq         182. aPackage           183. WinImage');
  WriteLn('184. GCA        185. PPMN          186. SAPCAR             187. Compressia');
  WriteLn('188. UHBC       189. PKZip6/BZip2');
  {$IFNDEF ENGLISH}
  WriteLn('251. unbekannt  255. n.gefunden');
  {$ELSE}
  WriteLn('251. unknown    255. not found');
  {$ENDIF}
 END;

BEGIN
 If ParamCount=0 then Hilfe ELSE
  BEGIN
   FileMode := 0 { $80 };
   Archiv := ParamStr(1);
   ArcTyp := ArchiveType(Archiv);
   Str(ArcTyp,EL);
   s := '';
   AssignPackerNames;
   {$IFNDEF ENGLISH}
   If (mv and av) then s := '  (Multiple Volume-Archiv + AV-gesch�tzt/locked)' ELSE
   If mv then s := '  (Multiple Volume-Archiv)' ELSE
   If av then s := '  (AV-gesch�tzt/locked)';
   WriteLn('Archiv-Typ = '+PackerNames^[ArcTyp]+s+' ('+EL+')');
   {$ELSE}
   If (mv and av) then s := '  (multiple volume archive + AV-secured/locked)' ELSE
   If mv then s := '  (multiple volume archive)' ELSE
   If av then s := '  (AV-secured/locked)';
   WriteLn('Archive type = '+PackerNames^[ArcTyp]+s+' ('+EL+')');
   {$ENDIF}
   DisposePackerNames;
   Halt(ArcTyp);
  END;
END.
