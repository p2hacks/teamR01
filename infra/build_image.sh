#!/bin/bash

cd ../server

for i in "bff" "card-info" "get-message" "send-message" "present" "pair" "rate" "user"
do
    docker build -t omamama/$i:1.1.0 $i/
    kind load docker-image omamama/$i:1.1.0 --name omamama --nodes omamama-worker,omamama-worker2
done