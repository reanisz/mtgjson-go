#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)

cd $SCRIPT_DIR/data
wget "https://mtgjson.com/api/v5/AllPrintings.sqlite.zip"
rm AllPrintings.sqlite
unzip AllPrintings.sqlite.zip
rm AllPrintings.sqlite.zip
