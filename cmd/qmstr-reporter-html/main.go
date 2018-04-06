package main

import (
	"log"
	"os"

	"github.com/QMSTR/qmstr/pkg/master"
	"github.com/QMSTR/qmstr/pkg/reporter/htmlreporter"
	"github.com/QMSTR/qmstr/pkg/reporting"
)

func main() {
	reporter := reporting.NewReporter(&htmlreporter.HTMLReporter{Keep: false})
	if err := reporter.RunReporterPlugin(); err != nil {
		log.Printf("%v failed: %v\n", htmlreporter.ModuleName, err)
		os.Exit(master.ReturnReportServiceCommFailed)
	}
	log.Printf("%v completed successfully\n", htmlreporter.ModuleName)
}
