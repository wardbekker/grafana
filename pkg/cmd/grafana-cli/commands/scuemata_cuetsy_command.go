package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/grafana/grafana/pkg/cmd/grafana-cli/utils"
	"github.com/grafana/grafana/pkg/schema"
	"github.com/grafana/grafana/pkg/schema/load"
	"github.com/sdboyer/cuetsy/encoder"
)

var generatedFileName = 0

func (cmd Command) generateDashboardTypeScripts(c utils.CommandLine) error {
	dest := c.String("dest")
	if err := generateTypeScriptFromCUE(dest, paths, load.BaseDashboardFamily); err != nil {
		return err
	}

	if err := generateTypeScriptFromCUE(dest, paths, load.DistDashboardFamily); err != nil {
		return err
	}

	return nil
}

func generateTypeScriptFromCUE(dest string, p load.BaseLoadPaths, loader func(p load.BaseLoadPaths) (schema.VersionedCueSchema, error)) error {
	dash, err := loader(p)
	if err != nil {
		return err
	}

	inst, _ := dash.CUE().Reference()
	b, err := encoder.Generate(inst, encoder.Config{})
	if err != nil {
		return err
	}
	writeTypeScriptFiles(dest, string(b))
	return nil
}

func writeTypeScriptFiles(dest string, content string) error {
	fd, err := os.Create(filepath.Join(dest, "generatedFileName"+strconv.Itoa(generatedFileName)+"ts"))
	generatedFileName++
	if err != nil {
		return err
	}

	fmt.Fprint(fd, content)
	return nil
}
