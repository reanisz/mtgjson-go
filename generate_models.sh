#!/bin/bash

xo 'file:data/AllPrintings.sqlite?loc=auto' -o models --template-path templates/

grep -l 'type' ./models/*.xo.go | xargs sed -i -e 's/sql\.Null/Null/g'
goimports -w models/*.xo.go

