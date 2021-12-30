
#Note: The input file looks like:
#{AADDF, C_FREG, C_NONE, C_FREG, 32, 4, 0, 0, 0},
#{AADDF, C_FREG, C_REG, C_FREG, 32, 4, 0, 0, 0},
#{ACMPEQF, C_FREG, C_REG, C_NONE, 32, 4, 0, 0, 0},
#{AABSF, C_FREG, C_NONE, C_FREG, 33, 4, 0, 0, 0},
#{AMOVVF, C_FREG, C_NONE, C_FREG, 33, 4, 0, sys.LOONG64, 0},
#{AMOVF, C_FREG, C_NONE, C_FREG, 33, 4, 0, 0, 0},
#{AMOVD, C_FREG, C_NONE, C_FREG, 33, 4, 0, 0, 0},
#{AMOVW, C_REG, C_NONE, C_SEXT, 7, 4, REGSB, sys.LOONG64, 0},
#{AMOVWU, C_REG, C_NONE, C_SEXT, 7, 4, REGSB, sys.LOONG64, 0},

INPUT=$1

#C_SCON
sed -i 's/^.*{A\(.*\), C_SCON.*C_NONE.*C_REG.*/\1 $4, R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_SCON.*C_REG.*C_REG.*/\1 $4, R4, R5/g' $INPUT

#C_UCON
sed -i 's/^.*{A\(.*\), C_UCON.*C_NONE.*C_REG.*/\1 $0x10000, R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_UCON.*C_REG.*C_REG.*/\1 $0x10000, R4, R5/g' $INPUT

#C_ADD0CON
sed -i 's/^.*{A\(.*\), C_ADD0CON.*C_NONE.*C_REG.*/\1 $-1, R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_ADD0CON.*C_REG.*C_REG.*/\1 $-1, R4, R5/g' $INPUT

#C_ADDCON
sed -i 's/^.*{A\(.*\), C_ADDCON.*C_NONE.*C_REG.*/\1 $-1, R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_ADDCON.*C_NONE.*C_FREG.*/\1 $-1, F4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_ADDCON.*C_REG.*C_REG.*/\1 $-1, R4, R5/g' $INPUT

#C_AND0CON
sed -i 's/^.*{A\(.*\), C_AND0CON.*C_NONE.*C_REG.*/\1 $1, R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_AND0CON.*C_REG.*C_REG.*/\1 $1, R4, R5/g' $INPUT

#C_ANDCON
sed -i 's/^.*{A\(MOV.*\), C_ANDCON.*C_NONE.*C_REG.*/\1 $1, R4/g' $INPUT
sed -i 's/^.*{A\(MOV.*\), C_ANDCON.*C_NONE.*C_FREG.*/\1 $1, F4/g' $INPUT
sed -i 's/^.*{A\(MOV.*\), C_ANDCON.*C_REG.*C_REG.*/\1 $1, R4, R5/g' $INPUT

#C_LCON
sed -i 's/^.*{A\(.*\), C_LCON.*C_NONE.*C_REG.*/\1 $0x12345, R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_LCON.*C_NONE.*C_NONE.*/\1 $0x12345/g' $INPUT
sed -i 's/^.*{A\(.*\), C_LCON.*C_REG.*C_REG.*/\1 $0x12345, R4, R5/g' $INPUT

#C_SACON
sed -i 's/^.*{A\(.*\), C_SACON.*C_NONE.*C_REG.*/\1 $4(R4), R5/g' $INPUT

#C_LACON
sed -i 's/^.*{A\(.*\), C_LACON.*C_NONE.*C_REG.*/\1 $0x10000(R4), R5/g' $INPUT

#C_STCON
#TODO


#C_SOREG
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_SOREG.*/\1 R4, 1(R5)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_SOREG.*C_NONE.*C_REG.*/\1 1(R5), R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_NONE.*C_SOREG.*/\1 F4, 1(R5)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_SOREG.*C_NONE.*C_FREG.*/\1 1(R5), F4/g' $INPUT

#C_LOREG
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_LOREG.*/\1 R4, 0x10000(R5)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_LOREG.*C_NONE.*C_REG.*/\1 0x10000(R5), R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_NONE.*C_LOREG.*/\1 F4, 0x10000(R5)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_LOREG.*C_NONE.*C_FREG.*/\1 0x10000(R5), F4/g' $INPUT

#C_SAUTO
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_SAUTO.*/\1 R4, result+16(FP)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_SAUTO.*C_NONE.*C_REG.*/\1 y+8(FP), R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_NONE.*C_SAUTO.*/\1 F4, result+16(FP)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_SAUTO.*C_NONE.*C_FREG.*/\1 y+8(FP), F4/g' $INPUT

#C_LAUTO
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_LAUTO.*/\1 R4, result+0x10004(FP)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_LAUTO.*C_NONE.*C_REG.*/\1 y+0x10004(FP), R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_NONE.*C_LAUTO.*/\1 F4, result+0x10004(FP)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_LAUTO.*C_NONE.*C_FREG.*/\1 y+0x10004(FP), F4/g' $INPUT

#C_SBRA
sed -i 's/^.*{A\(.*\), C_REG.*C_REG.*C_SBRA.*/\1 R4, R5, 1(PC)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_SBRA.*/\1 R4, 1(PC)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_NONE.*C_NONE.*C_SBRA.*/\1 1(PC)/g' $INPUT

#C_ADDR
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_ADDR.*/\1 R4, name(SB)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_NONE.*C_ADDR.*/\1 F4, name(SB)/g' $INPUT
sed -i 's/^.*{A\(.*\), C_ADDR.*C_NONE.*C_REG.*/\1 name(SB), R4/g' $INPUT
sed -i 's/^.*{A\(.*\), C_ADDR.*C_NONE.*C_FREG.*/\1 name(SB), F4/g' $INPUT

#REG <-> FREG
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_FREG.*/\1 R4, F5/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_NONE.*C_REG.*/\1 F4, R5/g' $INPUT

#others
sed -i 's/^.*{A\(.*\), C_NONE.*C_NONE.*C_NONE.*/\1/g' $INPUT
sed -i 's/^.*{A\(.*\), C_REG.*C_REG.*C_REG.*/\1 R4, R5, R6/g' $INPUT
sed -i 's/^.*{A\(.*\), C_REG.*C_NONE.*C_REG.*/\1 R4, R5/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_NONE.*C_FREG.*/\1 F4, F5/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_REG.*C_FREG.*/\1 F4, R5, F6/g' $INPUT
sed -i 's/^.*{A\(.*\), C_FREG.*C_REG.*C_NONE.*/\1 F4, R5/g' $INPUT
