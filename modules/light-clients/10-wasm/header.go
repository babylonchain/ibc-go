package wasm

import (
	"fmt"

	"github.com/cosmos/ibc-go/v5/modules/core/exported"
)

var _ exported.ClientMessage = &Header{}
var _ exported.LCHeader = &Header{}

func (m Header) ClientType() string {
	return exported.Wasm
}

func (m Header) ValidateBasic() error {
	if m.Data == nil || len(m.Data) == 0 {
		return fmt.Errorf("data cannot be empty")
	}

	return nil
}

func (h Header) HeaderHeight() exported.Height {
	// TODO what exactly is this wasm header height ? Does it represent heigh of the
	// chain inside wasm smart contract ?
	return h.Height
}
