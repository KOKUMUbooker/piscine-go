#! /bin/bash

# 0 - none
# 1 - exec
# 2 - write
# 3 - wx
# 4 - read
# 5 - rx
# 6 - rw
# 7 - rwx

# clear file contents
rm -rf 0  1  2  3  4  5  6  7  8  9  A done.tar

mkdir 0
chmod 401 0
touch -d "1986-01-05 00:00:00 UTC" 0 # touch -d "YYYY-MM-DD HH:MM:SS UTC" filename

touch 1
chmod 402 1
touch -d "1986-11-13 00:01:00 UTC" 1

touch 2
chmod 604 2
touch -d "1988-03-05 00:10:00 UTC" 2

# Create symlink file 3 to dir 0
touch 3
chmod 777 3
ln -sf 0 3 # used -sf to update the file to a symlink
touch -d "1990-02-16 00:11:00 UTC" -h 3 # add -h to modify links timestamps & not what its pointing to

touch 4
chmod 510 4
touch -d "1990-10-07 01:00:00 UTC" 4

touch 5
chmod 460 5
touch -d "1990-11-07 01:01:00 UTC" 5

touch 6
chmod 460 6
touch -d "1991-02-08 01:10:00 UTC" 6

touch 7
chmod 510 7
touch -d "1991-03-08 01:11:00 UTC" 7

touch 8
chmod 604 8
touch -d "1994-05-20 10:00:00 UTC" 8

touch 9
chmod 402 9
touch -d "1994-06-10 10:01:00 UTC" 9

mkdir A
chmod 401 A
touch -d "1995-04-10 10:10:00 UTC" A

tar -cf done.tar 0  1  2  3  4  5  6  7  8  9  A 