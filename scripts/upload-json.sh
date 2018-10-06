#!/usr/bin/env bash

ROOT=$(cd "$(dirname ${BASH_SOURCE[0]})/.."; pwd)

CONFIG_URL=$(grep 'const URL' "$ROOT/fuzzy/config.go" | cut -d'"' -f2)
curl -XPOST -d "@$ROOT/config.json" $CONFIG_URL
echo
