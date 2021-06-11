package commands

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/require"
)

func TestCuetsyBasics(t *testing.T) {
	t.Run("Testing generate typeScripts from cue schemas", func(t *testing.T) {
		genCue, err := os.ReadFile("testdata/dashboard_cuetsy.cue")
		require.NoError(t, err)

		filesystem := fstest.MapFS{
			"cue/ui/gen.cue": &fstest.MapFile{Data: genCue},
		}

		path, err := os.Getwd()
		require.NoError(t, err)

		err = generateTypeScriptFromFS(filesystem, filepath.Join(path, "testdata"))

		// err = validateScuemata(baseLoadPaths, load.BaseDashboardFamily)
		require.NoError(t, err, "error when generate typeScript for base dashboard scuemata")

		// err = validateScuemata(baseLoadPaths, load.DistDashboardFamily)
		// assert.EqualError(t, err, "all schema should be valid with respect to basic CUE rules, Family.lineages.0.0: field #Panel not allowed")
	})
}
