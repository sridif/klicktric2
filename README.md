klicktric
=========

gamification app for mo

on a new instance..

reroute the port 80
sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 9921

store this git under 
/home/ubuntu/dev/src


in ~/.profile 
add 

export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/home/ubuntu/dev

to open 
