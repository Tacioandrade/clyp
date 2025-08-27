#!/usr/bin/env bash

case $1 in
  run)
    if [[ $2 == "theme-light" ]]; then
      GTK_THEME="Adwaita:light"
    else
      GTK_THEME="Adwaita:dark"
    fi
    go build -mod=vendor -v -o dist/clyp_linux_amd64_v1/clyp . \
      && GTK_THEME=$GTK_THEME dist/clyp_linux_amd64_v1/clyp
    ;;
esac
