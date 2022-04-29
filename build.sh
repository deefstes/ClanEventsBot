#!/bin/bash

# Use git describe to set build number
BuildNumber=$(git describe --tags --dirty --always)
# Remove leading v from buildnumber if in the format v1.2.3
BuildNumber=$(sed -E "s|v([[:digit:]]+\.[[:digit:]]+\.[[:digit:]]+)|\1|" <<< $BuildNumber)

if [ ! -z "$2" ]
then
	BuildNumber=$2
fi

case $1 in
	docker)
		echo Building docker image, BuildNumber=$BuildNumber
		docker build --build-arg BuildNumber=$BuildNumber -t claneventsbot:$BuildNumber .
		;;
	win32)
		echo Building ClanEventsBot32.exe, BuildNumber=$BuildNumber
		GOOS=windows GOARCH=386 go build -o bin/ClanEventsBot32.exe -ldflags "-X main.buildNumber=$BuildNumber" *.go
		;;
	win64)
		echo Building ClanEventsBot64.exe, BuildNumber=$BuildNumber
		GOOS=windows GOARCH=amd64 go build -o bin/ClanEventsBot64.exe -ldflags "-X main.buildNumber=$BuildNumber" *.go
		;;
	linux32)
		echo Building ClanEventsBot32, BuildNumber=$BuildNumber
		GOOS=linux GOARCH=386 go build -o bin/ClanEventsBot32 -ldflags "-X main.buildNumber=$BuildNumber" *.go
		;;
	linux64)
		echo Building ClanEventsBot64, BuildNumber=$BuildNumber
		GOOS=linux GOARCH=amd64 go build -o bin/ClanEventsBot64 -ldflags "-X main.buildNumber=$BuildNumber" *.go
		;;
	all)
		;;
	*)
		echo "Invalid OS and Architecture specified ($1)"
		echo Accepted Values:
		echo docker
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
fi