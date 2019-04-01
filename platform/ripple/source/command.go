package source

import "encoding/json"

var cmdId uint64

func nextCommandId() uint64 {
	cmdId++
	return cmdId
}

type Result map[string]interface{}

type CommandResponse struct {
	Id     uint64          `json:"id"`
	Status string          `json:"status"`
	Type   string          `json:"type"`
	Error  string          `json:"error"`
	Result json.RawMessage `json:"result"`
}

type Command interface {
	Id() uint64
	Command() string
}

type baseCommand struct {
	BaseId      uint64 `json:"id"`
	BaseCommand string `json:"command"`
}

func (c baseCommand) Id() uint64 {
	return c.BaseId
}

func (c baseCommand) Command() string {
	return c.BaseCommand
}

func SubscribeCommand(streams ...string) Command {
	type cmd struct {
		baseCommand
		Streams []string `json:"streams"`
	}

	return cmd{
		baseCommand{
			BaseId:      nextCommandId(),
			BaseCommand: "subscribe",
		},
		streams,
	}
}

func LedgerCommand(hash string) Command {
	type cmd struct {
		baseCommand
		LedgerHash   string `json:"ledger_hash"`
		Expand       bool   `json:"expand"`
		Transactions bool   `json:"transactions"`
	}

	return cmd{
		baseCommand{
			BaseId:      nextCommandId(),
			BaseCommand: "ledger",
		},
		hash,
		true,
		true,
	}
}
