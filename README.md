# go-whosonfirst-fetch-refresh

Tools for refreshing an index of Who's On First records.

## Tools

### refresh

```
$> ./bin/refresh -h
Refresh an index of Who's On First records.

Usage:
  ./bin/refresh [options] [path1 path2 ... pathN]

Options:
  -belongs-to value
    	One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.
  -dryrun
    	List all the records to be refreshed (but don't refresh them).
  -indexer-uri string
    	A valid go-whosonfirst-index URI. (default "repo://")
  -reader-uri string
    	A valid go-reader URI. (default "whosonfirst-data://")
  -retries int
    	The maximum number of attempted retries when fetching a record. (default 3)
  -writer-uri string
    	A valid go-writer URI. (default "null://")
```

Refresh an index of Who's On First records. For example the [sfomuseum-data-whosonfirst](https://github.com/sfomuseum-data/sfomuseum-data-whosonfirst) repository which contains a subset of the larger WOF data set and that are updated with SFO Museum-specific properties. The data in the `sfomuseum-data-whosonfirst` repository is "refreshed" like this:

```
$> ./bin/refresh \
	-reader-uri whosonfirst-data:// \
	-writer-uri fs:///usr/local/data/sfomuseum-data-whosonfirst/data/ \
	-indexex-uri repo:// \
	/usr/local/data/sfomuseum-data-whosonfirst/
```

_Note: SFO Museum-specific properties are appended after the fact by another process._

Under the hood this tool uses a number of other packages for handling specific tasks:

* It uses the `go-whosonfirst-indexer` package for determining which records to refresh.
* It uses the `go-whosonfirst-fetch` package for actually retrieving records, and their ancestors.
* It uses the `go-reader` package to read Who's On First records (that will update the index).
* It uses the `go-writer` package to write the updated Who's On First records (in the index).

### Readers

Readers exported by the following packages are available to the `refresh` tool:

* https://github.com/whosonfirst/go-reader
* https://github.com/whosonfirst/go-reader-whosonfirst-data
* https://github.com/whosonfirst/go-reader-http
* https://github.com/whosonfirst/go-reader-github

### Writers

Writers exported by the following packages are available to the `refresh` tool:

* https://github.com/whosonfirst/go-writer

### Indexers

Indexers exported by the following packages are available to the `refresh` tool:

* https://github.com/whosonfirst/go-whosonfirst-index

## See also

* https://github.com/whosonfirst/go-whosonfirst-fetch
* https://github.com/whosonfirst/go-whosonfirst-indexer
* https://github.com/whosonfirst/go-reader-whosonfirst-data
* https://github.com/whosonfirst/go-reader
* https://github.com/whosonfirst/go-writer
