# Swarm ECS

This is a simple Go program that generates a text file suitable for the [swarm file discovery backend](https://github.com/docker/swarm/tree/master/discovery#using-a-static-file-describing-the-cluster)

## Installation

```
go get -u github.com/ejholmes/swarm-ecs/cmd/swarm-ecs
```

## Usage

List the hosts within the cluster:

```console
$ swarm-ecs list
<node_ip1:2375>
<node_ip2:2375>
<node_ip3:2375>
```

Periodically write it to a file that swarm is watching:

```console
$ swarm manage -H tcp://<swarm_ip:swarm_port> file:///tmp/my_cluster &
$ while true; do swarm-ecs list > /tmp/my_cluster; sleep 1; done
```
