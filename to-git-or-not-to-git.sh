#/bin/bash

# Fetch superheroes & use jq to access item whose id is 170
obj_170=$(curl -s "https://learn.zone01kisumu.ke/assets/superhero/all.json" | jq '.[] | select(.id == 170)')

# Added -r flag to ensure output is plain strings ie lack ""
name=$(echo "$obj_170" | jq -r '.name')
power=$(echo "$obj_170" | jq -r '.powerstats.power')
gender=$(echo "$obj_170" | jq -r '.appearance.gender')

echo ${name}
echo ${power}
echo ${gender}