#!/bin/bash;

curl -XGET -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIwLTA4LTI4VDE4OjI0OjU3LjYzNDU3MTI2OFoiLCJpYXQiOjE1OTg2MjgyOTcsImlkIjo1MiwibmJmIjoxNTk4NjI4Mjk3LCJ1c2VybmFtZSI6InQ2In0.32x-knAhzulrk87aHSnLTvLeniN62NB_93qEz-fkGPY" -H "Content-Type: application/json" https://127.0.0.1:8101/v1/users --cacert conf/server.crt --cert conf/server.crt --key conf/server.key 
