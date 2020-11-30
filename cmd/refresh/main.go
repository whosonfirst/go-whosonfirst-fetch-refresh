package main

import (
	_ "github.com/whosonfirst/go-reader-http"
	_ "github.com/whosonfirst/go-reader-whosonfirst-data"
)

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"github.com/whosonfirst/go-whosonfirst-fetch"
	"github.com/whosonfirst/go-whosonfirst-index"
	_ "github.com/whosonfirst/go-whosonfirst-index/fs"
	"github.com/whosonfirst/go-whosonfirst-uri"
	"github.com/whosonfirst/go-writer"
	"io"
	"log"
	"os"
)

func main() {

	reader_uri := flag.String("reader-uri", "whosonfirst-data://", "A valid go-reader URI.")
	writer_uri := flag.String("writer-uri", "null://", "A valid go-writer URI.")
	indexer_uri := flag.String("indexer-uri", "repo://", "A valid go-whosonfirst-index URI.")
	retries := flag.Int("retries", 3, "The maximum number of attempted retries when fetching a record.")

	dryrun := flag.Bool("dryrun", false, "List all the records to be refreshed (but don't refresh them).")

	var belongs_to flags.MultiString
	flag.Var(&belongs_to, "belongs-to", "One or more placetypes that a given ID may belong to to also fetch. You may also pass 'all' as a short-hand to fetch the entire hierarchy for a place.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Refresh an index of Who's On First records.\n\n")		
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [options] [path1 path2 ... pathN]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r, err := reader.NewReader(ctx, *reader_uri)

	if err != nil {
		log.Fatal(err)
	}

	wr, err := writer.NewWriter(ctx, *writer_uri)

	if err != nil {
		log.Fatal(err)
	}

	fetcher_opts, err := fetch.DefaultOptions()

	if err != nil {
		log.Fatal(err)
	}

	fetcher_opts.Retries = *retries

	fetcher, err := fetch.NewFetcher(ctx, r, wr, fetcher_opts)

	if err != nil {
		log.Fatal(err)
	}

	cb := func(ctx context.Context, fh io.Reader, args ...interface{}) error {

		path, err := index.PathForContext(ctx)

		if err != nil {
			return err
		}

		id, _, err := uri.ParseURI(path)

		if err != nil {
			log.Printf("Failed to parse '%s', %v\n", path, err)
			return nil
		}

		if *dryrun {
			log.Printf("Refresh %d (%s)\n", id, path)
			return nil
		}

		err = fetcher.FetchIDs(ctx, []int64{id}, belongs_to...)

		if err != nil {
			return err
		}

		return nil
	}

	i, err := index.NewIndexer(*indexer_uri, cb)

	if err != nil {
		log.Fatalf("Failed to create new indexer, %v", err)
	}

	paths := flag.Args()

	err = i.Index(ctx, paths...)

	if err != nil {
		log.Fatalf("Failed to index paths, %v", err)
	}

}
