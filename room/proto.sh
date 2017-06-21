protoc \
--plugin=protoc-gen-ts=/home/ender/.nvm/versions/node/v6.10.2/bin/protoc-gen-ts \
--js_out=import_style=commonjs,binary:generated \
--ts_out=service=true:generated \
-I ./proto \
proto/*.proto
