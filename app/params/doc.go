/*
Package params defines the simulation parameters in the gaia.

It contains the default weights used for each transaction used on the module's
simulation. These weights define the chance for a transaction to be simulated at
any gived operation.

You can repace the default values for the weights by providing a params.json
file with the weights defined for each of the transaction operations:

	{
		"op_weight_msg_send": 60,
		"op_weight_msg_delegate": 100,
	}

In the example above, the `MsgSend` has 60% chance to be simulated, while the
`MsgDelegate` will always be simulated.
*/
package params

// Default simulation operation weights for messages and gov proposals
const (
	DefaultWeightMsgStoreCode           int = 50
	DefaultWeightMsgInstantiateContract int = 100
	DefaultWeightMsgExecuteContract     int = 100
	DefaultWeightMsgUpdateAdmin         int = 25
	DefaultWeightMsgClearAdmin          int = 10
	DefaultWeightMsgMigrateContract     int = 50

	DefaultWeightStoreCodeProposal                   int = 5
	DefaultWeightInstantiateContractProposal         int = 5
	DefaultWeightUpdateAdminProposal                 int = 5
	DefaultWeightExecuteContractProposal             int = 5
	DefaultWeightClearAdminProposal                  int = 5
	DefaultWeightMigrateContractProposal             int = 5
	DefaultWeightSudoContractProposal                int = 5
	DefaultWeightPinCodesProposal                    int = 5
	DefaultWeightUnpinCodesProposal                  int = 5
	DefaultWeightUpdateInstantiateConfigProposal     int = 5
	DefaultWeightStoreAndInstantiateContractProposal int = 5
)
