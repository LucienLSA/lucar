set PROTO_PATH=.\\rental\\api
set GO_OUT_PATH=.\\rental\\api\\gen\\v1
powershell -Command "New-Item -ItemType Directory -Path '%GO_OUT_PATH%' -Force"

protoc -I=%PROTO_PATH% --go_out=paths=source_relative:%GO_OUT_PATH% rental.proto

protoc -I=%PROTO_PATH% --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:%GO_OUT_PATH% rental.proto

protoc -I=%PROTO_PATH% --grpc-gateway_out=paths=source_relative,grpc_api_configuration=%PROTO_PATH%/rental.yaml:%GO_OUT_PATH% rental.proto

set PBTS_BIN_DIR=..\\wx\\miniprogram\\node_modules\\.bin
set PBTS_OUT_DIR=..\\wx\\miniprogram\\service\\proto_gen\\rental

%PBTS_BIN_DIR%\\pbjs -t static -w es6 %PROTO_PATH%\\rental.proto --no-create --no-encode --no-decode --no-verify --no-delimited --force-number -o %PBTS_OUT_DIR%\\rental_pb_tmp.js

echo import * as $protobuf from "protobufjs"; > "%PBTS_OUT_DIR%\rental_pb.js"
type "%PBTS_OUT_DIR%\rental_pb_tmp.js" >> "%PBTS_OUT_DIR%\rental_pb.js"
if exist "%PBTS_OUT_DIR%\rental_pb_tmp.js" del "%PBTS_OUT_DIR%\rental_pb_tmp.js"
"%PBTS_BIN_DIR%\pbts" -o "%PBTS_OUT_DIR%\rental_pb.d.ts" "%PBTS_OUT_DIR%\rental_pb.js"