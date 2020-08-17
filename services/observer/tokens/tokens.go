// Package tokens populates the address <=> token mappings cache.
//
// TODO Is this package in the right spot?
//
// It implements a lazy, incremental cache. The cache key is the (coin, address) tuple,
// and the cache entry is the set of tokens the address has interacted with.
//
// The cache entry is created when the tokens set is queried from an authoritative source (hence "lazy").
// This happens in the FullTokensUpdate() method. From this point on, the addresses' tokens are marked as **tracked**.
//
// A background routine will eavesdrop all token transactions happening on the network.
// It populates the cache incrementally, i.e. adds tokens to the cache entry without re-doing the tokens list every time.
// It does this using the IncrementalTokensUpdate() function.
package tokens

import (
	"fmt"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Service struct {
	DB   *db.Instance
	Coin uint
}

// FullTokensUpdate updates the cache forcefully for a single address.
// It also marks the address as tracked from now on.
// This would be called after fetching the set of tokens from an expensive query.
func (s *Service) FullTokensUpdate(address string, tokens []string) error {
	if err := s.DB.UpsertAddressTokens(s.Coin, address, tokens); err != nil {
		return fmt.Errorf("failed to insert initial address <=> tokens mappings: %w", err)
	}
	if err := s.DB.MarkAddressTokensTracked(s.Coin, address); err != nil {
		return fmt.Errorf("failed to mark address <=> tokens as tracked: %w", err)
	}
	return nil
}

// IncrementalTokensUpdate updates the cache incrementally from new transactions.
// It ignores addresses that are not marked as tracked.
// This would be called in the background when new txs come in.
func (s *Service) IncrementalTokensUpdate(txs blockatlas.Txs) error {
	// Collect list of addresses with token txs.
	activeAddrsMap := make(map[string]bool)
	for _, tx := range txs {
		if tx.Type != blockatlas.TxTokenTransfer {
			// TODO Also respect NativeTokenTransfers?
			continue
		}
		activeAddrsMap[tx.From] = true
		activeAddrsMap[tx.To] = true
	}
	activeAddrs := make([]string, 0, len(activeAddrsMap))
	for addr := range activeAddrsMap {
		activeAddrs = append(activeAddrs, addr)
	}
	// Filter by tracked addresses.
	matched, err := s.DB.AddressTokensTrackedMulti(s.Coin, activeAddrs)
	if err != nil {
		return fmt.Errorf("failed to get tracked addrs: %w", err)
	}
	matchedMap := make(map[string]bool)
	for _, match := range matched {
		matchedMap[match] = true
	}
	// Send out updates.
	var mappings []models.AddressToken
	for _, tx := range txs {
		if tx.Type != blockatlas.TxTokenTransfer {
			// TODO Also respect NativeTokenTransfers?
			continue
		}
		if matchedMap[tx.From] {
			mappings = append(mappings, models.AddressToken{
				Coin: s.Coin, Address: tx.From, Token: "", // FIXME how to get token?
			})
		}
		if matchedMap[tx.To] {
			mappings = append(mappings, models.AddressToken{
				Coin: s.Coin, Address: tx.To, Token: "", // FIXME how to get token?
			})
		}
	}
	if err := s.DB.UpsertAddressTokensMulti(mappings); err != nil {
		return fmt.Errorf("failed to upsert address <=> token mappings: %w", err)
	}
	return nil
}
