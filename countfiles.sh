#! /bin/bash

# wc -l -> for counting line length

# Find only files 
fileCount=$(find . -type f -name "*" | wc -l)

# Find only directories
dirCount=$(find . -type d -name "*" | wc -l)

count=$((fileCount + dirCount + 1 )) # Add 1 to ensure current dir is included

echo ${count} 