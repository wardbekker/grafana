package commands

import (
	"testing"

	"github.com/grafana/grafana/pkg/schema/load"
	"github.com/stretchr/testify/require"
)

func TestCuetsyBasics(t *testing.T) {
	t.Run("Testing generate ts with cue schemas", func(t *testing.T) {
		var baseLoadPaths = load.BaseLoadPaths{
			BaseCueFS:       defaultBaseLoadPaths.BaseCueFS,
			DistPluginCueFS: defaultBaseLoadPaths.DistPluginCueFS,
		}

		err := generateTypeScriptFromCUE("testdata/", baseLoadPaths, load.BaseDashboardFamily)
		require.NoError(t, err, "error while generating base dashboard scuemata")

		err = generateTypeScriptFromCUE("testdata/", baseLoadPaths, load.DistDashboardFamily)
		require.NoError(t, err, "error while generating dist dashboard scuemata")
	})
}
