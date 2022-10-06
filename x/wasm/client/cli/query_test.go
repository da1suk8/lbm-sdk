package cli

import (
	"context"
	"encoding/hex"
	"os"
	"testing"

	"github.com/line/lbm-sdk/client"
	"github.com/line/lbm-sdk/codec"
	"github.com/line/lbm-sdk/x/wasm/lbmtypes"
	"github.com/line/lbm-sdk/x/wasm/types"
	ocabcitypes "github.com/line/ostracon/abci/types"
	ocrpcmocks "github.com/line/ostracon/rpc/client/mocks"
	ocrpctypes "github.com/line/ostracon/rpc/core/types"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	codeID       = "1"
	accAddress   = "link1yxfu3fldlgux939t0gwaqs82l4x77v2kasa7jf"
	queryJson    = `{"a":"b"}`
	queryJsonHex = hex.EncodeToString([]byte(queryJson))
)

func TestGetQueryCmd(t *testing.T) {
	tests := []struct {
		name string
		want *cobra.Command
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetQueryCmd()
			assert.NotNilf(t, cmd, "GetQueryCmd()")
		})
	}
}

func TestGetCmdLibVersion(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdLibVersion()
			assert.Equalf(t, tt.want, cmd.RunE(cmd, nil), "GetCmdLibVersion()")
		})
	}
}

func TestGetCmdListCode(t *testing.T) {
	res := types.QueryCodesResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdListCode()
			cmd.SetContext(ctx)
			assert.Equalf(t, tt.want, cmd.RunE(cmd, nil), "GetCmdListCode()")
		})
	}
}

func TestGetCmdListContractByCode(t *testing.T) {
	res := types.QueryContractsByCodeResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdListContractByCode()
			cmd.SetContext(ctx)
			args := []string{codeID}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdListContractByCode()")
		})
	}
}

func TestGetCmdQueryCode(t *testing.T) {
	res := types.QueryCodeResponse{Data: []byte{0}}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdQueryCode()
			cmd.SetContext(ctx)
			args := []string{codeID}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdQueryCode()")
			downloaded := "contract-" + codeID + ".wasm"
			assert.FileExists(t, downloaded)
			assert.NoError(t, os.Remove(downloaded))
		})
	}
}

func TestGetCmdQueryCodeInfo(t *testing.T) {
	res := types.QueryCodeResponse{CodeInfoResponse: &types.CodeInfoResponse{}}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdQueryCodeInfo()
			cmd.SetContext(ctx)
			args := []string{codeID}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdQueryCodeInfo()")
		})
	}
}

func TestGetCmdGetContractInfo(t *testing.T) {
	res := types.QueryContractInfoResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdGetContractInfo()
			cmd.SetContext(ctx)
			args := []string{accAddress}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdGetContractInfo()")
		})
	}
}

func TestGetCmdGetContractState(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdGetContractState()
			assert.Equalf(t, tt.want, cmd.RunE(cmd, nil), "GetCmdGetContractState()")
		})
	}
}

func TestGetCmdGetContractStateAll(t *testing.T) {
	res := types.QueryAllContractStateResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdGetContractStateAll()
			cmd.SetContext(ctx)
			args := []string{accAddress}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdGetContractStateAll()")
		})
	}
}

func TestGetCmdGetContractStateRaw(t *testing.T) {
	res := types.QueryRawContractStateResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdGetContractStateRaw()
			cmd.SetContext(ctx)
			args := []string{accAddress, queryJsonHex}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdGetContractStateRaw()")
		})
	}
}

func TestGetCmdGetContractStateSmart(t *testing.T) {
	res := types.QueryRawContractStateResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdGetContractStateSmart()
			cmd.SetContext(ctx)
			args := []string{accAddress, queryJson}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdGetContractStateSmart()")
		})
	}
}

func TestGetCmdGetContractHistory(t *testing.T) {
	res := types.QueryContractHistoryResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdGetContractHistory()
			cmd.SetContext(ctx)
			args := []string{accAddress}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdGetContractHistory()")
		})
	}
}

func TestGetCmdListPinnedCode(t *testing.T) {
	res := types.QueryPinnedCodesResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdListPinnedCode()
			cmd.SetContext(ctx)
			assert.Equalf(t, tt.want, cmd.RunE(cmd, nil), "GetCmdListPinnedCode()")
		})
	}
}

func TestGetCmdListInactiveContracts(t *testing.T) {
	res := lbmtypes.QueryInactiveContractsResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdListInactiveContracts()
			cmd.SetContext(ctx)
			assert.Equalf(t, tt.want, cmd.RunE(cmd, nil), "GetCmdListInactiveContracts()")
		})
	}
}

func TestGetCmdIsInactiveContract(t *testing.T) {
	res := lbmtypes.QueryInactiveContractResponse{}
	bz, err := res.Marshal()
	require.NoError(t, err)
	ctx := makeContext(bz)
	tests := []struct {
		name string
		want error
	}{
		{"execute success", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := GetCmdIsInactiveContract()
			cmd.SetContext(ctx)
			args := []string{accAddress}
			assert.Equalf(t, tt.want, cmd.RunE(cmd, args), "GetCmdIsInactiveContract()")
		})
	}
}
func makeContext(bz []byte) context.Context {
	result := ocrpctypes.ResultABCIQuery{Response: ocabcitypes.ResponseQuery{Value: bz}}
	mockClient := ocrpcmocks.RemoteClient{}
	mockClient.On("ABCIQueryWithOptions",
		mock.Anything, mock.Anything, mock.Anything, mock.Anything,
	).Return(&result, nil)
	cli := client.Context{}.WithClient(&mockClient).WithCodec(codec.NewProtoCodec(nil))
	return context.WithValue(context.Background(), client.ClientContextKey, &cli)
}
