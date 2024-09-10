#!/bin/bash
set -euxo pipefail

ARCHIVE_LINK=https://github.com/SteamDatabase/Protobufs/archive/master.tar.gz

mkdir -p ./deadlock/tmp

if [ -f tmp/master.tar.gz ]; then
  tar -xzf tmp/master.tar.gz --strip-components=1 -C ./deadlock/tmp
else
  curl -L -o - ${ARCHIVE_LINK} | tar -xz --strip-components=1 -C ./deadlock/tmp
fi

cp -a ./deadlock/tmp/deadlock/*.proto ./deadlock/ && rm -rf ./deadlock/tmp

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
