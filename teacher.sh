#! /bin/bash

# Get interview number of this witness
iNum=$(head -n 179 streets/Buckingham_Place | tail -n 1 | grep -o '[0-9]\+')

# Print out env variable
echo $iNum

# Print out the contents f the interview
cat "interviews/interview-${iNum}"

# Print out contents of main suspect
echo "$MAIN_SUSPECT"