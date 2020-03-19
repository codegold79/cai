#!/bin/bash
#
# Developed to work on MacOS Catalina using Zsh shell.
# Depends on sound player afplay and bash calculator bc (in order to be able to use float numbers,
# as bash only allows for integer calculations).

echo "Timer interval (in minutes)?"
read int

echo "Timer number of executions?"
read count

echo "Sound choice? [1] door chime [2] sickbay replicator [3] TNG replicator"
read sound

for (( i=0; i<count; i++ ))
do
	echo "Next chime in $int minute(s)..."
	sleep $(echo "$int*60" | bc)

	case $sound in
		1)
			echo "Play door chime"
			afplay ~/Downloads/audio/voy_door_chime.mp3
			;;
		2)
			echo "Play sickbay replicator"
			afplay ~/Downloads/audio/voy_sickbay_replicator.mp3
			;;
		3)
			echo "Play TNG replicator"
			afplay ~/Downloads/audio/tng_replicator.mp3
			;;
	esac
done
