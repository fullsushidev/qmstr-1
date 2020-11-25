package main

import (
	"context"
	"fmt"
	"github.com/QMSTR/qmstr/lib/go-qmstr/module"
	"github.com/QMSTR/qmstr/lib/go-qmstr/service"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"io"
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
	//files := []*spdx.File2_1{}
	//hashes := []string{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Retrieving all PackageNodes
	// FIXME To delete
	//packageNodes, err := masterClient.GetPackageNodes()
	//if err != nil {
	//	return fmt.Errorf("couldn't get PackageNodes, %v", err)
	//}

	// Response object initialization
	response := &service.FastenResponse{}
	response.PluginName = "QMSTR"
	response.PluginVersion = "0.0.1"
	response.Input = &service.FastenResponse_FastenInput{}
	//response.Input.Product = // TODO fill in
	//response.Input.Forge = // TODO fill in
	//response.Input.Version = // TODO fill in
	response.CreatedAt = ptypes.TimestampNow()
	response.Payload = &service.FastenResponse_Report{Report: &service.FastenReport{
		PackageMetadata: &service.FastenReport_PackageMetadata{},
		FileMetadata:    []*service.FastenReport_FileMetadata{},
	}}

	// Retrieving FileNodes with metadata
	stream, err := masterClient.RptSvcClient.GetAllFileNodesMetadata(ctx, &service.InfoDataRequest{
		Infotype: "license",
		Datatype: "name",
	})
	if err != nil {
		// TODO return FastenError proto message
		return fmt.Errorf("couldn't get FileNodes, %v", err)
	}
	for {
		fileNode, err := stream.Recv()

		// Skipping last and empty FileNodes
		if err == io.EOF || fileNode == nil {
			break
		}

		// Filling file response objects
		switch payload := response.Payload.(type) {
		case *service.FastenResponse_Report:
			payload.Report.FileMetadata = append(payload.Report.FileMetadata,
				&service.FastenReport_FileMetadata{
					FileName: fileNode.GetPath(),
					Sha_1:    fileNode.GetFileData().GetHash(),
					Licenses: []string{fileNode.GetFileData().GetAdditionalInfo()[0].GetDataNodes()[0].GetData()},
				})
		default:
			// TODO return FastenError proto message
			return fmt.Errorf("response initialization error, %v", err)
		}
	}




	//// Retrieving license and compliance information for every PackageNode
	//for _, packageNode := range packageNodes {
	//
	//	// Retrieving all FileNodes of the current package
	//	fileNodes, err := packageNode.GetTargets()
	//	if err != nil {
	//		// TODO return FastenError proto message
	//		return fmt.Errorf("couldn't get FileNodes of package %v:%v, %v",
	//			packageNode.Name, packageNode.Version, err)
	//	}
	//
	//	// Retrieving files metadata
	//	for _, fileNode := range fileNodes {
	//
	//		// Retrieving file licenses
	//		log.Printf("FileNode %v has FileDataNode %v", fileNode.GetUid(), fileNode.GetFileData().GetUid())
	//		licenses, err := masterClient.RptSvcClient.GetInfoData(ctx, &service.InfoDataRequest{
	//			RootID: fileNode.GetFileData().GetUid(), // FIXME empty FileDataNode uid
	//			Infotype: "license",
	//			Datatype: "name",
	//		})
	//		if err != nil {
	//			// TODO return FastenError proto message
	//			return fmt.Errorf("couldn't get license node, %v", err)
	//		}
	//
	//		// Filling file response objects
	//		switch payload := response.Payload.(type) {
	//		case *service.FastenResponse_Report:
	//			payload.Report.FileMetadata = append(payload.Report.FileMetadata,
	//				&service.FastenReport_FileMetadata{
	//					FileName: fileNode.GetPath(),
	//					Sha_1:    fileNode.FileData.GetHash(),
	//					Licenses: licenses.GetData(),
	//				})
	//		default:
	//			// TODO return FastenError proto message
	//			return fmt.Errorf("response initialization error, %v", err)
	//		}
	//	}
	//}

	// Printing JSON to standard output for debugging purposes
	marshaler := jsonpb.Marshaler{}
	formattedResponse, err := marshaler.MarshalToString(response)
	if err != nil {
		return fmt.Errorf("couldn't format response to stdout")
	}
	log.Printf("Result: %v", formattedResponse)

	return nil
}

func (f FastenReporter) PostReport() error {
	return nil
}
