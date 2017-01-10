#!/bin/bash

#possibility of setting global_conf settings by enviroment variables - probably stupid way
sed -i "s/\(\"gateway_ID\":\s\)\"\w*\"/\1\"$PACKET_FORWARDER_GATEWAY_ID\"/g" ./global_conf.json
sed -i "s/\(\"server_address\":\s\)\"\w*\"/\1\"$PACKET_FORWARDER_SERVER_ADDRESS\"/g" ./global_conf.json
sed -i "s/\(\"serv_port_up\":\s\)\w*/\1$PACKET_FORWARDER_SERV_PORT_UP/g" ./global_conf.json
sed -i "s/\(\"serv_port_down\":\s\)\w*/\1$PACKET_FORWARDER_SERV_PORT_DOWN/g" ./global_conf.json
sed -i "s/\(\"keepalive_interval\":\s\)\w*/\1$PACKET_FORWARDER_KEEPALIVE_INTERVAL/g" ./global_conf.json
sed -i "s/\(\"stat_interval\":\s\)\w*/\1$PACKET_FORWARDER_STAT_INTERVAL/g" ./global_conf.json
sed -i "s/\(\"push_timeout_ms\":\s\)\w*/\1$PACKET_FORWARDER_PUSH_TIMEOUT_MS/g" ./global_conf.json
sed -i "s/\(\"forward_crc_valid\":\s\)\w*/\1$PACKET_FORWARDER_FORWARD_CRC_VALID/g" ./global_conf.json
sed -i "s/\(\"forward_crc_error\":\s\)\w*/\1$PACKET_FORWARDER_FORWARD_CRC_ERROR/g" ./global_conf.json
sed -i "s/\(\"forward_crc_disabled\":\s\)\w*/\1$PACKET_FORWARDER_FORWARD_CRC_DISABLED/g" ./global_conf.json
