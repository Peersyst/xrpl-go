package transactions

import (
	"fmt"

	"github.com/Peersyst/xrpl-go/pkg/typecheck"
)

// https://xrpl.org/docs/references/protocol/transactions/types/ticketcreate#ticketcreate-fields
const MAX_TICKETS = 250

// A TicketCreate transaction sets aside one or more sequence numbers as Tickets.
type TicketCreate struct {
	// Base transaction fields
	BaseTx

	//How many Tickets to create. This must be a positive number and cannot cause
	// the account to own more than 250 Tickets after executing this transaction.
	TicketCount uint32
}

func (*TicketCreate) TxType() TxType {
	return TicketCreateTx
}

func (t *TicketCreate) Flatten() FlatTransaction {
	flattened := t.BaseTx.Flatten()

	flattened["TransactionType"] = "TicketCreate"

	if t.TicketCount != 0 {
		flattened["TicketCount"] = t.TicketCount
	}

	return flattened
}

func ValidateTicketCreate(tx FlatTransaction) error {
	err := ValidateBaseTransaction(tx)
	if err != nil {
		return err
	}

	err = ValidateRequiredField(tx, "TicketCount", typecheck.IsUint)
	if err != nil {
		return err
	}

	// check if the ticket count is between 1 and MAX_TICKETS
	if tx["TicketCount"].(uint) < 1 || tx["TicketCount"].(uint) > MAX_TICKETS {
		return fmt.Errorf("field 'TicketCount' must be between 1 and %v", MAX_TICKETS)
	}

	return nil
}
