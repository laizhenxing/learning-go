#!/bin/bash;

for n in $(seq 1 1 10)
do
    nohup curl -XGET -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIwLTA4LTI4VDE4OjI0OjU3LjYzNDU3MTI2OFoiLCJpYXQiOjE1OTg2MjgyOTcsImlkIjo1MiwibmJmIjoxNTk4NjI4Mjk3LCJ1c2VybmFtZSI6InQ2In0.32x-knAhzulrk87aHSnLTvLeniN62NB_93qEz-fkGPY" http://uapiserver.com/v1/users &>/dev/null
done
