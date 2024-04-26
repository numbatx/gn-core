//go:generate protoc -I=. -I=$GOPATH/src/github.com/numbatx/gn-core/data/block -I=$GOPATH/src -I=$GOPATH/src/github.com/numbatx/protobuf/protobuf --gogoslick_out=$GOPATH/src firehoseBlock.proto

package firehose
