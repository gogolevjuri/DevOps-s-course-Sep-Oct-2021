#!/bin/bash
docker image build -t my_docker .
docker container run --rm -dp 8080:8080 --name my_docker my_docker
status=$(docker inspect my_docker --format '{{ .State.Status }}')
for i in {1..10}
do
 if [[ $status=="running" ]];then
  sleep $i
  curl -XPOST -k -d'{"animal":"elephant", "sound":"whoooaaa", "count": 5}' localhost:8080
  break
 fi
done
