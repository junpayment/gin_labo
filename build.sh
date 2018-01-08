#!/usr/bin/env bash
GOPATH=/var/app/current
GOBIN=/var/app/current/bin
PATH=$PATH:$GOBIN

curl https://glide.sh/get | sh

cd $GOPATH/src/github.com/junpayment/gin_labo/;
glide install;
cd $GOPATH/src/github.com/junpayment/gin_labo/vendor/github.com/suapapa/go_sass/
sh ./install_libsass.sh

cd $GOPATH
go build github.com/junpayment/gin_labo
