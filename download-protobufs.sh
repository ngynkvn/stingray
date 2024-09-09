#!/bin/bash
set -euxo pipefail

ARCHIVE_LINK=https://github.com/SteamDatabase/GameTracking-Deadlock/archive/master.tar.gz

mkdir -p ./deadlock/tmp

if [ -f tmp/master.tar.gz ]; then
  tar -xzf tmp/master.tar.gz --strip-components=1 -C ./deadlock/tmp
else
  curl -L -o - ${ARCHIVE_LINK} | tar -xz --strip-components=1 -C ./deadlock/tmp
fi
cp -a ./deadlock/tmp/Protobufs/*.proto ./deadlock/ && rm -rf ./deadlock/tmp
