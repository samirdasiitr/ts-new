export PATH=$PATH:/home/ubuntu/work/bin/
export GOPATH=/home/ubuntu/work

go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/jsonpb
go get github.com/gogo/protobuf/protoc-gen-gogo
go get github.com/gogo/protobuf/gogoproto

#protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto \
#   --gogo_out=plugins=grpc,paths=source_relative,\
#Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
#Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
#Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
#Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
#Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:. \
#model/user.proto

protoc -I=. -I=${GOPATH}/src --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative model/user.proto
