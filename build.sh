#!/usr/bin/env bash

case $1 in
  run-gui)
    if [[ $2 == "theme-light" ]]; then
      GTK_THEME="Adwaita:light"
    else
      GTK_THEME="Adwaita:dark"
    fi
    go build -mod=vendor -v -o dist/clyp_linux_amd64_v1/clyp . \
      && cd dist/clyp_linux_amd64_v1/ \
      && GTK_THEME=$GTK_THEME RUN_ENV=dev ./clyp
    ;;
  run-watcher)
    go build -mod=vendor -v -o dist/clyp_linux_amd64_v1/clyp . \
      && cd dist/clyp_linux_amd64_v1/ \
      && ./clyp --watch
    ;;
  package-snapshot)
    goreleaser release --snapshot --clean \
      && nfpm pkg --packager deb --target dist/ \
      && nfpm pkg --packager archlinux --target dist/
    ;;
esac
