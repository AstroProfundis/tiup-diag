package main

import (
	"flag"

	"github.com/pingcap/tidb-foresight/bootstrap"
	"github.com/pingcap/tidb-foresight/server"
	"github.com/pingcap/tidb-foresight/version"
	log "github.com/sirupsen/logrus"
)

func main() {
	printVersion := flag.Bool("V", false, "print version info")
	homepath := flag.String("home", "/tmp/tidb-foresight", "tidb-foresight work home")
	address := flag.String("address", "0.0.0.0:8888", "tidb foresight listen address")

	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if *printVersion {
		version.PrintReleaseInfo()
		return
	}

	config, db := bootstrap.MustInit(*homepath)
	defer db.Close()

	s := server.NewServer(config, db)

	log.Panic(s.Run(*address))
}
