package pb

/*
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/golang/protobuf/protoc-gen-go
*/

// --python_out=plugins=grpc:./../server service.proto
// --go_out=plugins=grpc:./../client service.proto
//go:generate protoc -I. --python_out=plugins=grpc:./../server service.proto
