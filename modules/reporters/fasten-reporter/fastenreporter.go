package main

import (
	"context"
	"fmt"
	"github.com/QMSTR/qmstr/lib/go-qmstr/module"
	"github.com/QMSTR/qmstr/lib/go-qmstr/service"
)

const (
	ModuleName  = "reporter-fasten"
	outFileName = "%s.spdx"
)

type FastenReporter struct {
}

func (f FastenReporter) Configure(configMap map[string]string) error {
	return nil
}

func (f FastenReporter) Report(masterClient *module.MasterClient) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	infoData, err := masterClient.RptSvcClient.GetInfoData(ctx, &service.InfoDataRequest{})
	if err != nil {
		return fmt.Errorf("failed to get info data: %v", err)
	}

	fmt.Printf("Data: %v", infoData)

	return nil
}

func (f FastenReporter) PostReport() error {
	return nil
}
