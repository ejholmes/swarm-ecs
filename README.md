# Swarm ECS

This is a simple Go program that generates a text file suitable for the [swarm file discovery backend](https://github.com/docker/swarm/tree/master/discovery#using-a-static-file-describing-the-cluster)

## Usage

Watch the cluster and write to `/tmp/my_cluster`:

```console
swarm-ecs watch --cluster default /tmp/my_cluster
swarm manage -H tcp://<swarm_ip:swarm_port> file:///tmp/my_cluster
```

List the hosts within the cluster:

```console
swarm-ecs list --cluster default
<node_ip1:2375>
<node_ip2:2375>
<node_ip3:2375>
```
