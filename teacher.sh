#! /bin/bash

# Locate base directory (mystery folder)
mystery_dir="$(find . -type d -name mystery )"

people_file="$mystery_dir/people"
vehicles_file="$mystery_dir/vehicles"

# Get interview number of this witness
interviews_dir="$mystery_dir/interviews"
iNum=$(grep -R "Church" "$interviews_dir" | awk '{print $1}' | sed 's/.*interview-//' | sed 's/:.*//')

# Isolate interview number into an env variable
export interviewnum=$iNum

# Print out env variable
echo $interviewnum

# Print out the contents f the interview
cat "$interviews_dir/interview-$interviewnum"

# Print out contents of main suspect
echo "$MAIN_SUSPECT"