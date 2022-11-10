#!/bin/bash
cd "${0%/*}"
root_dir="$(pwd)"

version="${VERSION:-1.19}"
image="golang:${version}"

container_app_dir="/app"
container_cache_dir="output/cache"
container_gopath_dir="gopath"
cache_dir="${root_dir}/${container_cache_dir}"

if [ ! -d "${cache_dir}" ]; then
  mkdir -p "${cache_dir}"
fi

docker run --rm \
  -u $(id -u ${USER}):$(id -g ${USER}) \
  --workdir "${container_app_dir}/src" \
  --volume "${root_dir}:/${container_app_dir}" \
  --env "GOCACHE=${container_app_dir}/${container_cache_dir}" \
  --env "GOPATH=${container_app_dir}/${container_gopath_dir}" \
  "${image}" go $@
