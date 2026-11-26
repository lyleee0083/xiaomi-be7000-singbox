#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
    echo "Usage: bash release/build-be7000.sh <output-directory>" >&2
    exit 2
fi

release_dir=$(cd -- "$(dirname -- "$0")" && pwd)
source_dir=$(cd -- "$release_dir/.." && pwd)
mkdir -p -- "$1"
output_dir=$(cd -- "$1" && pwd)
spec="$release_dir/be7000.json"
version=$(jq -er '.version | select(test("^[0-9]+[.][0-9]+[.][0-9]+$"))' "$spec")
upstream=$(jq -er '.upstream_version' "$spec")
revision=$(jq -er '.upstream_revision' "$spec")
name="sing-box-xiaomi-be7000-$version"
source_status=$(git -C "$source_dir" status --porcelain)
if [[ -n "$source_status" ]] || ! git -C "$source_dir" merge-base --is-ancestor "$revision" HEAD; then
    echo "Use a clean fork checkout based on upstream $revision." >&2
    exit 1
fi

cd -- "$source_dir"
tags="with_utls,with_musl,badlinkname,tfogo_checklinkname0"
ldflags=$(cat release/LDFLAGS)
CC=musl-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm64 GOARM64=v8.0 GOTOOLCHAIN=local GOWORK=off \
    go build -mod=readonly -trimpath -buildvcs=true -tags "$tags" \
        -ldflags "$ldflags -linkmode=external -extldflags=-static -X github.com/sagernet/sing-box/constant.Version=$upstream -X main.distributionVersion=$version -s -w -buildid=" \
        -o "$output_dir/$name" ./cmd/sing-box

cd -- "$output_dir"
sha256sum -- "$name" > SHA256SUMS.txt
