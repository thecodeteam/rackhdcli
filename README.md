# RackHDCLI [![Build Status](https://travis-ci.org/codedellemc/rackhdcli.svg?branch=master)](https://travis-ci.org/codedellemc/rackhdcli)
A Go-based CLI for RackHD

## Description
The `rackhdcli` command can be used to interact with [RackHD](https://github.com/RackHD/RackHD).

Currently the CLI supports the **1.1** version of the RackHD API, and can list nodes, SKUs, and tags.

RackHDCLI is built on top of [gorackhd](https://github.com/codedellemc/gorackhd).

## Building

Binaries are not yet provided, but can be built with the following commands:

```
go get github.com/Masterminds/glide
go get -d github.com/codedellemc/rackhdcli
cd $GOPATH/src/github.com/codedellemc/rackhdcli
glide install
go install
```

## Using the CLI

The tool includes robust help via the `--help` flag
```
$ rackhdcli --help
rackhdcli is a command line interface to to interact with a
RackHD server. One can query and modify various components of the RackHD
setup and status

Usage:
  rackhdcli [command]

Available Commands:
  nodes       Interact with RackHD nodes
  skus        Interace with RackHD SKUs
  tags        Interact with tags on RackHD

Flags:
      --endpoint string    API endoint of RackHD (default "localhost:9090")
      --transport string   http or https (default "http")

Use "rackhdcli [command] --help" for more information about a command.
```

The primary option to be aware of is `--endpoint`, which can be used to set the domain and port where your RackHD server is listening.

Example Output:
```
$ rackhdcli --endpoint localhost:9090 skus list
Using config file: /Users/travis/.rackhdcli.yaml
+------------------+--------------------------------------+----------------------------------+
|       NAME       |                  ID                  |        DISCOVERY WORKFLOW        |
+------------------+--------------------------------------+----------------------------------+
| RancherSmallNode | 9c7d91e8-9da9-4c7c-812e-3f89c6778c37 | Graph.DefaultRancherNode.Standby |
| RancherLargeNode | 5202fedc-2f60-409a-8458-f062debab528 | Graph.DefaultRancherNode.Standby |
+------------------+--------------------------------------+----------------------------------+
```

## Licensing
rackhdcli is freely distributed under the [MIT License](http://codedellemc.github.io/sampledocs/LICENSE "LICENSE"). See LICENSE for details.

## Support
Please file bugs and issues on the Github issues page for this project. This is to help keep track and document everything related to this repo. For general discussions and further support you can join the [{code} by Dell EMC Community slack team](http://community.codedellemc.com/) and join the **#rackhd** channel. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
