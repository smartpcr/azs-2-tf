package mocks

import (
	"github.com/smartpcr/azs-2-tf/utils"
	"os"
	"path/filepath"
)

var (
	_ utils.Settings = &MockSettings{}
)

type MockSettings struct {
	appName          string
	appFolderPath    string
	configFolderPath string
	configFileName   string
	logFolderPath    string
	logFileName      string
}

func NewMockSettings() *MockSettings {
	testFolder, err := filepath.Abs("./unit_test")
	if err != nil {
		panic(err)
	}

	if _, e := os.Stat(testFolder); os.IsNotExist(e) {
		err := os.Mkdir(testFolder, 0700)
		if err != nil {
			panic(err)
		}
	}

	return &MockSettings{
		appName:          "azs-2-tf",
		appFolderPath:    testFolder,
		configFolderPath: filepath.Join(testFolder, "config"),
		configFileName:   "config.json",
		logFolderPath:    filepath.Join(testFolder, "logs"),
		logFileName:      "azs-2-tf.log",
	}
}

func (m *MockSettings) GetAppName() string {
	return m.appName
}

func (m *MockSettings) GetAppFolderPath() string {
	return m.appFolderPath
}

func (m *MockSettings) GetConfigFolderPath() string {
	return m.configFolderPath
}

func (m *MockSettings) GetConfigFileName() string {
	return m.configFileName
}

func (m *MockSettings) GetLogFolderPath() string {
	return m.logFolderPath
}

func (m *MockSettings) GetLogFileName() string {
	return m.logFileName
}
