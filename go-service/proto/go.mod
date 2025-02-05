module proto

go 1.22.5

require (
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.1
)

require github.com/google/go-cmp v0.5.9 // indirect

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20220722212130-b98a9ff5e252 // indirect
)

replace (
	// comm => ../../comm
	proto => ./proto
)
