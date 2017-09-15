# escat - Simple Elasticsearch cat API tool
Have you found yourself repetedly doing `curl -s localhost:9200/_cat/whatever` over and over again? Well, with `escat` you can query [Elasticsearch cat API](https://www.elastic.co/guide/en/elasticsearch/reference/current/cat.html)s with ease and joy.

## Installation
```
go install -u github.com/inkel/escat
```

## Usage
```
Usage: escat [flags] <command>:

  -bytes string
    	use this numeric format instead of 'human' format
  -headers string
    	force only these columns to appear
  -help
    	output available columns
  -verbose
    	use verbose output (i.e. add a header) (default true)


<command> is one of the following:

  aliases        shows information about currently configured aliases to
                 indices including filter and routing infos.
  allocation     provides a snapshot of how many shards are allocated to each
                 data node and how much disk space they are using.
  count          provides quick access to the document count of the entire
                 cluster, or individual indices.
  fielddata      shows how much heap memory is currently being used by
                 fielddata on every data node in the cluster.
  health         allows to get a very simple status on the health of the
                 cluster.
  indices        provides a cross-section of each index; this information
                 spans nodes.
  master         displays the master’s node ID, IP address, and node name.
  nodeattrs      shows custom node attributes.
  nodes          shows the cluster topology.
  pending_tasks  returns a list of any cluster-level changes (e.g. create
                 index, update mapping, allocate or fail shard) which have
                 not yet been executed.
  plugins        provides a view per node of running plugins; this information
                 spans nodes.
  recovery       view of index shard recoveries, both on-going and previously
                 completed.
  repositories   shows the snapshot repositories registered in the cluster.
  thread_pool    shows cluster wide thread pool statistics per node.
  shards         detailed view of what nodes contain which shards.
  segments       provides low level information about the segments in the
                 shards of an index.
  snapshots      shows all snapshots that belong to a specific repository.
```

## License
See [LICENSE](LICENSE).
