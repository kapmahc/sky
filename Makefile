dist=build
theme=bootstrap
pkg=github.com/kapmahc/sky/web

VERSION=`git rev-parse --short HEAD`
BUILD_TIME=`date -R`
AUTHOR_NAME=`git config --get user.name`
AUTHOR_EMAIL=`git config --get user.email`
COPYRIGHT=`head -n 1 LICENSE`
USAGE=`sed -n '3p' README.md`


build: backend frontend
	tar jcvf dist.tar.bz2 $(dist)

frontend:
	cd dashboard && npm run build
	mkdir -pv $(dist)/public
	-cp -rv dashboard/build $(dist)/public/my

backend:
	go build -ldflags "-s -w -X ${pkg}.Version=${VERSION} -X '${pkg}.BuildTime=${BUILD_TIME}' -X '${pkg}.AuthorName=${AUTHOR_NAME}' -X ${pkg}.AuthorEmail=${AUTHOR_EMAIL} -X '${pkg}.Copyright=${COPYRIGHT}' -X '${pkg}.Usage=${USAGE}'" -o ${dist}/sky main.go
	-cp -rv locales db templates $(dist)/
	mkdir -pv $(dist)/themes/${theme}
	-cp -rv themes/${theme}/package.json themes/${theme}/views themes/${theme}/assets $(dist)/themes/${theme}


clean:
	-rm -rv $(dist) dist.tar.bz2
	-rm -rv dashboard/build


init:
	go get -u github.com/kardianos/govendor
	govendor sync
	npm install
