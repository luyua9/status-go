package appgeneral

import (
	"context"

	accscommon "github.com/status-im/status-go/internal/accounts-management/common"
	"github.com/status-im/status-go/pkg/version"
)

type API struct {
	s *Service
}

func NewAPI(s *Service) *API {
	return &API{s: s}
}

// Returns a list of currencies for user's selection
func (api *API) GetCurrencies(context context.Context) []*Currency {
	return GetCurrencies()
}

func (api *API) Version(context context.Context) string {
	return version.Version()
}

// GetRandomMnemonic creates a random BIP-39 mnemonic with the default length.
func (api *API) GetRandomMnemonic(context context.Context) (string, error) {
	return accscommon.CreateRandomMnemonicWithDefaultLength()
}
