package thor

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "0.0.1"

// Thor package struct
type Thor struct {
	AppName  string
	IsDebug  bool
	Version  string
	ErrorLog *log.Logger
	InforLog *log.Logger
	RootPath string
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

	// .env
	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	// loggers
	infoLog, errorLog := t.startLoggers()
	t.InforLog = infoLog
	t.ErrorLog = errorLog

	t.IsDebug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	t.Version = version

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

func (t *Thor) startLoggers() (infoLog *log.Logger, errorLog *log.Logger) {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return
}
