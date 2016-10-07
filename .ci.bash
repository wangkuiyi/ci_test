cd ..
GOPATH=`pwd`
mkdir -p src/github.com/wangkuiyi/
ln -sf `pwd`/ci_test $GOPATH/src/github.com/wangkuiyi/ci_test
cd $GOPATH/src/github.com/wangkuiyi/ci_test

echo Working in `pwd`
echo Assuming GOPATH=$GOPATH

go get -u ./...
go run example.go
