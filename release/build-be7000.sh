#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 2 ]]; then
    echo "Usage: bash release/build-be7000.sh <upstream-checkout> <output-directory>" >&2
    exit 2
fi

release_dir=$(cd -- "$(dirname -- "$0")" && pwd)
source_dir=$(cd -- "$1" && pwd)
mkdir -p -- "$2"
output_dir=$(cd -- "$2" && pwd)
spec="$release_dir/be7000.json"
version=$(jq -er '.version | select(test("^[0-9]+[.][0-9]+[.][0-9]+$"))' "$spec")
upstream=$(jq -er '.upstream_version' "$spec")
revision=$(jq -er '.upstream_revision' "$spec")
name="sing-box-xiaomi-be7000-$version"
patch="$release_dir/be7000-version.patch"
source_status=$(git -C "$source_dir" status --porcelain)
source_revision=$(git -C "$source_dir" rev-parse HEAD)
if [[ "$source_revision" != "$revision" || -n "$source_status" ]]; then
    echo "Use a clean upstream checkout at $revision." >&2
    exit 1
fi

git -C "$source_dir" apply --unidiff-zero "$patch"
trap 'git -C "$source_dir" apply --reverse --unidiff-zero "$patch"' EXIT
cd -- "$source_dir"
tags="$(cat release/DEFAULT_BUILD_TAGS),with_musl"
ldflags=$(cat release/LDFLAGS)
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 GOARM64=v8.0 GOTOOLCHAIN=local GOWORK=off \
    go build -mod=readonly -trimpath -buildvcs=true -tags "$tags" \
        -ldflags "$ldflags -X github.com/sagernet/sing-box/constant.Version=$upstream -X main.distributionVersion=$version -s -w -buildid=" \
        -o "$output_dir/$name" ./cmd/sing-box

cd -- "$output_dir"
sha256sum -- "$name" > SHA256SUMS.txt
