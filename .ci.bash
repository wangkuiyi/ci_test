cd ..
GOPATH=`pwd`
mkdir -p src/github.com/wangkuiyi/
ln -sf `pwd`/ci_test $GOPATH/src/github.com/wangkuiyi/ci_test
cd $GOPATH/src/github.com/wangkuiyi/ci_test

go get -t ./... # Don't use -u
go test ./...
