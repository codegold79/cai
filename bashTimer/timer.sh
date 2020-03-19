#!/bin/bash
# Developed to work on MacOS Catalina using Zsh shell.

echo "Timer interval (in minutes)?"
read int

echo "Timer number of executions?"
read count

for (( i=0; i<count; i++ ))
do
	echo "Next chime in $int minute(s)..."
	sleep $((int*60))
	afplay ~/Downloads/audio/voy_door_chime.mp3
done
