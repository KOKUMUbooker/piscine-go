#! /bin/bash

# Get interview number of this witness
iNum=$(grep "Church" interviews/* | awk '{print $1}' | sed 's/.*interview-//' | sed 's/:.*//')

# Print out env variable
echo $iNum

# Print out the contents f the interview
cat "interviews/interview-${iNum}"

# Print out contents of main suspect
echo "$MAIN_SUSPECT"