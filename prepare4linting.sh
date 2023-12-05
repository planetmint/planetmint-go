#!/bin/bash
export LINT="//lint:file-ignore SA1019 Ignore all deprecation errors, it's generated"

# Add lint-ignore comment to beginning of files                                
sed -i "1i${LINT}" ./x/asset/types/query.pb.gw.go
sed -i "1i${LINT}" ./x/machine/types/query.pb.gw.go
sed -i "1i${LINT}" ./x/dao/types/query.pb.gw.go