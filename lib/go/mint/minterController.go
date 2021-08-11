package mint

import (
	"errors"

	"github.com/bjartek/go-with-the-flow/gwtf"
	util "github.com/flow-usdc/flow-usdc"
)

func CreateMinterController(
	g *gwtf.GoWithTheFlow,
	account string,
) (events []*gwtf.FormatedEvent, err error) {
	txFilename := "../../../transactions/mint/create_new_minter_controller.cdc"
	txScript := util.ParseCadenceTemplate(txFilename)
	e, err := g.TransactionFromFile(txFilename, txScript).
		SignProposeAndPayAs(account).
		AccountArgument(account).
		Run()
	events = util.ParseTestEvents(e)
	return
}

func GetMinterControllerUUID(g *gwtf.GoWithTheFlow, minterControllerAcct string) (uuid uint64, err error) {
	filename := "../../../scripts/get_minter_controller_uuid.cdc"
	script := util.ParseCadenceTemplate(filename)
	r, err := g.ScriptFromFile(filename, script).AccountArgument(minterControllerAcct).RunReturns()
	uuid, ok := r.ToGoValue().(uint64)
	if !ok {
		err = errors.New("returned not uint64")
	}
	return
}

func GetManagedMinter(g *gwtf.GoWithTheFlow, minterController uint64) (uuid uint64, err error) {
	filename := "../../../scripts/get_managed_minter.cdc"
	script := util.ParseCadenceTemplate(filename)
	r, err := g.ScriptFromFile(filename, script).UInt64Argument(minterController).RunReturns()
	uuid, ok := r.ToGoValue().(uint64)
	if !ok {
		err = errors.New("returned nil")
	}
	return
}

func ConfigureMinterAllowance(
	g *gwtf.GoWithTheFlow,
	minterControllerAcct string,
	amount string,
) (events []*gwtf.FormatedEvent, err error) {
	txFilename := "../../../transactions/mint/configure_minter_allowance.cdc"
	txScript := util.ParseCadenceTemplate(txFilename)
	e, err := g.TransactionFromFile(txFilename, txScript).
		SignProposeAndPayAs(minterControllerAcct).
		UFix64Argument(amount).
		Run()
	events = util.ParseTestEvents(e)
	return
}

func RemoveMinter(
	g *gwtf.GoWithTheFlow,
	minterControllerAcct string,
) (events []*gwtf.FormatedEvent, err error) {
	txFilename := "../../../transactions/mint/remove_minter.cdc"
	txScript := util.ParseCadenceTemplate(txFilename)
	e, err := g.TransactionFromFile(txFilename, txScript).
		SignProposeAndPayAs(minterControllerAcct).
		Run()
	events = util.ParseTestEvents(e)
	return
}

func IncreaseOrDecreaseMinterAllowance(
	g *gwtf.GoWithTheFlow,
	minterControllerAcct string,
	absDelta string,
	inc uint,
) (events []*gwtf.FormatedEvent, err error) {
	var txFilename string
	if inc == 1 {
		txFilename = "../../../transactions/mint/increase_minter_allowance.cdc"
	} else {
		txFilename = "../../../transactions/mint/decrease_minter_allowance.cdc"
	}

	txScript := util.ParseCadenceTemplate(txFilename)
	e, err := g.TransactionFromFile(txFilename, txScript).
		SignProposeAndPayAs(minterControllerAcct).
		UFix64Argument(absDelta).
		Run()
	events = util.ParseTestEvents(e)
	return
}