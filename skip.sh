#! /bin/bash
 
# awk 'NR % 2 == 0'
#   NR -> Represents number of records ie the number of lines 
#   NR % 2 == 0 ie numOfLines % 2 == 0 evaluates to true for even numbered lines
#   thus awk will only print even numbered lines 
ls -l | awk 'NR % 2 == 0'
