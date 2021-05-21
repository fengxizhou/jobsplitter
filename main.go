package main

import (
	"flag"
	"github.com/fengxizhou/pbslog"
	"log"
	"os"
	"path/filepath"
)

var inDirFlag = flag.String("in_dir", "/var/spool/xdmod/palmetto/accounting", "the directory that stores the pbs accounting logs")
var outDirFlag = flag.String("out_dir", "/var/spool/xdmod/skylight/accounting", "the directory that stores the filtered accounting logs")
var hostPattern = flag.String("host_prefix",	"", "a regular expression string for the matched hostnames")
var queuePattern = flag.String("queue_prefix", "", "a regular expression string for the matched queues")
var logFile = flag.String("log_path", "/var/log/jobsplitter.log", "the log file path")

func main() {
	var (
		logFH *os.File = nil
		logger = log.Default()
	)
	logger.SetPrefix("pbslog: ")
	logger.SetFlags(log.Ldate | log.Ltime)
	if *logFile != "" {
		var err error
		logFH, err = os.OpenFile(*logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			logger.Fatal(err)
		}
		logger.SetOutput(logFH)
	}
	defer func(){
		if logFH != nil {
			logFH.Close()
		}
	}()
	
	flag.Parse()

	filenamePattern := `\d{8}`
	statesStorePath := filepath.Join(filepath.Dir(*outDirFlag), "states.json")
	lfs := pbslog.NewLogFileStates(*inDirFlag, filenamePattern, statesStorePath)

	var filters []pbslog.JobFilter
	if *hostPattern != "" {
		filters = append(filters, pbslog.NewHostPrefixFilter(*hostPattern))
	}
	if *queuePattern != "" {
		filters = append(filters, pbslog.NewQueueFilter(*queuePattern))
	}
	if len(filters) >= 1 {
		pbslog.FilterLogsBatch(lfs, *outDirFlag, filters)
	}
}
