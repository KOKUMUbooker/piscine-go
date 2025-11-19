#! /bin/bash

# Expected target output :
# Key witness - Annabel Church
# Witness interview number - 699607
# Car color & make of main suspect - Blue Honda
# Names of 3 other suspects :
#         - Joe Germuska
#         - Hellen Maher
#         - Erika Owens

# Locate base directory (mystery folder)
mystery_dir="$(find . -type d -name mystery )"

people_file="$mystery_dir/people"
vehicles_file="$mystery_dir/vehicles"

# Display first & last name of your key witness
grep "Annabel Church" "$people_file" | awk '{print $1, $2}'

# Display interview number of this witness
interviews_dir="$mystery_dir/interviews"
grep -R "Church" "$interviews_dir" | awk '{print $1}' | sed 's/.*interview-//' | sed 's/:.*//'

# Display colour and the make of the car of the main suspect
grep "License Plate L337" -A 6  "$vehicles_file" | grep "Owner: Dartey Henv" -B 2 | sed "3d" | awk '{print $2}' | sort | paste -sd\ \

# List out 3 other suspects
grep "License Plate L337" -A 6 "$vehicles_file" | 
    grep "Color: Blue" -A 4 -B 2 | 
    grep "Height: 6'" -A 2 -B 4 |
    grep "Make: Honda" -A 5 -B 1 | 
    grep "Owner:" | grep -v "Dartey Henv" | awk '{print $2 , $3}' | sort -k2
