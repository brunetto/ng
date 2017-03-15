package new

var BuildScript string = `
#!/usr/bin/env bash
#set -x

cwd=$(pwd)

echo "Test"
cd ../../
go test ./...
cd $cwd

# Check test result, soon or later

echo "Remove old builds (just to be sure)"
rm ../../docker/standalone/{{bin}}

version=` + "`date -u +%Y%m%d-%H%M%S`" + `

# Generate version reminder
rm last_version_number_*
touch "last_version_number_$version"

echo "Build"
go build -ldflags "-X {{package}}.version=$version"

echo "Cross-compile for linux (may take a while)"
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$version" -o {{bin}}_linux main.go

echo "Copy executable to docker scratch folder"
cp {{bin}}_linux ../../docker/standalone/{{bin}}

echo "Building ng docker standalone service"
cd ../../docker/
ws_build.sh
cd ${cwd}

echo "Done"
`
