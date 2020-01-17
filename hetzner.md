# Overview

This script contains shell commands to spawn and delete [Hetzner cloud](https://www.hetzner.de/cloud) resources for faster render times.

## Set API key

    brew install hcloud
    hcloud context create default
    hcloud server list

## Spawn new machine

    # Real machines
    # ccx31 ->  8 cores     0.143 EUR/h
    # ccx41 -> 16 cores     0.286 EUR/h ; increase CPU limit with a support ticket
    # ccx51 -> 32 cores     0.571 EUR/h ; increase CPU limit with a support ticket
    
    hcloud server create --image ubuntu-18.04 --name raytracer --type ccx41 --ssh-key m

## Build for linux and copy

    env GOOS=linux GOARCH=amd64 go build
    scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null raytracing root@$(hcloud server list|grep raytracer|cut -b37-50):

## Log into instance and start raytracer

    ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null root@$(hcloud server list|grep raytracer|cut -b37-51)
    ./raytracer
    apt install -y zip
    zip demo.zip demo-*.png
    
## Copy files to local machine

    # scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null root@$(hcloud server list|grep raytracer|cut -b37-50):demo.png demo-$(date +%s).png
    scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null root@$(hcloud server list|grep raytracer|cut -b37-50):demo.zip demo.zip
    
    unzip demo.zip
    ./create-animation.sh
    
## Delete server    
     
    hcloud server delete raytracer
