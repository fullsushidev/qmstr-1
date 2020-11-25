package module

import (
	"github.com/QMSTR/qmstr/lib/go-qmstr/service"
)

type FileNodeProxy struct {
	service.FileNode
	masterClient *MasterClient
}

//func (m *MasterClient) GetAllFileNodesMetadata() ([]*FileNodeProxy, error) {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	fnps := []*FileNodeProxy{}
//	var fileNodeStream, err = m.RptSvcClient.GetAllFileNodesMetadata(ctx, &service.InfoDataRequest{
//		Infotype: "license",
//		Datatype: "name",
//	})
//	if err != nil {
//		return nil, fmt.Errorf("Couldn't get file node metadata, %v", err)
//	}
//
//	for {
//		fileNode, err := fileNodeStream.Recv()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			return nil, fmt.Errorf("Error receiving file nodes metadata, %v", err)
//		}
//		fnps = append(fnps, &FileNodeProxy{*fileNode, m})
//	}
//	return fnps, nil
//}
