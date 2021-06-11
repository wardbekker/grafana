package commands

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"github.com/grafana/grafana/pkg/cmd/grafana-cli/utils"
	"github.com/sdboyer/cuetsy/encoder"
)

func (cmd Command) generateDashboardTypeScripts(c utils.CommandLine) error {
	dest := c.String("dest")
	if err := generateTypeScriptFromFS(paths.BaseCueFS, dest); err != nil {
		return err
	}

	if err := generateTypeScriptFromFS(paths.DistPluginCueFS, dest); err != nil {
		return err
	}

	return nil
}

func generateTypeScriptFromFS(cueFS fs.FS, dest string) error {
	err := fs.WalkDir(cueFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".cue" {
			if err := os.MkdirAll(dest, os.ModePerm); err != nil {
				return err
			}
			fmt.Println("<<<<<<<<<<<<<<<<<1111111111", path)
			b, err := generateTypeScriptFromFile(path)
			if err != nil {
				return err
			}
			fmt.Println("<<<<<<<<<<<<<<<<<222222222222")
			if err := writeTypeScriptFiles(path, dest, b); err != nil {
				return err
			}
		}
		// fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func writeTypeScriptFiles(path string, dest string, content string) error {
	fileName := filepath.Base(path)
	fd, err := os.Create(filepath.Join(dest, fileName[:len(fileName)-3]) + "ts")
	if err != nil {
		return err
	}

	fmt.Fprint(fd, content)
	return nil
}

func generateTypeScriptFromFile(origin string) (string, error) {
	loadedInstances := load.Instances([]string{origin}, nil)
	instances := cue.Build(loadedInstances)
	fmt.Println("<<<<<<<<<<<<<<<<<33")
	b, err := encoder.Generate(instances[0], encoder.Config{})
	fmt.Println("<<<<<<<<<<<<<<<<<44", b)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
