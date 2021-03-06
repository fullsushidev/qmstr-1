package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/QMSTR/qmstr/lib/go-qmstr/service"

	yaml "gopkg.in/yaml.v2"
)

type Analysis struct {
	Name       string
	PosixName  string
	Analyzer   string
	TrustLevel int64
	PathSub    []*service.PathSubstitution
	Config     map[string]string
}

type Reporting struct {
	Name      string
	PosixName string
	Reporter  string
	Config    map[string]string
}

type ServerConfig struct {
	RPCAddress string `yaml:"rpcAddress"`
	DBAddress  string `yaml:"dbAddress"`
	DBWorkers  int `yaml:"dbWorkers"`
	OutputDir  string `yaml:"outputDir"`
	CacheDir   string `yaml:"cacheDir"`
	ImageName  string `yaml:"image"`
	Debug      bool
	ExtraEnv   map[string]string `yaml:"extraEnv"`
	ExtraMount map[string]string `yaml:"extraMount"`
	BuildPath  string `yaml:"buildPath"`
	PathSub    []*service.PathSubstitution `yaml:"pathSub"`
}

type MasterConfig struct {
	Name      string
	MetaData  map[string]string
	Server    *ServerConfig
	Analysis  []*Analysis
	Reporting []*Reporting
}

type QmstrConfig struct {
	Project *MasterConfig
}

func getDefaultConfig() *QmstrConfig {
	return &QmstrConfig{
		Project: &MasterConfig{
			Server: &ServerConfig{DBWorkers: 2, RPCAddress: ":50051", DBAddress: "localhost:9080",
				ExtraEnv: map[string]string{}, ExtraMount: map[string]string{},
			},
		},
	}
}

func ReadConfigFromFiles(configfiles ...string) (*MasterConfig, error) {
	fileNotExistCount := 0
	config := getDefaultConfig()
	for _, configfile := range configfiles {
		if _, err := os.Stat(configfile); os.IsNotExist(err) {
			log.Printf("File %s not found", configfile)
			fileNotExistCount++
			continue
		}

		log.Printf("Reading configuration from %s\n", configfile)
		data, err := ConsumeFile(configfile)
		if err != nil {
			return nil, err
		}

		if err := readConfig(data, config); err != nil {
			return nil, fmt.Errorf("Failed to read config from %s: %v", configfile, err)
		}

	}

	if fileNotExistCount == len(configfiles) {
		return nil, errors.New("No configuration file found")
	}

	return config.Project, nil
}

func ReadConfigFromBytes(data []byte) (*MasterConfig, error) {
	config := getDefaultConfig()
	err := readConfig(data, config)
	if err != nil {
		return nil, err
	}
	return config.Project, err
}

func readConfig(data []byte, configuration *QmstrConfig) error {
	err := yaml.Unmarshal(data, configuration)
	if err != nil {
		return err
	}
	configEnvOverride(configuration.Project)
	err = validateConfig(configuration.Project)
	if err != nil {
		return err
	}
	return nil
}

func SerializeConfig(config *MasterConfig) ([]byte, error) {
	data, err := yaml.Marshal(QmstrConfig{Project: config})
	return data, err
}

func ConsumeFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CreateProjectNode(masterConfig *MasterConfig) *service.ProjectNode {
	projectNode := &service.ProjectNode{Name: masterConfig.Name}
	metaData := &service.InfoNode{Type: "metadata"}
	for key, val := range masterConfig.MetaData {
		metaData.DataNodes = append(metaData.DataNodes, &service.InfoNode_DataNode{Type: key, Data: val})
	}

	if len(metaData.DataNodes) > 0 {
		projectNode.AdditionalInfo = []*service.InfoNode{metaData}
	}

	return projectNode
}

func configEnvOverride(masterConfig *MasterConfig) {
	if dbaddress := os.Getenv("SERVER_DBADDRESS"); dbaddress != "" {
		masterConfig.Server.DBAddress = dbaddress
	}
	if rpcaddress := os.Getenv("SERVER_RPCADDRESS"); rpcaddress != "" {
		masterConfig.Server.RPCAddress = rpcaddress
	}
	if buildpath := os.Getenv("SERVER_BUILDPATH"); buildpath != "" {
		masterConfig.Server.BuildPath = buildpath
	}
}
