#! /bin/bash

curl -s "https://learn.zone01kisumu.ke/assets/superhero/all.json" | 
  jq -r ".[] | select(.id == ${HERO_ID}) | .connections.relatives" | 
  tr '\n' ' ' | sed 's/ $//'