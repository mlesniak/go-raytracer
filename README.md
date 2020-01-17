# Overview

This is a simple raytracer based on the book (series) [Raytracing in one weekend](https://raytracing.github.io/). In addition to the book which is rather short on the underlying mathematics I've used the famous, albeit difficult to obtain physically, [An introduction to raytracing](https://www.realtimerendering.com/raytracing/An-Introduction-to-Ray-Tracing-The-Morgan-Kaufmann-Series-in-Computer-Graphics-.pdf) for theoretical concepts of raytracing.

The result is this raytracer written in Go which allows to produce images (which computation took around five minutes with 16 cores) such as

!(full.png)

Using a bit of [Imagemagick](https://imagemagick.org/index.php)'s magic, we are able to create animations (computation time ca. 23 minutes with 16 cores) such as 

!(animation.gif)

## Parallelization support

The program is optimized for multicore machines, albeit we have a strange performance regression, i.e. larger machines with, e.g. 16 or 32 CPUs not all cores are fully utlilized -- which is rather strange given our parallelization approach of computing rows of the image in parallel. 

!(parllel-2.png)
!(parllel-1.png)

## Open topics

[ ] **Cleanup and refactor code**
[ ] Improve parallel performance / fix regression
[ ] Add scene description format
[ ] Add CLI options
[ ] Add support for triangles 
[ ] Allow import (and raytracing) of arbitrary models based on triangles