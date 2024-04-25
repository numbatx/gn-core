//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/numbatx/protobuf/protobuf --gogoslick_out=. guardians.proto
package guardians
