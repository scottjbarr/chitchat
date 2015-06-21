#!/bin/bash
#
# Build the binaries.
#
# Author : Scott Barr
# Date   : 21 Jun 2015
#

binaries="chitchatpub chitchatsub"

for b in $binaries; do
    cd ./cmd/$b && go install && cd -
done
