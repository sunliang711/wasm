#!/bin/bash
for d in $(find . -type d -depth 1);do
    cd $d && go install
    cd -
done
