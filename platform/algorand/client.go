package algorand

//
//import (
//	"fmt"
//	"github.com/trustwallet/blockatlas"
//	"net/http"
//)
//
//type Client struct {
//	Request blockatlas.Request
//	URL     string
//}
//
//func InitClient(URL string) Client {
//	return Client{
//		Request: blockatlas.Request{
//			HttpClient: http.DefaultClient,
//			ErrorHandler: func(res *http.Response, uri string) error {
//				return nil
//			},
//		},
//		URL: URL,
//	}
//}
//
//func (c *Client) GetLatestBlock() (int64, error) {
//	var status Status
//	err := c.Request.Get(&status, c.URL, "v1/status", nil)
//	if err != nil {
//		return 0, err
//	}
//	return int64(status.LastRound), nil
//}
//
//func (c *Client) GetTxsInBlock(number int64) ([]Transaction, error) {
//	path := fmt.Sprintf("v1/block/%d", number)
//	var resp BlockResponse
//	err := c.Request.Get(&resp, c.URL, path, nil)
//	if err != nil {
//		return nil, err
//	}
//	return resp.Transactions.Transactions, nil
//}
//
//func (c *Client) GetTxsOfAddress(address string) ([]Transaction, error) {
//	var response TransactionsResponse
//	path := fmt.Sprintf("v1/account/%s/transactions", address)
//	err := c.Request.Get(&response, c.URL, path, nil)
//	if err != nil {
//		return nil, blockatlas.ErrSourceConn
//	}
//	return response.Transactions, err
//}
