package thor

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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
	Routes   *chi.Mux

	config config
}

type config struct {
	port     string
	renderer string
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
	t.RootPath = rootPath

	// config
	t.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	//routes
	t.Routes = t.routes().(*chi.Mux)

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

// ListenAndServe starts the webserver
func (t *Thor) ListenAndServe() {
	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     t.ErrorLog,
		Handler:      t.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	t.InforLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	if err != nil {
		t.ErrorLog.Fatal(err)
	}
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
