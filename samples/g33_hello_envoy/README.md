# Envoy Hello

# check version (docker)
docker run --rm \
  envoyproxy/envoy-dev:068861364d84ba86f8b6bdc57c637afb10eb6692 \
  --version

# help
docker run --rm \
  envoyproxy/envoy-dev:068861364d84ba86f8b6bdc57c637afb10eb6692 \
  --help

# proxy 9901 to 10000
docker run --rm -it \
  -v $(pwd)/envoy-local.yaml:/envoy-local.yaml \
  -p 9901:9901 \
  -p 10000:10000 \
  envoyproxy/envoy-dev:068861364d84ba86f8b6bdc57c637afb10eb6692 \
  -c /envoy-local.yaml
