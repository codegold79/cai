#! /bin/bash
# A bash script that fails after a delay

user=$(whoami)

echo "Hello, $user. Here are the files and their sizes in bytes in this directory"
echo -n "$(pwd)"
echo ":"

for i in * 
    do wc -c $i;
done

echo "Send exit code 1 after two delays"
echo "Start first delay"
sleep 3
echo "Start second delay"
sleep 20
exit 1