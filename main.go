package main

import (
	"flag"
	"github.com/fengxizhou/pbslog"
	"path/filepath"
)

var inDirFlag = flag.String("in_dir", "/var/spool/xdmod/palmetto", "the directory that stores the pbs accounting logs")
var outDirFlag = flag.String("out_dir", "/var/spool/xdmod/skylight", "the directory that stores the filtered accounting logs")
var hostPattern = flag.String("host_prefix", "^sky", "a regular expression string for the matched hostnames")

func main() {
	flag.Parse()

	filter := pbslog.NewHostPrefixFilter("^sky")
	filenamePattern := `\d{8}`
	statesStorePath := filepath.Join(*outDirFlag, "states.json")
	lfs := pbslog.NewLogFileStates(*inDirFlag, filenamePattern, statesStorePath)
	pbslog.FilterLogsBatch(lfs, *outDirFlag, filter)
}
