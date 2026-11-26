package traust_test

import (
	"testing"

	"github.com/traust-security/traust-sdk/go/v1/traust"
)

// Regeneration keeps the exported constant names consumers compile against.
func TestEcosystemConstantsKeepTheirNames(t *testing.T) {
	for _, tt := range []struct {
		name string
		got  traust.Ecosystem
		want string
	}{
		{"docker", traust.EcosystemDocker, "docker"},
		{"go", traust.EcosystemGo, "go"},
		{"npm", traust.EcosystemNpm, "npm"},
		{"ruby", traust.EcosystemRuby, "ruby"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.got) != tt.want {
				t.Fatalf("%s = %q", tt.name, tt.got)
			}
		})
	}
}
