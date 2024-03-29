#!/bin/bash
#
# Developed to work on Ubuntu 19 using bash shell.
# Depends on sound player paplay and bash calculator bc (in order to be able to use float numbers,
# as bash only allows for integer calculations).

if [[ $# -eq 0 ]] ; then
	echo "Timer interval (in minutes)?"
	read int

	echo "Timer number of executions?"
	read count

	echo "Sound choice? [1] door chime [2] sickbay replicator [3] TNG replicator [4] ding [5] Tibeten bowl"
	read sound

else
	int=$1
	count=${2:-50}
	sound=${3:-5}
fi

for (( i=0; i<count; i++ ))
do
	echo `date +"[%D %H:%M %p] "`"Next alert in $int minute(s)..."
	sleep $(echo "$int*60" | bc)

	case $sound in
		1)
			echo "Play door chime"
			paplay ~/Downloads/audio/processed/voy_door_chime.wav
			;;
		2)
			echo "Play sickbay replicator"
			paplay ~/Downloads/audio/processed/voy_sickbay_replicator.wav
			;;
		3)
			echo "Play TNG replicator"
			paplay ~/Downloads/audio/processed/tng_replicator.wav
			;;
		4)
			echo "Play gkillhour_ding-4"
			paplay /media/hdd/Sounds/freesound.org/16494__gkillhour__dings/267337__gkillhour__ding-4.wav
			;;
		5)
			echo "Play tibetan-bowl"
			paplay ~/storage/audio/converted/tibetan-bowl.wav
	esac

	printf "$((count-i-1)) chimes remain.\n\n"
done

echo `date +"[%D %H:%M %p] "`"Timer complete"
