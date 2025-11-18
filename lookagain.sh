#! /bin/bash

# "sort -r" -> r flag to reverse the default ascending order
# "cut -c 3-" -> -c flag to specify character count that I want
#             -> "3-" specify to take charactes starting from index 3 to the end ie exclude "./" 
# 'sed "s/.sh//g"' -> Replace all instances of ".sh" with nothing ie specified by //
find . -name "*.sh" | sort -r | cut -c 3- | sed "s/.sh//g"