#!/usr/bin/env bash
set -euo pipefail

NAME="${1:?usage: build_benzhi_docker.sh <image-name> <platform>}"
PLATFORM="${2:-linux/amd64}"
IMAGE="benzhi/${NAME}:latest"

docker build --platform "${PLATFORM}" -f benzhi.Dockerfile -t "${IMAGE}" .
echo "built ${IMAGE} (${PLATFORM})"
