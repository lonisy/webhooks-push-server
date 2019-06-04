#!/usr/bin/env bash
src=/tmp/www.myweb.com
dist=/data/html/www.myweb.com
pushd $src && git pull && popd
rsync -ap $src/* $dist
chmod -R 777 $dist/web/cache
chmod -R 777 $dist/web/uploads
