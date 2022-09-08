#!/bin/bash

if [[ -z ${CI} ]]; then
    echo "Executing openapi-gen to generate the api ..."
    if hash openapi-gen 2>/dev/null; then
      for version in `find pkg/apis/apicur -maxdepth 1 -mindepth 1 -type d -printf "%f\n"`
      do
          echo "Generating api for ${version} ..."
          openapi-gen --logtostderr=true -o "" \
              --go-header-file ./boilerplate/boilerplate.go.txt \
              -i ./pkg/apis/apicur/${version} -O zz_generated.openapi -p ./pkg/apis/apicur/${version}
          if [ $? != 0 ]; then
              echo "Error: openapi-gen failed to generate the API"
              exit 1
          fi
      done
    else
        echo "skipping go openapi generation"
    fi

    osdk_version=$(operator-sdk version | sed -n 's/.*version: "v\([^"]*\)".*/\1/p')
    if [[ ${osdk_version} == 0.* ]]; then
      echo "operator-sdk >= 1.0.0 required. Please upgrade ..."
      exit 1
    else
      echo "Executing config/Makefile::generate ..."
      # Calls the config/Makefile which in turn uses controller-gen
      # As described by the operator-sdk documentation
      make -C config generate
    fi
fi
go vet ./...
