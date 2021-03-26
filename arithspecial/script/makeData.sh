#!/bin/bash
cd ~/go/src/github.com/filgor84/languages/arithspecial/data
### make 20 MB file
cat 10MB.txt > 20MB.txt
echo " + " >> 20MB.txt
cat 10MB.txt >> 20MB.txt
###make 100MB
cat 10MB.txt >100MB.txt
for i in {1..9}
do
echo " + " >> 100MB.txt
cat 10MB.txt >> 100MB.txt
done
#make 1GB
cat 10MB.txt >1GB.txt
for i in {1..99}
do
echo " + " >> 1GB.txt
cat 10MB.txt >> 1GB.txt
done