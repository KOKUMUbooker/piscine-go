#! /bin/bash

# --argjson hId "$HERO_ID" -> safely pass env var HERO_ID under a new variable name hId for access under the jq query
# seg 's/"//g' -
#   s -> substitute, the action
#   " -> pattern to substitute
#   // -> what to be used for substituting ie nothing
#   g -> Ensures action is executed for every occurence of the pattern

curl -s "https://learn.zone01kisumu.ke/assets/superhero/all.json" | jq --argjson hId "$HERO_ID" '.[] | select(.id == $hId) | .connections.relatives' | sed 's/"//g'