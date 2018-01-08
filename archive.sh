#!/usr/bin/env bash

DEST_DIR=/tmp/working_`date +%s`
REP_URL="git@github.com:junpayment/gin_labo.git"

mkdir -p $DEST_DIR/go
mkdir -p $DEST_DIR/go/bin $DEST_DIR/go/pkg $DEST_DIR/go/src/github.com/junpayment/
git clone --depth 1 $REP_URL $DEST_DIR/go/src/github.com/junpayment/gin_labo

cp -a $DEST_DIR/go/src/github.com/junpayment/gin_labo/.ebextensions $DEST_DIR/go/
cp $DEST_DIR/go/src/github.com/junpayment/gin_labo/build.sh $DEST_DIR/go/
cp $DEST_DIR/go/src/github.com/junpayment/gin_labo/Buildfile $DEST_DIR/go/
cp $DEST_DIR/go/src/github.com/junpayment/gin_labo/Procfile $DEST_DIR/go/

cd $DEST_DIR/go
zip ../go.zip -r * .[^.]*
