#!/bin/sh
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
go build -o ./bin/crawlGoCN *.go

ps -aux|grep crawlGoCN |grep -v grep

if [ $? -ne 0 ]
then
		./bin/crawlGoCN &
else
	echo "crawlGoCN running"

fi

##"0 */2 * * * SHELL_FOLDER/build.sh" >> /etc/crontab
