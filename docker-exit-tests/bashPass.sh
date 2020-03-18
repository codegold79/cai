#!/bin/bash 
# A bash script that doesn't (shouldn't) have any errors.

user=$(whoami) 

echo "Hi $user. Here are the files and their line counts in pwd," 
echo -n $(pwd)
echo ":"

for i in .* *;
  do if [ ! -d "$i" ]; then 
    echo "$(wc -l $i)";
  fi
done
