make:
  - cd $GOPATH/src/github.com/junpayment/gin_labo/;
  - glide install;
  - cd $GOPATH/src/github.com/junpayment/gin_labo/vendor/github.com/suapapa/go_sass/libsass;
  - ./configure;
  - make;
  - sudo make install;
  - cd $GOPATH
  - go build github.com/junpayment/gin_labo