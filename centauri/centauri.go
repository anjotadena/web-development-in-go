package centauri

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/anjotadena/centauri/render"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Centauri struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	Render   *render.Render
	config   config
}

type config struct {
	port     string
	renderer string
}

func (c *Centauri) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	err := c.Init(pathConfig)

	if err != nil {
		return err
	}

	err = c.checkDotEnv(rootPath)

	if err != nil {
		return nil
	}

	err = godotenv.Load(rootPath + "/.env")

	if err != nil {
		return err
	}

	// create loggers
	infoLog, errorLog := c.startLoggers()

	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath
	c.Routes = c.routes().(*chi.Mux)

	c.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	c.createRenderer()

	return nil
}

func (c *Centauri) Init(p initPaths) error {
	root := p.rootPath

	for _, path := range p.folderNames {
		// Create folder if it doesn't exists
		err := c.CreateDirIfNotExist(root + "/" + path)

		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Centauri) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     c.ErrorLog,
		Handler:      c.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	c.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))

	err := srv.ListenAndServe()

	if err != nil {
		c.ErrorLog.Fatal(err)
	}
}

func (c *Centauri) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}

func (c *Centauri) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (c *Centauri) createRenderer() {
	r := render.Render{
		Renderer: c.config.renderer,
		RootPath: c.RootPath,
		Port:     c.config.port,
	}

	c.Render = &r
}
