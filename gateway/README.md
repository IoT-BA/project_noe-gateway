## Install

Connect to Raspberry PI zero using SSH - username: pi  password: raspberry

Run the following commands:


    sudo raspi-config
    # enable SPI in advanced features
    sudo reboot
    # login to RPI again
    sudo apt-get -y install git
    sudo curl -sSL https://get.docker.com | sh
    sudo pip install docker-compose
    cd /opt
    sudo git clone https://github.com/IoT-BA/ECN.git
    cd ECN/gateway
    docker-compose up --build --force-recreate -d --remove-orphans
