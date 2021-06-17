package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grafana/grafana/pkg/cmd/grafana-cli/utils"
	"github.com/grafana/grafana/pkg/schema"
	"github.com/grafana/grafana/pkg/schema/load"
	"github.com/sdboyer/cuetsy/encoder"
)

func (cmd Command) generateDashboardTypeScripts(c utils.CommandLine) error {
	dest := c.String("dest")
	from := c.String("from")
	if err := generateTypeScriptFromCUE(from, dest, paths, load.BaseDashboardFamily); err != nil {
		return err
	}

	if err := generateTypeScriptFromCUE(from, dest, paths, load.DistDashboardFamily); err != nil {
		return err
	}

	return nil
}

func generateTypeScriptFromCUE(from string, dest string, p load.BaseLoadPaths, loader func(p load.BaseLoadPaths) (schema.VersionedCueSchema, error)) error {
	dash, err := loader(p)
	if err != nil {
		return err
	}

	inst, _ := dash.CUE().Reference()
	b, err := encoder.Generate(inst, encoder.Config{})
	if err != nil {
		return err
	}
	writeTypeScriptFiles(from, dest, string(b))

	return nil
}

func writeTypeScriptFiles(from string, dest string, content string) error {
	fileName := filepath.Base(from)
	fd, err := os.Create(filepath.Join(dest, fileName[:len(fileName)-3]) + "ts")
	if err != nil {
		return err
	}

	fmt.Fprint(fd, content)
	return nil
}
