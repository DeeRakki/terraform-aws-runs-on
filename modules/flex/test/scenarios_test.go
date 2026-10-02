package test

import (
	"context"
	"testing"

	"github.com/runs-on/terraform-aws-runs-on/modules/flex/test/internal/validationimage"
)

func TestScenarioMatrix(t *testing.T) {
	testCases := []struct {
		name          string
		validationEnv string
		skipInShort   bool
		configure     func(*ScenarioConfig)
	}{
		{
			name:          "basic",
			validationEnv: validationimage.ScenarioBasic,
			configure:     func(*ScenarioConfig) {},
		},
		{
			name:          "private",
			validationEnv: validationimage.ScenarioPrivate,
			skipInShort:   true,
			configure: func(cfg *ScenarioConfig) {
				cfg.EnableNAT = true
				cfg.PrivateMode = "true"
				cfg.EnableCacheIsolation = true
				cfg.EnableStickyDiskIsolation = true
			},
		},
		{
			name:          "full",
			validationEnv: validationimage.ScenarioFull,
			skipInShort:   true,
			configure: func(cfg *ScenarioConfig) {
				cfg.EnableNAT = true
				cfg.EnableEFS = true
				cfg.EnableECR = true
				cfg.EnableCacheIsolation = true
				cfg.EnableStickyDiskIsolation = true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.skipInShort && testing.Short() {
				t.Skip("Skipping expensive scenario in short mode")
			}

			requireValidationEnv(t, tc.validationEnv, validationimage.EnvModeDirect)

			cfg := DefaultScenarioConfig()
			tc.configure(&cfg)

			result := deployScenario(t, cfg)
			t.Logf("Stack %s deployed; ingress %s", result.StackName(), result.IngressURL())

			runBaselineValidations(t, NewAWSClients(context.Background()), result)
		})
	}
}
