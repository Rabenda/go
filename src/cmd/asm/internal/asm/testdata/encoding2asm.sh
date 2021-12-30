#!/bin/sh

set -x

# Note: The input file looks like:
#85 e0 40 00 a5 e0 48 00 85 00 41 00 a5 00 45 00
#de 83 10 00 c4 03 80 29 de 83 10 00 de ff 10 00
#c4 03 00 29 de 83 10 00 de ff 10 00 c4 03 00 29
#de 83 10 00 c4 03 80 28 de 83 10 00 de ff 10 00
#c4 03 00 28 de 83 10 00 c4 03 00 2a 04 02 00 14
#84 00 80 03 04 02 00 14 84 00 80 03 44 02 00 14
#84 14 8d 03 44 02 00 14 84 14 8d 03 85 04 80 02
#84 04 80 02 85 04 c0 02 84 04 c0 02 1e fc ff 03
#85 f8 14 00 1e fc ff 03 84 f8 14 00 00 05 00 48
#00 00 00 00 de 83 10 00 de ff 10 00 c4 03 00 2b
#de 83 10 00 de ff 10 00 c4 03 80 2b de 83 10 00
#de ff 10 00 c4 03 40 2b de 83 10 00 de ff 10 00
#c4 03 c0 2b 1e fc ff 03 04 fc 14 01 1e 04 80 02
#04 fc 14 01 85 10 00 58 00 00 2a 00 04 10 00 58
#00 00 2a 00


INPUT=$1
OUTPUT=${INPUT%.*}-word.s

cp $INPUT $OUTPUT

sed -i 's/.\{12\}/&\n/g' $OUTPUT
sed -i 's/^\(.*\) $/\1/g' $OUTPUT
sed -i 's/^\(.*\) \(.*\) \(.*\) \(.*\)/.word 0x\4\3\2\1/g' $OUTPUT

loong64-linux-gnu-gcc -c $OUTPUT
loong64-linux-gnu-objdump -d ${OUTPUT%.*}.o
