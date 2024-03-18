package pathtransfer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/pathtransfer"
)

func TestCallFunc(t *testing.T) {
	transferLine := `
func.vocabulary.SetLimit.input.index@int:Dictionary.pagination.index
func.vocabulary.SetLimit.input.size@int:Dictionary.pagination.size
func.vocabulary.SetLimit.output.offset@int:Dictionary.limit.offset
func.vocabulary.SetLimit.output.size@int:Dictionary.limit.size
	`
	transfer := pathtransfer.Parse(transferLine)
	script, err := transfer.GetCallFnScript("go")
	require.NoError(t, err)
	fmt.Println(script)

}

func TestExplainFuncPath(t *testing.T) {
	t.Run("bas arg", func(t *testing.T) {
		funcPath := "func.vocabulary.SetLimit.input.index@int"
		arg, err := pathtransfer.ExplainFuncPath(funcPath)
		require.NoError(t, err)
		fmt.Println(arg)
	})
	t.Run("object arg", func(t *testing.T) {
		funcPath := "func.vocabulary.SetLimit.input.pagination.index@int"
		arg, err := pathtransfer.ExplainFuncPath(funcPath)
		require.NoError(t, err)
		fmt.Println(arg)
	})

}

func TestGetTransferFuncname(t *testing.T) {

	transfers := make(pathtransfer.Transfers, 0)

	transfers = append(transfers, pathtransfer.Transfer{
		Src: pathtransfer.TransferUnit{Path: "func.vocabulary.TrimName.input.name", Type: "string"},
		Dst: pathtransfer.TransferUnit{Path: "data.userName"},
	})

	transfers = append(transfers, pathtransfer.Transfer{
		Src: pathtransfer.TransferUnit{Path: "func.vocabulary.TrimName.output.name", Type: "string"},
		Dst: pathtransfer.TransferUnit{Path: "data.userName"},
	})

	transfers = append(transfers, pathtransfer.Transfer{
		Src: pathtransfer.TransferUnit{Path: "user.name", Type: "string"},
		Dst: pathtransfer.TransferUnit{Path: "data.userName"},
	})

	transfers = append(transfers, pathtransfer.Transfer{
		Src: pathtransfer.TransferUnit{Path: "func.vocabulary.SetLimit.input.index", Type: "int"},
		Dst: pathtransfer.TransferUnit{Path: "data.pagination.index"},
	})

	transfers = append(transfers, pathtransfer.Transfer{
		Src: pathtransfer.TransferUnit{Path: "func.vocabulary.SetLimit.input.size", Type: "int"},
		Dst: pathtransfer.TransferUnit{Path: "data.pagination.size"},
	})

	transfers = append(transfers, pathtransfer.Transfer{
		Src: pathtransfer.TransferUnit{Path: "func.vocabulary.SetLimit.output.offset", Type: "int"},
		Dst: pathtransfer.TransferUnit{Path: "data.limit.offset"},
	})

	transfers = append(transfers, pathtransfer.Transfer{
		Src: pathtransfer.TransferUnit{Path: "func.vocabulary.SetLimit.output.size", Type: "int"},
		Dst: pathtransfer.TransferUnit{Path: "data.limit.size"},
	})

	var data = `{"data":{"pagination":{"index":1,"size":20}},"user":{"name":"testName"}}`

	dstKey := []string{
		"data.limit.offset",
		"data.limit.size",
	}
	funcName := pathtransfer.GetTransferFuncname(transfers, data, dstKey)
	require.Equal(t, "func.vocabulary.SetLimit", funcName)

}
