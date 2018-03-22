#!/bin/bash

GOOS=windows
GOARCH=amd64
FileName=ClanEventsBot.exe
BuildNumber=$(date +%y%m%d%H%M%S)

if [ ! -z "$2" ]
then
	BuildNumber=$2
fi

case $1 in
	win32)
		GOOS=windows
		GOARCH=386
		FileName=ClanEventsBot32.exe
		;;
	win64)
		GOOS=windows
		GOARCH=amd64
		FileName=ClanEventsBot64.exe
		;;
	linux32)
		GOOS=linux
		GOARCH=386
		FileName=ClanEventsBot32
		;;
	linux64)
		GOOS=linux
		GOARCH=amd64
		FileName=ClanEventsBot64
		;;
	all)
		;;
	*)
		echo "Invalid OS and Architecture specified ($1)"
		echo Accepted Values:
		echo win32
		echo win64
		echo linux32
		echo linux64
		echo all
		exit
		;;
esac

if [ $1 == 'all' ]
then
	counter=1
	while [ $counter -le 4 ]
	do
		case $counter in
			1)
				platform=win32
				;;
			2)
				platform=win64
				;;
			3)
				platform=linux32
				;;
			4)
				platform=linux64
				;;
		esac
		./$0 $platform $BuildNumber
		((counter++))
	done
else
	echo Building $FileName, BuildNumber=$BuildNumber
	GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$FileName -ldflags "-X main.buildNumber=$BuildNumber" main.go configuration.go commands.go events.go
fi
cp ClanEventsBot.yaml bin