## Packet_forwarder

Dockerfile for packet_forwarder which will run packet_forwarder with default  settings specified as environment variables.
### Build image

`sudo docker build -t packet_forwarder_image:beta -f Dockerfile-packet_forwarder .`
### Test-Run container

`sudo docker run --name forwarder --device /dev/spidev0.0:/dev/spidev0.0 --device /dev/spidev0.1:/dev/spidev0.1 -v /sys/:/sys/ -it packet_forwarder_image:beta`
### Daemon-Run container

`sudo docker run --name forwarder --device /dev/spidev0.0:/dev/spidev0.0 --device /dev/spidev0.1:/dev/spidev0.1 --restart=on-failure -v /sys/:/sys/ packet_forwarder_image:beta`

