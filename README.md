### Building

Run the following to build the image, tag it, and push it out to dockerhub

```
make package tag push
```

### Running

Output fromt the `--health` flag

```
A simple healthcheck webserver. It expects an inbound ping within the heartbeat interval.

Usage:
  webhealth [flags]

Flags:
      --grace int       number of intervals that can be missed before considered unhealthy (default 3)
      --heartbeat int   heartbeat interval in seconds (default 10)
```
