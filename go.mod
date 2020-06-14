module github.com/roguesoftware/tla-location

go 1.14

replace github.com/roguesoftware/tla-proto => ../tla-proto

require (
	github.com/roguesoftware/tla-proto v0.0.0-20200614165752-7172642b658f
	google.golang.org/grpc v1.29.1
)
