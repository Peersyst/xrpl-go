package websocket

import (
	"github.com/Peersyst/xrpl-go/xrpl/currency"
	"github.com/Peersyst/xrpl-go/xrpl/queries/account"
	"github.com/Peersyst/xrpl-go/xrpl/queries/common"
	"github.com/Peersyst/xrpl-go/xrpl/queries/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/queries/server"
	"github.com/Peersyst/xrpl-go/xrpl/transaction/types"
)

// GetAccountInfo retrieves information about an account on the XRP Ledger.
// It takes an AccountInfoRequest as input and returns an AccountInfoResponse,
// along with the raw XRPL response and any error encountered.
func (c *Client) GetAccountInfo(req *account.InfoRequest) (*account.InfoResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var air account.InfoResponse
	err = res.GetResult(&air)
	if err != nil {
		return nil, err
	}
	return &air, nil
}

// GetAccountObjects retrieves a list of objects owned by an account on the XRP Ledger.
// It takes an AccountObjectsRequest as input and returns an AccountObjectsResponse,
// along with any error encountered.
func (c *Client) GetAccountObjects(req *account.ObjectsRequest) (*account.ObjectsResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.ObjectsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// GetXrpBalance retrieves the XRP balance of a given account address.
// It returns the balance as a string in XRP (not drops) and any error encountered.
func (c *Client) GetXrpBalance(address types.Address) (string, error) {
	res, err := c.GetAccountInfo(&account.InfoRequest{
		Account: address,
	})
	if err != nil {
		return "", err
	}

	xrpBalance, err := currency.DropsToXrp(res.AccountData.Balance.String())
	if err != nil {
		return "", err
	}

	return xrpBalance, nil
}

// GetAccountLines retrieves the lines associated with an account on the XRP Ledger.
// It takes an AccountLinesRequest as input and returns an AccountLinesResponse,
// along with any error encountered.
func (c *Client) GetAccountLines(req *account.LinesRequest) (*account.LinesResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var acr account.LinesResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// Returns the index of the most recently validated ledger.
func (c *Client) GetLedgerIndex() (common.LedgerIndex, error) {
	res, err := c.Request(&ledger.Request{
		LedgerIndex: common.LedgerTitle("validated"),
	})
	if err != nil {
		return 0, err
	}

	var lr ledger.Response
	err = res.GetResult(&lr)
	if err != nil {
		return 0, err
	}
	return lr.LedgerIndex, err
}

// GetServerInfo retrieves information about the server.
// It takes a ServerInfoRequest as input and returns a ServerInfoResponse,
// along with any error encountered.
func (c *Client) GetServerInfo(req *server.InfoRequest) (*server.InfoResponse, error) {
	res, err := c.Request(req)
	if err != nil {
		return nil, err
	}
	var sir server.InfoResponse
	err = res.GetResult(&sir)
	if err != nil {
		return nil, err
	}
	return &sir, err
}
