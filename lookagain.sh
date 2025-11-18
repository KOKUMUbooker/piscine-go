#! /bin/bash

# "sort -r" -> r flag to reverse the default ascending order
# 'sed "s/.*\///"' -> substitute charachters from "." to '/' with nothing ie remove parent directories
# 'sed "s/.sh//g"' -> Replace all instances of ".sh" with nothing ie specified by //
find . -name "*.sh" | sed 's/.*\///' | sed "s/.sh//g" | sort -r 