package helpers

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"strings"
)

func B64ToHex(b64Str string) (hexStr string, err error) {
	bts, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		return hexStr, fmt.Errorf("base64.StdEncoding.DecodeString: %s", err.Error())
	}
	return hex.EncodeToString(bts), nil
}

func ValAddressFromBech32(address string) (addr types.ValAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return types.ValAddress{}, fmt.Errorf("empty address string is not allowed")
	}

	bech32PrefixValAddr := "evmosvaloper"

	bz, err := types.GetFromBech32(address, bech32PrefixValAddr)
	if err != nil {
		return nil, err
	}

	err = types.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func Bech32Addr(aa types.AccAddress) string {
	prefix := "evmos"
	bech32Addr, err := bech32.ConvertAndEncode(prefix, aa.Bytes())
	if err != nil {
		panic(err)
	}
	return bech32Addr
}
