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
	rm -rf \
		deadlock/base_gcmessages.proto \
		deadlock/c_peer2peer_netmessages.proto \
		deadlock/citadel_clientmessages.proto \
		deadlock/citadel_gcmessages_client.proto \
		deadlock/citadel_gcmessages_server.proto \
		deadlock/citadel_usercmd.proto \
		deadlock/clientmessages.proto \
		deadlock/connectionless_netmessages.proto \
		deadlock/econ_gcmessages.proto \
		deadlock/econ_shared_enums.proto \
		deadlock/engine_gcmessages.proto \
		deadlock/enums_clientserver.proto \
		deadlock/gcsystemmsgs.proto \
		deadlock/network_connection.proto \
		deadlock/networksystem_protomessages.proto \
		deadlock/steamdatagram_messages_auth.proto \
		deadlock/steamdatagram_messages_sdr.proto \
		deadlock/steammessages_base.proto \
		deadlock/steammessages_cloud.steamworkssdk.proto \
		deadlock/steammessages_gamenetworkingui.proto \
		deadlock/steammessages_helprequest.steamworkssdk.proto \
		deadlock/steammessages_int.proto \
		deadlock/steammessages_oauth.steamworkssdk.proto \
		deadlock/steammessages_player.steamworkssdk.proto \
		deadlock/steammessages_publishedfile.steamworkssdk.proto \
		deadlock/steamnetworkingsockets_messages.proto \
		deadlock/steamnetworkingsockets_messages_certs.proto \
		deadlock/steamnetworkingsockets_messages_udp.proto \
		deadlock/uifontfile_format.proto \
		deadlock/usercmd.proto
	$(SED) -i '/^import "network_connection\.proto"/d' deadlock/networkbasetypes.proto
	$(SED) -i '/^import "google\/protobuf\/descriptor\.proto"/d' deadlock/citadel_gameevents.proto
	$(SED) -i '/^import "citadel_gcmessages_common.proto"/d' deadlock/citadel_gamemessages.proto
	$(SED) -i '1isyntax = "proto2";\n\noption go_package = "github.com/ngynkvn/stingray/deadlock;deadlock";\n' deadlock/*.proto
	protoc -I deadlock --go_out=paths=source_relative:deadlock  deadlock/*.proto

generate:
	go run gen/callbacks.go

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
