# Overview

This is a simple raytracer based on the book (series) [Raytracing in one weekend](https://raytracing.github.io/). In addition to the book which is rather short on the underlying mathematics I've used the famous, albeit difficult to obtain physically, [An introduction to raytracing](https://www.realtimerendering.com/raytracing/An-Introduction-to-Ray-Tracing-The-Morgan-Kaufmann-Series-in-Computer-Graphics-.pdf) for theoretical concepts of raytracing.

The result is this raytracer written in Go which allows to produce images such as

!(full.png)

Using a bit of [Imagemagick](https://imagemagick.org/index.php)'s magic, we are able to create animations such as 

TOOD

hcloud context create default
hcloud server list

# Real machines
# ccx31 ->  8 cores     0.143 EUR/h
# ccx41 -> 16 cores     0.286 EUR/h
# ccx51 -> 32 cores     0.571 EUR/h

hcloud server create --image ubuntu-18.04 --name raytracer --type ccx51 --ssh-key m


env GOOS=linux GOARCH=amd64 go build
scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null raytracing root@$(hcloud server list|grep raytracer|cut -b37-50):

ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null root@$(hcloud server list|grep raytracer|cut -b37-51)

scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null root@$(hcloud server list|grep raytracer|cut -b37-50):demo.png demo-$(date +%s).png 
hcloud server delete raytracer


## Open topics

[ ] Cleanup and refactor code
[ ] Add scene description format
[ ] Add CLI options
[ ] Add support for triangles 
[ ] Allow import (and raytracing) of arbitrary models