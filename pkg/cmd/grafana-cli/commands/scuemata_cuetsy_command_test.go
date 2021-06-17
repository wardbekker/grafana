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

		err := generateTypeScriptFromCUE(baseLoadPaths, load.BaseDashboardFamily)
		require.NoError(t, err, "error while loading base dashboard scuemata")

		err = generateTypeScriptFromCUE(baseLoadPaths, load.DistDashboardFamily)
		require.NoError(t, err, "error while loading dist dashboard scuemata")
	})
}
