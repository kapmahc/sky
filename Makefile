dist=build
pkg=github.com/kapmahc/sky/web

VERSION=`git rev-parse --short HEAD`
BUILD_TIME=`date -R`
AUTHOR_NAME=`git config --get user.name`
AUTHOR_EMAIL=`git config --get user.email`
COPYRIGHT=`head -n 1 LICENSE`
USAGE=`sed -n '3p' README.md`


build:
	go build -ldflags "-s -w -X ${pkg}.Version=${VERSION} -X '${pkg}.BuildTime=${BUILD_TIME}' -X '${pkg}.AuthorName=${AUTHOR_NAME}' -X ${pkg}.AuthorEmail=${AUTHOR_EMAIL} -X '${pkg}.Copyright=${COPYRIGHT}' -X '${pkg}.Usage=${USAGE}'" -o ${dist}/sky main.go
	-cp -rv package.json locales db templates themes $(dist)/
	tar jcvf dist.tar.bz2 $(dist)



clean:
	-rm -rv $(dist) dist.tar.bz2


init:
	go get -u github.com/kardianos/govendor
	govendor sync
	npm install
