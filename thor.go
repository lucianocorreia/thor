package thor

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

	return nil
}

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
