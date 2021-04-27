#!/bin/bash
DIR="/dev/shm/data/"
rm -r $DIR
mkdir $DIR
### make 20 MB file
cat ../data/10MB.txt > "${DIR}20MB.txt"
echo " + " >> "${DIR}20MB.txt"
cat ../data/10MB.txt >> "${DIR}20MB.txt"
###make 100MB
cat ../data/10MB.txt > "${DIR}100MB.txt"
for i in {1..9}
do
echo " + " >> "${DIR}100MB.txt"
cat ../data/10MB.txt >> "${DIR}100MB.txt"
done
#make 1GB
cat ../data/10MB.txt >"${DIR}1GB.txt"
for i in {1..99}
do
echo " + " >> "${DIR}1GB.txt"
cat ../data/10MB.txt >> "${DIR}1GB.txt"
done
#make 10GB
cat ../data/10MB.txt >"${DIR}10GB.txt"
for i in {1..999}
do
echo " + " >> "${DIR}10GB.txt"
cat ../data/10MB.txt >> "${DIR}10GB.txt"
done
#


