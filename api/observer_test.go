package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/observer"
	"testing"
)

func Test_parseSubscriptions(t *testing.T) {
	tests := []struct {
		name          string
		subscriptions map[string][]string
		webhook       string
		wantSubs      []observer.Subscription
	}{
		{
			name: "webhook with 2 coins",
			subscriptions: map[string][]string{
				"2": {"zpub6rH4MwgyTmuexAX6HAraks5cKv5BbtmwdLirvnU5845ovUJb4abgjt9DtXK4ZEaToRrNj8dQznuLC6Nka4eMviGMinCVMUxKLpuyddcG9Vc"},
				"0": {"xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA", "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL"},
			},
			webhook: "http://127.0.0.1:8080",
			wantSubs: []observer.Subscription{
				{
					Coin: uint(2), Webhooks: []string{"http://127.0.0.1:8080"},
					Address: "zpub6rH4MwgyTmuexAX6HAraks5cKv5BbtmwdLirvnU5845ovUJb4abgjt9DtXK4ZEaToRrNj8dQznuLC6Nka4eMviGMinCVMUxKLpuyddcG9Vc",
				},
				{
					Coin: uint(0), Webhooks: []string{"http://127.0.0.1:8080"},
					Address: "xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA",
				},
				{
					Coin: uint(0), Webhooks: []string{"http://127.0.0.1:8080"},
					Address: "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL",
				},
			},
		}, {
			name: "webhook with 1 coin",
			subscriptions: map[string][]string{
				"0": {"xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA", "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL"},
			},
			webhook: "http://127.0.0.1:8080",
			wantSubs: []observer.Subscription{
				{
					Coin: uint(0), Webhooks: []string{"http://127.0.0.1:8080"},
					Address: "xpub6BpYi6J1GZzfY3yY7DbhLLccF3efQa18nQngM3jaehgtNSoEgk6UtPULpC3oK5oA3trczY8Ld34LFw1USMPfGHwTEizdD5QyGcMyuh2UoBA",
				},
				{
					Coin: uint(0), Webhooks: []string{"http://127.0.0.1:8080"},
					Address: "xpub6CYwPfnPJLPquufPkb98coSb3mdy1CgaZrWUtYWGJTJ4VWZUbzH9HLGy7nHpP7DG4UdTkYYpirkTWQSP7pWHsrk24Nos5oYNHpfr4BgPVTL",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSubs := parseSubscriptions(tt.subscriptions, tt.webhook)
			assert.ElementsMatch(t, tt.wantSubs, gotSubs)
		})
	}
}
