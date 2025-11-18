#! /bin/bash

# col 9 seems to be the last column that lists out the file names on "ls -l" output 
# sed '1~2c\' -> replace odd number lines with nothing ie delete them
#             -> 1~2 specifies the range
#                where 1 is the starting point
#                ~2 specifies number of steps to be incremented till the end is reached
#                eg 1, 3(1+2), 5(3+2), 7(5+2)
#              -> c - to specify action is change
#              -> \ - specify the action to be used to achieve the action ie
#                 change all odd numbered line with nothing ie delete them
ls -l | awk 'NR % 2 == 0'
