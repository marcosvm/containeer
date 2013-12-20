# Containeer

Containeer is a tool to create multiple containers on the [Rackspace CloudFiles storage service](http://www.rackspace.com/cloud/files/) using go concurrency primitives

As is known Cloud Files will throttle you in 4 req/s if the number of objects in a container goes over 500,000 which is not [*web scale*](https://www.youtube.com/watch?v=b2F-DItXtZs)

To overcame this limitation one solution is to create multiple containers and keep the number of objects below 500,000.

## Setup

```bash
go install github.com/marcosvm/containeer
```

## Usage

You must define the following enviroment variables for the connection

```bash
SWIFT_API_USER
SWIFT_API_KEY
SWIFT_AUTH_URL
```

```bash
Usage of ./containeer:
  -concurrency=50: how many concurrent requests
  -cpuprofile="": write cpu profile to file
  -dry=false: dry run, won't create any container
  -list=false: prints a list of existent containers and exits
  -list_filter="": a filter to list the containers
  -num=10000: number of containers to create
  -prefix="development_": prefix for the containers names
  -single="": create a single container
```

## Example creating containers

```bash
SWIFT_API_USER=your_api_user \
SWIFT_API_KEY=your_api_key \
SWIFT_AUTH_URL=https://lon.auth.api.rackspacecloud.com/v2.0 \
containeer -num=10000 -prefix="production_"
```

The command above will create 10000 containers on the London endpoint using production_ as a prefix using 50 concurrent requests.

Result is:
```
production_00001 production_00002 ... production_10000 are created
```

## Example listing containers

```bash
SWIFT_API_USER=your_api_user \
SWIFT_API_KEY=your_api_key \
SWIFT_AUTH_URL=https://lon.auth.api.rackspacecloud.com/v2.0 \
containeer -list -list_filter=production_09000
```

The command above will print existent containers *above* production_09000

A list is printed:
```
production_09000 production_09001 ... production_10000
```

## Example creating a single container
```bash
SWIFT_API_USER=your_api_user \
SWIFT_API_KEY=your_api_key \
SWIFT_AUTH_URL=https://lon.auth.api.rackspacecloud.com/v2.0 \
containeer -single="single_container"
```

## To do

* more tests
* documentation (go docs)

## Contributing

Contributions are welcome, create a GitHub issue or even better, a *Pull Request*, with code and tests.

(c) 2013 Marcos Oliveira
