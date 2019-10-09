#!/usr/bin/env bash

readonly CI_FLAG=ci

RED='\033[0;31m'
GREEN='\033[0;32m'
INVERTED='\033[7m'
NC='\033[0m' # No Color

echo -e "${INVERTED}"
echo "USER: " + $USER
echo "PATH: " + $PATH
echo "GOPATH:" + $GOPATH
echo -e "${NC}"

##
# Tidy dependencies
##
go mod tidy
tidyResult=$?
if [ ${tidyResult} != 0 ]; then
	echo -e "${RED}✗ go mod tidy${NC}\n$tidyResult${NC}"
	exit 1
else echo -e "${GREEN}√ go mod tidy${NC}"
fi

##
# GO BUILD
##
buildEnv=""
if [ "$1" == "$CI_FLAG" ]; then
	# build binary statically
	buildEnv="env CGO_ENABLED=0"
fi

${buildEnv} go build -o bin/terraform-provider-gardener

goBuildResult=$?
if [ ${goBuildResult} != 0 ]; then
	echo -e "${RED}✗ go build${NC}\n$goBuildResult${NC}"
	exit 1
else echo -e "${GREEN}√ go build${NC}"
fi

##
# Verify dependencies
##
echo "? go mod verify"
depResult=$(go mod verify)
if [ $? != 0 ]; then
	echo -e "${RED}✗ go mod verify\n$depResult${NC}"
	exit 1
else echo -e "${GREEN}√ go mod verify${NC}"
fi

##
# GO TEST
##
echo "? go test"
go test ./...
# Check if tests passed
if [ $? != 0 ]; then
	echo -e "${RED}✗ go test\n${NC}"
	exit 1
else echo -e "${GREEN}√ go test${NC}"
fi

goFilesToCheck=$(find . -type f -name "*.go" | egrep -v "\/vendor\/|_*/automock/|_*/testdata/|_*export_test.go")

#
# GO FMT
#
goFmtResult=$(echo "${goFilesToCheck}" | xargs -L1 go fmt)
if [ $(echo ${#goFmtResult}) != 0 ]
	then
    	echo -e "${RED}✗ go fmt${NC}\n$goFmtResult${NC}"
    	exit 1;
	else echo -e "${GREEN}√ go fmt${NC}"
fi

##
# GO VET
##
packagesToVet=("./kind/...")

for vPackage in "${packagesToVet[@]}"; do
	vetResult=$(go vet ${vPackage})
	if [ $(echo ${#vetResult}) != 0 ]; then
		echo -e "${RED}✗ go vet ${vPackage} ${NC}\n$vetResult${NC}"
		exit 1
	else echo -e "${GREEN}√ go vet ${vPackage} ${NC}"
	fi
done