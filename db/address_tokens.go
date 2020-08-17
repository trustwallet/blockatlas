package db

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/trustwallet/blockatlas/db/models"
)

// GetAddressTokens returns the set of known tokens for an address.
// If the set is unknown, "ok" is set to false.
func (i *Instance) GetAddressTokens(coin uint, address string) (tokens []string, ok bool, err error) {
	// Check if the address is tracked (i.e. was MarkAddressTokensTracked() called?).
	ok, err = i.AddressTokensTracked(coin, address)
	if err != nil || !ok {
		return
	}
	// The address is tracked. Read the list of tokens.
	if err := i.Gorm.Table("address_tokens").
		Where("coin = ? AND address = ?", coin, address).
		Pluck("token", &tokens).
		Error; err != nil {
		return nil, false, err
	}
	return
}

// UpsertAddressTokens adds in a new set of tokens for an address.
// Future calls to GetTokens() will return the set union of the previous set and the new set,
// if MarkAddressTokensTracked() was called at a previous time.
// This func is idempotent.
func (i *Instance) UpsertAddressTokens(coin uint, address string, tokens []string) error {
	if len(tokens) == 0 {
		return nil
	}
	// Insert token mapping
	stmt, values := upsertTokensStmt(coin, address, tokens)
	return i.Gorm.Exec(stmt, values...).Error
}

// UpsertAddressTokensMulti is like UpsertAddressTokens() but for multiple mappings at once.
// This func is idempotent.
func (i *Instance) UpsertAddressTokensMulti(mappings []models.AddressToken) error {
	if len(mappings) == 0 {
		return nil
	}
	// Insert token mapping
	stmt, values := upsertTokensMultiStmt(mappings)
	return i.Gorm.Exec(stmt, values...).Error
}

// AddressTokensTracked returns whether the set of tokens for a single address are marked "tracked".
// In other words, was MarkAddressTokensTracked() called previously?
func (i *Instance) AddressTokensTracked(coin uint, address string) (bool, error) {
	trackedEntry := &models.AddressTokenTracker{Coin: coin, Address: address}
	if err := i.Gorm.Take(trackedEntry).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // The address is not fully tracked for tokens.
	} else if err != nil {
		return false, err // Error during lookup
	}
	return true, nil
}

// AddressTokensTrackedMulti checks AddressTokensTracked() for multiple addresses at once.
// It returns the list of address for which AddressTokensTracked() resolves to (true, nil).
func (i *Instance) AddressTokensTrackedMulti(coin uint, addresses []string) (matched []string, err error) {
	err = i.Gorm.Where("coin = ? AND address IN (?)", coin, addresses).
		Pluck("address", &matched).
		Error
	return
}

// MarkAddressTokensTracked marks the set of tokens as "tracked" for a single address.
// This will activate GetAddressTokens().
// This func is idempotent.
func (i *Instance) MarkAddressTokensTracked(coin uint, address string) error {
	trackedEntry := &models.AddressTokenTracker{Coin: coin, Address: address}
	return i.Gorm.FirstOrCreate(trackedEntry).Error
}

func upsertTokensStmt(coin uint, address string, tokens []string) (stmt string, values []interface{}) {
	mappings := make([]models.AddressToken, len(tokens))
	for i, token := range tokens {
		mappings[i] = models.AddressToken{
			Coin:    coin,
			Address: address,
			Token:   token,
		}
	}
	return upsertTokensMultiStmt(mappings)
}

func upsertTokensMultiStmt(mappings []models.AddressToken) (stmt string, values []interface{}) {
	var batch strings.Builder
	batch.WriteString("INSERT INTO address_tokens (coin, address, token) VALUES\n")
	for i, mapping := range mappings {
		if i > 0 {
			batch.WriteString(",\n")
		}
		batch.WriteString("  (?, ?, ?)")
		values = append(values, mapping.Coin, mapping.Address, mapping.Token)
	}
	batch.WriteString("\nON CONFLICT DO NOTHING;")
	stmt = batch.String()
	return
}
