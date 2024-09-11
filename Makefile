SED=sed
ifeq ($(shell uname), Darwin)
	SED=gsed
endif

default: build

test:
	go test -cover -v

testnew:
	go test -cover -run=TestMatchNew -v

bench:
	go test -run=XXX -bench=BenchmarkMatch -benchtime=1m -v

cover:
	go test -cover -coverpkg github.com/ngynkvn/stingray,github.com/ngynkvn/stingray/vbkv -coverprofile /tmp/manta.cov -v
	go tool cover -html=/tmp/manta.cov

cpuprofile:
	go test -v -run=TestMatch2159568145 -test.cpuprofile=/tmp/manta.cpuprof
	go tool pprof -svg -output=/tmp/manta.cpuprof.svg manta.test /tmp/manta.cpuprof
	open /tmp/manta.cpuprof.svg

memprofile:
	go test -v -run=TestMatch2159568145 -test.memprofile=/tmp/manta.memprof -test.memprofilerate=1
	go tool pprof --alloc_space manta.test /tmp/manta.memprof

update: update-protobufs generate

update-protobufs:
	rm -rf deadlock
	./download-protobufs.sh
	$(SED) -i '/^import "network_connection\.proto"/d' deadlock/networkbasetypes.proto
	$(SED) -i '/^import "google\/protobuf\/descriptor\.proto"/d' deadlock/citadel_gameevents.proto
	$(SED) -i '/^import "citadel_gcmessages_common.proto"/d' deadlock/citadel_gamemessages.proto
	$(SED) -i '1isyntax = "proto2";\n\noption go_package = "github.com/ngynkvn/stingray/deadlock;deadlock";\n' deadlock/*.proto
	protoc -I deadlock --go_out=paths=source_relative:deadlock  deadlock/*.proto

generate:
	go run gen/callbacks.go

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
