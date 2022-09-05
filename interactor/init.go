package interactor

import "sso/contract"

type Interactor struct {
	store contract.Repository
}

func NewDB(store contract.Repository) Interactor {

	return Interactor{store: store}
}
