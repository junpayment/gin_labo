#!/usr/bin/env bash
GOPATH=${GOPATH}:/var/app/current

curl https://glide.sh/get | sh
cd $GOPATH/src/github.com/junpayment/gin_labo/;
glide install;
cd $GOPATH/src/github.com/junpayment/gin_labo/vendor/github.com/suapapa/go_sass/libsass;
./configure;
make;
sudo make install;
cd $GOPATH
go build github.com/junpayment/gin_labo
