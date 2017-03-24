# project Noe, Gateway
Software stack for project_noe gateways

## Install

1. Flash SD card with RASPBIAN JESSIE LITE, guide can be found [here](https://www.raspberrypi.org/documentation/installation/installing-images/)

2. Connect to Raspberry PI zero using SSH - username: pi  password: raspberry

	__YOU SHOULD CHANGE PASSWORD FOR pi USER__ or create another user and disable ssh for pi

3. Run the following command to enable SPI interface:
	```
	sudo raspi-config
	```
4. Enable SPI

![raspi-config->Interfacing Options](https://github.com/IoT-BA/project_noe-gateway/raw/master/img/raspi-config_interfacing-options.PNG "Select interfacing options")
![raspi-config->Interfacing Options->SPI](https://github.com/IoT-BA/project_noe-gateway/raw/master/img/raspi-config_interfacing-spi.PNG "Select SPI interface")
![raspi-config->Interfacing Options->SPI->Enable](https://github.com/IoT-BA/project_noe-gateway/raw/master/img/raspi-config_interfacing-spi-enable.PNG "Enable SPI interface")

5. Update and upgrade through apt-get
	```
	sudo apt-get update && sudo apt-get -y upgrade
	```

6. To apply changins RPi needs to reboot
   	```
	sudo reboot
	```

7. Login to RPi again with _new_ password which you set for pi user

8. Install git with apt-get
	```
	sudo apt-get install git -y
	```

9. Install docker
	```
	sudo curl -sSL https://get.docker.com | sh
	```

10. Install docker-compose
	```
	sudo pip install docker-compose
	```

11. Select directory where software for project_noe should reside
	```
	cd /opt
	```

12. Clone project_noe gateway repository
	```
	sudo git clone https://github.com/IoT-BA/project_noe-gateway
	```

13. Descend to project_noe-gateway directory
	```
	cd project_noe-gateway
	```

14. Change VALUES for `###GATEWAY_ID###`, `###MQTT_USER###`, `###MQTT_PASS###` for values which were provided for you during gateway registration

15. Run docker-compose which will run dockerized gateway stack
	```
	docker-compose up --build --force-recreate -d --remove-orphans
	```

