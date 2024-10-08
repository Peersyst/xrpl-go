package websocket

import (
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/account"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/common"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/ledger"
	"github.com/Peersyst/xrpl-go/xrpl/model/requests/server"
	"github.com/Peersyst/xrpl-go/xrpl/model/transactions/types"
	"github.com/Peersyst/xrpl-go/xrpl/utils"
)

// GetAccountInfo retrieves information about an account on the XRP Ledger.
// It takes an AccountInfoRequest as input and returns an AccountInfoResponse,
// along with the raw XRPL response and any error encountered.
func (c *WebsocketClient) GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, error) {
	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	var air account.AccountInfoResponse
	err = res.GetResult(&air)
	if err != nil {
		return nil, err
	}
	return &air, nil
}

// GetAccountObjects retrieves a list of objects owned by an account on the XRP Ledger.
// It takes an AccountObjectsRequest as input and returns an AccountObjectsResponse,
// along with any error encountered.
func (c *WebsocketClient) GetAccountObjects(req *account.AccountObjectsRequest) (*account.AccountObjectsResponse, error) {
	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	var acr account.AccountObjectsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// GetXrpBalance retrieves the XRP balance of a given account address.
// It returns the balance as a string in XRP (not drops) and any error encountered.
func (c *WebsocketClient) GetXrpBalance(address string) (string, error) {
	res, err := c.GetAccountInfo(&account.AccountInfoRequest{
		Account: types.Address(address),
	})
	if err != nil {
		return "", err
	}

	xrpBalance, err := utils.DropsToXrp(res.AccountData.Balance.String())
	if err != nil {
		return "", err
	}

	return xrpBalance, nil
}

func (c *WebsocketClient) GetAccountLines(req *account.AccountLinesRequest) (*account.AccountLinesResponse, error) {
	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	var acr account.AccountLinesResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, err
	}
	return &acr, nil
}

// Returns the index of the most recently validated ledger.
func (c *WebsocketClient) GetLedgerIndex() (*common.LedgerIndex, error) {
	res, err := c.sendRequest(&ledger.LedgerRequest{
		LedgerIndex: common.LedgerTitle("validated"),
	})
	if err != nil {
		return nil, err
	}

	var lr ledger.LedgerResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, err
	}
	return &lr.LedgerIndex, err
}

// GetServerInfo retrieves information about the server.
// It takes a ServerInfoRequest as input and returns a ServerInfoResponse,
// along with any error encountered.
func (c *WebsocketClient) GetServerInfo(req *server.ServerInfoRequest) (*server.ServerInfoResponse, error) {
	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	var sir server.ServerInfoResponse
	err = res.GetResult(&sir)
	if err != nil {
		return nil, err
	}
	return &sir, err
}
