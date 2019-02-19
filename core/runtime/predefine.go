package runtime

type PredefinedFunction struct {
	Name   string
	Action func(vm *VM, frame *Frame, params ...interface{}) interface{}
}

var (
	//TODO: add 'Action' field
	PredefinedFuncs = []PredefinedFunction{
		PredefinedFunction{Name: "contractCall"},
		PredefinedFunction{Name: "blockHash"},
		PredefinedFunction{Name: "blockCoinbase"},
		PredefinedFunction{Name: "blockGaslimit"},
		PredefinedFunction{Name: "blockNumber"},
		PredefinedFunction{Name: "blockTimestamp"},
		PredefinedFunction{Name: "gasleft"},
		PredefinedFunction{Name: "msgData"},
		PredefinedFunction{Name: "msgGas"},
		PredefinedFunction{Name: "msgSender"},
		PredefinedFunction{Name: "msgSig"},
		PredefinedFunction{Name: "msgValue"},
		PredefinedFunction{Name: "txGasprice"},
		PredefinedFunction{Name: "txOrigin"},
		PredefinedFunction{Name: "addressBalance"},
		PredefinedFunction{Name: "addressTransfer"},
		PredefinedFunction{Name: "addressSend"},
		PredefinedFunction{Name: "addressCall"},
		PredefinedFunction{Name: "addressCallcode"},
		PredefinedFunction{Name: "addressDelegateCall"},
	}
)
