#!/bin/bash

check_env_var() {
  if [ -z "${2}" ]; then
    echo "Error: ${1} env var not defined"
    exit 1
  fi
}

while getopts :b:g:i:p:t:v: opt; do
  case $opt in
    b)
      BUILD_TIME=${OPTARG}
      ;;
    g)
      GIT_COMMIT=${OPTARG}
      ;;
    i)
      IMAGE=${OPTARG}
      ;;
    p)
      PREV_VERSION=${OPTARG}
      ;;
    t)
      TAG=${OPTARG}
      ;;
    v)
      VERSION=${OPTARG}
      ;;
  esac
done

# Remove processed options from arguments
shift $(( OPTIND - 1 ));

check_env_var "BUILD_TIME" ${BUILD_TIME}
check_env_var "IMAGE" ${IMAGE}
check_env_var "PREV_VERSION" ${PREV_VERSION}
check_env_var "TAG" ${TAG}
check_env_var "VERSION" ${VERSION}

# Updates version/version.go
sed -i "s/Version      = .*/Version      = \"${VERSION}\""/ version/version.go
sed -i "s/PriorVersion = .*/PriorVersion = \"${PREV_VERSION}\""/ version/version.go

echo "Updating vendor directory ..."
go mod vendor

echo "Run generators ..."
go generate ./...

echo "Executing test build ..."
./scripts/go-test.sh

echo
echo "=== Compiling app ..."
export GO111MODULE=on
GOFLAGS="-X github.com/apicurio/apicurio-operators/apicurito/version.Version=${VERSION}"
GOFLAGS=${GOFLAGS}" -X github.com/apicurio/apicurio-operators/apicurito/version.PriorVersion=${PREV_VERSION}"
GOFLAGS=${GOFLAGS}" -X github.com/apicurio/apicurio-operators/apicurito/pkg.Version=${VERSION}"
GOFLAGS=${GOFLAGS}" -X github.com/apicurio/apicurio-operators/apicurito/pkg.BuildDateTime=${BUILD_TIME}"
GOFLAGS=${GOFLAGS}" -X github.com/apicurio/apicurio-operators/apicurito/pkg/cmd.GitCommit=${GIT_COMMIT}"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -a \
  -ldflags "${GOFLAGS}" \
  -o build/_output/bin/apicurito \
  -mod=vendor github.com/apicurio/apicurio-operators/apicurito/cmd/manager
if [ $? != 0 ]; then
  echo "Error: build failed"
  exit 1
fi

docker build . -t ${IMAGE}:${TAG}
