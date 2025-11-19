#! /bin/bash

# Get interview number of this witness
iNum=$(grep "Church" interviews/* | awk '{print $1}' | sed 's/.*interview-//' | sed 's/:.*//')

# Isolate interview number into an env variable
export interviewnum=$iNum

# Print out env variable
echo $interviewnum

# Print out the contents f the interview
cat "interviews/interview-$interviewnum"

# Print out contents of main suspect
echo "$MAIN_SUSPECT"