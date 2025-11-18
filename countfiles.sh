#! /bin/bash

# wc -l -> for counting line length

# Find only files 
fileCount=$(find . -type f -name "*" | wc -l)

# Find only directories
dirCount=$(find . -type d -name "*" | wc -l) # Already includes the current directory so no need to increment the final result by 1

count=$((fileCount + dirCount )) 

echo ${count} 