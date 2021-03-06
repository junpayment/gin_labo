#!/usr/bin/env bash
export GOPATH=/var/app/staging
export GOBIN=/var/app/staging/bin
export PATH=$PATH:$GOBIN

echo "install glide..."
curl https://glide.sh/get | sh
echo "install glide done"

echo "install go packages..."
cd $GOPATH/src/github.com/junpayment/gin_labo/;
glide install;
echo "install go packages done"

echo "install libsass..."
cd $GOPATH/src/github.com/junpayment/gin_labo/;
glide install;
cd $GOPATH/src/github.com/junpayment/gin_labo/vendor/github.com/suapapa/go_sass/
sh ./install_libsass.sh
echo "install libsass done"

echo "build go binary"
cd $GOPATH
echo $GOPATH; pwd
go install github.com/junpayment/gin_labo || echo error_occuered!
echo "build go binary"
