##Lora-Gateway Bridge
Dockerfile for lora-gateway-bridge

###Build image
    sudo docker build -t lora-gateway-bridge:beta -f Dockerfile-gateway_bridge .

###Test-Run container
    sudo docker run --name gateway-bridge -e MQTT_SERVER=tcp://139.59.213.253:1883 -it lora-gateway-bridge:beta

###Daemon-Run container
    sudo docker run --name gateway-bridge lora-gateway-bridge:beta

