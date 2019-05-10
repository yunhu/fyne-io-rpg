#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package static -name U1 u1.png > $FILE
$BIN/fyne bundle -package static -append -name U2 u2.png >> $FILE
$BIN/fyne bundle -package static -append -name U3 u3.png >> $FILE


$BIN/fyne bundle -package static -append -name D1 d1.png >> $FILE
$BIN/fyne bundle -package static -append -name D2 d2.png >> $FILE
$BIN/fyne bundle -package static -append -name D3 d3.png >> $FILE

$BIN/fyne bundle -package static -append -name L1 l1.png >> $FILE
$BIN/fyne bundle -package static -append -name L2 l2.png >> $FILE
$BIN/fyne bundle -package static -append -name L3 l3.png >> $FILE

$BIN/fyne bundle -package static -append -name R1 r1.png >> $FILE
$BIN/fyne bundle -package static -append -name R2 r2.png >> $FILE
$BIN/fyne bundle -package static -append -name R3 r3.png >> $FILE




