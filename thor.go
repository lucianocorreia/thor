package thor

import (
	"fmt"

	"github.com/joho/godotenv"
)

const version = "0.0.1"

// Thor package struct
type Thor struct {
	AppName string
	IsDebug bool
	Version string
}

func (t *Thor) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers", "migrations", "views", "models", "public", "tmp", "logs", "middlewares",
		},
	}

	err := t.Init(pathConfig)
	if err != nil {
		return err
	}

	err = t.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	return nil
}

// Init creates the defaul structure for the app
func (t *Thor) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := t.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Thor) checkDotEnv(path string) error {
	err := t.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}

	return nil
}
