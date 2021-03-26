#!/bin/bash
cd ~/go/src/github.com/filgor84/languages/arithspecial/data
DIR="/dev/shm/"
### make 20 MB file
cat 10MB.txt > "${DIR}20MB.txt"
echo " + " >> "${DIR}20MB.txt"
cat 10MB.txt >> "${DIR}20MB.txt"
###make 100MB
cat 10MB.txt > "${DIR}100MB.txt"
for i in {1..9}
do
echo " + " >> "${DIR}100MB.txt"
cat 10MB.txt >> "${DIR}100MB.txt"
done
#make 1GB
cat 10MB.txt >"${DIR}1GB.txt"
for i in {1..99}
do
echo " + " >> "${DIR}1GB.txt"
cat 10MB.txt >> "${DIR}1GB.txt"
done