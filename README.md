# project Noe, Gateway
Software stack for project_noe gateways

## Install

1. Flash SD card with RASPBIAN JESSIE LITE, guide can be found [here](https://www.raspberrypi.org/documentation/installation/installing-images/)

2. Connect to Raspberry PI zero using SSH - username: pi  password: raspberry
	_YOU SHOULD CHANGE PASSWORD FOR pi USER_ or create another user and disable ssh for pi

3. Run the following commands:
	```
    sudo raspi-config
    ```
4. Enable SPI
![raspi-config->config_interfacing]( "Logo Title Text 1")
    sudo reboot
    # login to RPI again
    sudo apt-get -y install git
    sudo curl -sSL https://get.docker.com | sh
    sudo pip install docker-compose
    cd /opt
    sudo git clone https://github.com/IoT-BA/ECN.git
    cd ECN/gateway
    docker-compose up --build --force-recreate -d --remove-orphans
