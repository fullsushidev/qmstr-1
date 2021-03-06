package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"syscall"

	"golang.org/x/net/context"

	"github.com/QMSTR/qmstr/lib/go-qmstr/analysis"
	"github.com/QMSTR/qmstr/lib/go-qmstr/master"
	"github.com/QMSTR/qmstr/lib/go-qmstr/module"
	"github.com/QMSTR/qmstr/lib/go-qmstr/service"
)

const (
	scancodeErrorCount = "Errors count:\\s*(\\d+)"
)

type ScancodeAnalyzer struct {
	scanData interface{}
}

func main() {
	analyzer := analysis.NewAnalyzer(&ScancodeAnalyzer{})
	if err := analyzer.RunAnalyzerModule(); err != nil {
		log.Printf("%v failed: %v\n", analyzer.GetModuleName(), err)
		os.Exit(master.ReturnAnalyzerFailed)
	}
}

func (scanalyzer *ScancodeAnalyzer) Configure(configMap map[string]string) error {
	cached := false
	if usecache, ok := configMap["cached"]; ok && usecache == "true" {
		cached = true
	}

	resultfile := ""
	if rf, ok := configMap["resultfile"]; ok {
		resultfile = rf
	}

	var err error
	if cached && resultfile != "" {
		scanalyzer.scanData, err = readScancodeFile(resultfile)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	if workdir, ok := configMap["workdir"]; ok {
		scanalyzer.scanData, err = scancode(workdir, runtime.NumCPU(), resultfile)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	return fmt.Errorf("Misconfigured scancode analyzer")
}

func (scanalyzer *ScancodeAnalyzer) Analyze(masterClient *module.MasterClient, token int64) error {
	stream, err := masterClient.AnaSvcClient.GetSourceFileNodes(context.Background(), &service.DummyRequest{})
	if err != nil {
		log.Printf("failed getting source file nodes %v", err)
		return err
	}

	infoNodeMsgs := []*service.InfoNodesMessage{}

	for {
		fileNode, err := stream.Recv()
		if err == io.EOF {
			break
		}

		log.Printf("Analyzing file %s", fileNode.Path)

		infoNodes := []*service.InfoNode{}

		licenseInfo, err := scanalyzer.detectLicenses(fileNode.GetPath())
		if err == nil {
			infoNodes = append(infoNodes, licenseInfo...)
		}
		copyrights, err := scanalyzer.detectCopyrights(fileNode.GetPath())
		if err == nil {
			infoNodes = append(infoNodes, copyrights)
		}

		if len(infoNodes) > 0 {
			infoNodeMsgs = append(infoNodeMsgs, &service.InfoNodesMessage{Token: token, Infonodes: infoNodes, Uid: fileNode.FileData.Uid})
		}
	}

	sendStream, err := masterClient.AnaSvcClient.SendInfoNodes(context.Background())
	if err != nil {
		return err
	}
	for _, inodeMsg := range infoNodeMsgs {
		sendStream.Send(inodeMsg)
	}

	reply, err := sendStream.CloseAndRecv()
	if err != nil {
		return err
	}
	if reply.Success {
		log.Println("Scancode Analyzer sent InfoNodes")
	}

	return nil
}

func (scanalyzer *ScancodeAnalyzer) PostAnalyze() error {
	return nil
}

func scancodeExitHandler(err error, output []byte) {
	// print to stdout so master server can log it
	fmt.Print(string(output))
	if output != nil {
		// scancode might have failed, let's see ...
		re := regexp.MustCompile(scancodeErrorCount)
		errors := re.FindSubmatch(output)
		if len(errors) > 1 && string(errors[1]) == "0" {
			return
		}
	}
	if err != nil {
		log.Printf("Scancode-toolkit failed %s\n", err)
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				// preserve non-zero return code
				os.Exit(status.ExitStatus())
			}
		}
		log.Fatalln("scancode return code was not found. Something went seriously wrong.")
	}
}

func readScancodeFile(resultfile string) (interface{}, error) {
	re := regexp.MustCompile("{.+")
	scanResult, err := ioutil.ReadFile(resultfile)
	if err != nil {
		return nil, err
	}
	jsonScanResult := re.Find(scanResult)
	var scanData interface{}
	err = json.Unmarshal(jsonScanResult, &scanData)
	if err != nil {
		return nil, err
	}
	return scanData, nil
}

func scancode(workdir string, jobs int, resultfilepath string) (interface{}, error) {
	if resultfilepath == "" {
		tmpfile, err := ioutil.TempFile("", "qmstr-scancode")
		resultfilepath = tmpfile.Name()
		if err != nil {
			return nil, err
		}
		defer os.Remove(tmpfile.Name())
	}

	cmdlineargs := []string{"--full-root", "--license", "--copyright", "--json", resultfilepath}
	if jobs > 1 {
		cmdlineargs = append(cmdlineargs, "--processes", fmt.Sprintf("%d", jobs))
	}
	cmd := exec.Command("scancode", append(cmdlineargs, workdir)...)
	log.Printf("Calling %s\n", cmd.Args)

	output, err := cmd.CombinedOutput()
	scancodeExitHandler(err, output)

	return readScancodeFile(resultfilepath)
}

func (scanalyzer *ScancodeAnalyzer) detectLicenses(srcFilePath string) ([]*service.InfoNode, error) {
	licenseNodes := []*service.InfoNode{}
	scanDatamap := scanalyzer.scanData.(map[string]interface{})
	for _, file := range scanDatamap["files"].([]interface{}) {
		fileData := file.(map[string]interface{})
		if fileData["path"] == srcFilePath {
			log.Printf("Found %s", srcFilePath)
			for _, licenseData := range fileData["licenses"].([]interface{}) {
				license := licenseData.(map[string]interface{})
				log.Println("Found license information")
				tempDataNodes := []*service.InfoNode_DataNode{{
					Type: "key",
					Data: license["key"].(string),
				},
					{
						Type: "score",
						Data: strconv.FormatFloat(license["score"].(float64), 'f', 2, 64),
					},
				}

				name := license["short_name"]
				if name != nil {
					tempDataNodes = append(tempDataNodes, &service.InfoNode_DataNode{
						Type: "name",
						Data: name.(string),
					})
				}

				spdxIdent := license["spdx_license_key"]
				if spdxIdent != nil {
					tempDataNodes = append(tempDataNodes, &service.InfoNode_DataNode{
						Type: "spdxIdentifier",
						Data: spdxIdent.(string),
					})
				}
				licenseNodes = append(licenseNodes, &service.InfoNode{
					Type:            "license",
					ConfidenceScore: license["score"].(float64) / 100,
					DataNodes:       tempDataNodes,
				})
			}
			return licenseNodes, nil
		}
	}
	return nil, fmt.Errorf("No license found for %s", srcFilePath)
}

func (scanalyzer *ScancodeAnalyzer) detectCopyrights(srcFilePath string) (*service.InfoNode, error) {
	copyrights := []*service.InfoNode_DataNode{}
	scanDatamap := scanalyzer.scanData.(map[string]interface{})
	for _, file := range scanDatamap["files"].([]interface{}) {
		fileData := file.(map[string]interface{})
		if fileData["path"] == srcFilePath {
			for _, copyright := range fileData["holders"].([]interface{}) {
				copyrightData := copyright.(map[string]interface{})
				log.Printf("Found copyright holder %v", copyrightData["value"])
				copyrights = append(copyrights, &service.InfoNode_DataNode{
					Type: "copyrightHolder",
					Data: copyrightData["value"].(string),
				})
			}
			for _, author := range fileData["authors"].([]interface{}) {
				authorData := author.(map[string]interface{})
				log.Printf("Found author %v", authorData["value"])
				copyrights = append(copyrights, &service.InfoNode_DataNode{
					Type: "author",
					Data: authorData["value"].(string),
				})
			}
			if len(copyrights) > 0 {
				copyrightInfoNode := &service.InfoNode{
					Type:      "copyright",
					DataNodes: copyrights,
				}
				return copyrightInfoNode, nil
			}
		}
	}
	return nil, fmt.Errorf("No copyright info found for %s", srcFilePath)
}
