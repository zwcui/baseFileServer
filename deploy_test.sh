#!/bin/bash

set -e

runmode_from_conf=`awk '$1=="runmode" {print $3}' ./conf/app.conf`
version_from_conf=`awk '$1=="version" {print $3}' ./conf/app.conf`
httpport_from_conf=`awk '$1=="httpport" {print $3}' ./conf/app.conf`

if [ $# == 0 ] && [ -z $version_from_conf ]; then
    echo "baby, we need a version code"
    exit 1
fi

runmode=$runmode_from_conf
if [ $# == 1 ]; then
    runmode=$1
fi

echo $runmode

version=$version_from_conf
if [ $# == 1 ]; then
    version=$1
fi

echo $version

httpport=$httpport_from_conf
if [ $# == 1 ]; then
    httpport=$1
fi

echo $httpport

default_runmode="test"
runmode=`awk '$1=="runmode" {print $3}' ./conf/app.conf`

if [ $default_runmode != $runmode ]
then
    echo "$runmode is err,you should in $default_runmode"
	exit 1
fi

ssh  root@106.14.202.179 version=$version httpport=$httpport runmode=$runmode 'bash -se' <<'ENDSSH'
cd ~/app/api/baseFileServer/test/baseFileServer
git pull;
echo basefileserver\_$runmode
#go clean;
if docker build -t basefileserver\_$runmode:$version .
then
    echo "stop and rm old container,start new one..."
#    docker stop basefileserver\_$runmode
#    docker rm basefileserver\_$runmode
    docker run --restart=always --name basefileserver\_$runmode -v /data/baseFileServer:/go/src/baseFileServer/data -d -p $httpport:8088 basefileserver\_$runmode:$version
    docker ps
fi
ENDSSH