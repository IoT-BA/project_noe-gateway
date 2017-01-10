#!/bin/bash
#Script for reseting iC880A over reset jumper which is connected at GPIO 25
echo "25" > /sys/class/gpio/export
echo "out" > /sys/class/gpio/gpio25/direction
echo "1" > /sys/class/gpio/gpio25/value
sleep 5
echo "0" > /sys/class/gpio/gpio25/value
sleep 1
echo "0" > /sys/class/gpio/gpio25/value
echo "25" > /sys/class/gpio/unexport
