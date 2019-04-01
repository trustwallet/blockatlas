package source

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"net/http"
	"net/url"
)

type Client struct {
	HttpClient *http.Client
	Dialer     *websocket.Dialer
	RpcUrl     string
	WsUrl      string
	commands   map[uint64]Command
}

func (c *Client) GetTxsOfAddress(address string) ([]Transaction, error) {
	uri := fmt.Sprintf("%s/accounts/%s/transactions?type=Payment&result=tesSUCCESS&limit=%d",
		c.RpcUrl,
		url.PathEscape(address),
		models.TxPerPage)
	httpRes, err := c.HttpClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Ripple: Failed to get transactions")
		return nil, ErrSourceConn
	}

	var res Response
	err = json.NewDecoder(httpRes.Body).Decode(&res)

	if res.Result != "success" {
		return nil, ErrSourceConn
	}

	return res.Transactions, nil
}

func (c *Client) initCommands() {
	if c.commands == nil {
		c.commands = make(map[uint64]Command)
	}
}

func (c *Client) pushCommand(cmd Command) {
	c.initCommands()
	c.commands[cmd.Id()] = cmd
}

func (c *Client) popCommand(id uint64) (Command, bool) {
	c.initCommands()
	cmd, exists := c.commands[id]
	if exists {
		delete(c.commands, id)
	}
	return cmd, exists
}

func (c *Client) SubscribeLedger(cLedger chan Ledger, cError chan error) error {
	conn, _, err := c.Dialer.Dial(c.WsUrl, nil)
	if err != nil {
		return err
	}

	cmd := SubscribeCommand("ledger")
	c.pushCommand(cmd)

	go c.readMessages(conn, cLedger, cError)

	if err := conn.WriteJSON(cmd); err != nil {
		cError <- err
	}

	return nil
}

func (c *Client) readMessages(conn *websocket.Conn, cLegder chan Ledger, cError chan error) {
	MessageLoop:
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				logrus.WithError(err).Error("Failed to read message from stream")
				return
			}

			var resp CommandResponse
			if err := json.Unmarshal(msg, &resp); err != nil {
				cError <- err
				continue
			}

			switch resp.Type {
			case "response":
				cmd, hasCmd := c.popCommand(resp.Id)
				if !hasCmd {
					cError <- errors.New(fmt.Sprintf("Command id %d not found", resp.Id))
					continue MessageLoop
				}

				switch cmd.Command() {
				case "subscribe":
					if resp.Status == "error" {
						cError <- errors.New(fmt.Sprintf("Failed to subscribe: %s", resp.Error))
						continue MessageLoop
					}
				case "ledger":
					if resp.Status == "error" {
						cError <- errors.New(fmt.Sprintf("Failed to get ledger: %s", resp.Error))
						continue MessageLoop
					}

					var result struct {
						Ledger Ledger `json:"ledger"`
					}

					if err := json.Unmarshal(resp.Result, &result); err != nil {
						cError <- err
						continue MessageLoop
					}

					cLegder <- result.Ledger
				}
			case "ledgerClosed":
				var header struct{
					Hash    string `json:"ledger_hash"`
					Index   uint64 `json:"ledger_index"`
					TxCount uint   `json:"txn_count"`
				}

				if err := json.Unmarshal(msg, &header); err != nil {
					cError <- err
					continue MessageLoop
				}

				cmd := LedgerCommand(header.Hash)
				if err := conn.WriteJSON(cmd); err != nil {
					cError <- err
					continue MessageLoop
				}
				c.pushCommand(cmd)
			}
		}
}
