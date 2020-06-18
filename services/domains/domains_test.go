package domains

import (
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/naming"
)

type (
	ProviderOne struct{}
	ProviderTwo struct{}
)

func (p *ProviderOne) CanHandle(name string) bool {
	domain := naming.GetTopDomain(name, ".")
	return domain == ".one" || domain == ".zero"
}

func (p *ProviderOne) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	return []blockatlas.Resolved{}, nil
}

func (p *ProviderTwo) CanHandle(name string) bool {
	domain := naming.GetTopDomain(name, ".")
	return domain == ".two" || domain == ".zero"
}

func (p *ProviderTwo) Lookup(coins []uint64, name string) ([]blockatlas.Resolved, error) {
	return []blockatlas.Resolved{}, nil
}

func setupProviders() map[uint]blockatlas.NamingServiceAPI {
	return map[uint]blockatlas.NamingServiceAPI{
		1: &ProviderOne{},
		2: &ProviderTwo{},
	}
}

func TestFindHandlerApis(t *testing.T) {
	tests := []struct {
		name      string
		wantCount int
	}{
		{
			name:      "user.one",
			wantCount: 1,
		},
		{
			name:      "user.two",
			wantCount: 1,
		},
		{
			name:      "user.NOSUCHDOMAIN",
			wantCount: 0,
		},
		{
			name:      "user.zero",
			wantCount: 2,
		},
		{
			name:      "user.ONE",
			wantCount: 1,
		},
	}
	allApis := setupProviders()
	for _, tt := range tests {
		res := findHandlerApis(tt.name, allApis)
		if len(res) != tt.wantCount {
			t.Errorf("Wrong answer %v %v %v", tt.name, len(res), tt.wantCount)
		}
	}
}
