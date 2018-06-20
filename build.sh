#!/bin/bash
BUILD_FOLDER=build
VERSION=$(cat main.go | grep "version " | awk '{print $4}')

bin_dep() {
    BIN=$1
    which $BIN > /dev/null || { echo "@ Dependency $BIN not found !"; exit 1; }
}

create_exe_archive() {
  bin_dep 'zip'

  OUTPUT=$1

  echo "# Creating archive $OUTPUT ..."
  zip -j "$OUTPUT" godirsearch.exe ../README.md ../LICENSE.txt > /dev/null
  rm -rf godirsearch godirsearch.exe
}

create_archive() {
  bin_dep 'zip'

  OUTPUT=$1

  echo "# Creating archive $OUTPUT ..."
  zip -j "$OUTPUT" godirsearch ../README.md ../LICENSE.md > /dev/null
  rm -rf godirsearch godirsearch.exe
}

build_linux_amd64() {
  echo "# Building linux/amd64 ..."
  GOOS=linux GOARCH=amd64 go build -o godirsearch ..
}

build_macos_amd64() {
  echo "# Building darwin/amd64 ..."
  GOOS=darwin GOARCH=amd64 go build -o godirsearch ..
}

build_windows_amd64() {
  echo "# Building windows/amd64 ..."
  GOOS=windows GOARCH=amd64 go build -o godirsearch.exe ..
}

rm -rf $BUILD_FOLDER
mkdir $BUILD_FOLDER
cd $BUILD_FOLDER

build_linux_amd64 && create_archive godirsearch_linux_amd64_$VERSION.zip
build_macos_amd64 && create_archive godirsearch_macos_amd64_$VERSION.zip
build_windows_amd64 && create_exe_archive godirsearch_windows_amd64_$VERSION.zip
shasum -a 256 * > checksums.txt

echo
echo
du -sh *

cd --
