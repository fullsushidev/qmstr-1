package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/QMSTR/qmstr/lib/go-qmstr/module"
	"github.com/QMSTR/qmstr/lib/go-qmstr/service"
	"github.com/spdx/tools-golang/v0/spdx"
	"log"
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

	// Keeping track of all files and hash codes
	files := []*spdx.File2_1{}
	hashes := []string{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Retrieving all PackageNodes
	packageNodes, err := masterClient.GetPackageNodes()
	if err != nil {
		return fmt.Errorf("couldn't get PackageNodes, %v", err)
	}

	// Response object
	response := &service.FastenResponse{}
	response.PluginName = "QMSTR"
	response.PluginVersion = "0.0.1"
	response.Input = &service.FastenResponse_FastenInput{}
	//response.Input.Product = // TODO fill in
	//response.Input.Forge = // TODO fill in
	//response.Input.Version = // TODO fill in
	response.CreatedAt = ptypes.TimestampNow()
	response.Payload = &service.FastenResponse_Report{}

	// Retrieving license and compliance information for every PackageNode
	for _, packageNode := range packageNodes {

		// Retrieving package information
		switch payload := response.Payload.(type) {
		case *service.FastenResponse_Report:
			payload.Report.PackageMetadata = &service.FastenReport_PackageMetadata{
				Licenses: nil, // TODO fill in
				Authors:  nil, // TODO fill in
			}
			payload.Report.FileMetadata = []*service.FastenReport_FileMetadata{}
		default:
			// TODO return FastenError proto message
			return fmt.Errorf("response initialization error, %v", err)
		}

		// Retrieving all FileNodes of the current package
		fileNodes, err := packageNode.GetTargets()
		if err != nil {
			// TODO return FastenError proto message
			return fmt.Errorf("couldn't get FileNodes of package %v:%v, %v",
				packageNode.Name, packageNode.Version, err)
		}

		// Retrieving files metadata
		rserv := masterClient.RptSvcClient
		for _, fileNode := range fileNodes {
			licenses, err := rserv.GetInfoData(ctx, &service.InfoDataRequest{
				RootID: fileNode.Uid, Infotype: "license", Datatype: "name"})
			if err != nil {
				// TODO return FastenError proto message
				return fmt.Errorf("couldn't get license node, %v", err)
			}
			switch payload := response.Payload.(type) {
			case *service.FastenResponse_Report:
				payload.Report.FileMetadata = append(payload.Report.FileMetadata,
					&service.FastenReport_FileMetadata{
						FileName: fileNode.Path,
						Sha_1:    fileNode.FileData.Hash,
						Licenses: licenses.GetData(),
					})
			default:
				// TODO return FastenError proto message
				return fmt.Errorf("response initialization error, %v", err)
			}
		}
	}

	// Printing JSON to standard output for debugging purposes
	marshaler := jsonpb.Marshaler{}
	log.Printf("Result: %v", marshaler.MarshalToString(response))

	return nil
}

func (f FastenReporter) PostReport() error {
	return nil
}
