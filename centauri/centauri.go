package centauri

import (
	"fmt"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Centauri struct {
	AppName string
	Debug   bool
	Version string
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

func (c *Centauri) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}

	return nil
}
