package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const commands = `

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

More info at https://www.elastic.co/guide/en/elasticsearch/reference/current/cat.html
`

func main() {
	var (
		verbose = flag.Bool("v", true, "use verbose output (i.e. add a header)")
		help    = flag.Bool("help", false, "output available columns")
		headers = flag.String("h", "", "force only these columns to appear")
		bytes   = flag.String("bytes", "", "use this numeric format instead of 'human' format")
		host    = flag.String("H", "http://localhost:9200", "Elasticsearch host to connect to. Must include the protocol.")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <command>:\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, commands)
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	qs := url.Values{}

	if *help {
		qs.Set("help", "true")
	} else {
		if *verbose {
			qs.Set("v", "1")
		}
		if *headers != "" {
			qs.Set("h", *headers)
		}
		if *bytes != "" {
			qs.Set("bytes", *bytes)
		}
	}

	command := flag.Arg(0)

	base, err := url.Parse(*host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse %v as host: %v\n", *host, err)
		os.Exit(4)
	}
	if base.Scheme != "http" && base.Scheme != "https" {
		fmt.Fprintf(os.Stderr, "unsupported protocol: %v\n", base.Scheme)
		os.Exit(4)
	}

	esURL := base.String() + fmt.Sprintf("/_cat/%s?%s", command, qs.Encode())

	res, err := http.Get(esURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v\n", err)
		os.Exit(3)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
