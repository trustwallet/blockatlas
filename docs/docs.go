// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-27 22:58:34.193834 -0300 -03 m=+0.046249491

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/{coin}/transactions/{address}": {
            "get": {
                "description": "Get Transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tx"
                ],
                "summary": "Get transactions from address",
                "operationId": "tx_v2",
                "parameters": [
                    {
                        "type": "string",
                        "default": "tezos",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
                        "description": "the address to get transactions",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.TxPage"
                        }
                    }
                }
            }
        },
        "/v1/{coin}/{address}": {
            "get": {
                "description": "Get Transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tx"
                ],
                "summary": "Get transactions from address",
                "operationId": "tx_v1",
                "parameters": [
                    {
                        "type": "string",
                        "default": "tezos",
                        "description": "the coin name",
                        "name": "coin",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
                        "description": "the address to get transactions",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/blockatlas.TxPage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "blockatlas.Tx": {
            "type": "object",
            "properties": {
                "block": {
                    "description": "Height of the block the transaction was included in",
                    "type": "integer"
                },
                "coin": {
                    "description": "SLIP-44 coin index of the platform",
                    "type": "integer"
                },
                "date": {
                    "description": "Unix timestamp of the block the transaction was included in",
                    "type": "integer"
                },
                "direction": {
                    "description": "Transaction Direction",
                    "type": "string"
                },
                "error": {
                    "description": "Empty if the transaction was successful,\nelse error explaining why the transaction failed (optional)",
                    "type": "string"
                },
                "fee": {
                    "description": "Transaction fee (native currency)",
                    "type": "string"
                },
                "from": {
                    "description": "Address of the transaction sender",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier",
                    "type": "string"
                },
                "inputs": {
                    "description": "Input addresses",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/blockatlas.TxOutput"
                    }
                },
                "memo": {
                    "description": "Meta data object",
                    "type": "string"
                },
                "metadata": {
                    "type": "object"
                },
                "outputs": {
                    "description": "Output addresses",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/blockatlas.TxOutput"
                    }
                },
                "sequence": {
                    "description": "Transaction nonce or sequence",
                    "type": "integer"
                },
                "status": {
                    "description": "Status of the transaction",
                    "type": "string"
                },
                "to": {
                    "description": "Address of the transaction recipient",
                    "type": "string"
                },
                "type": {
                    "description": "Type of metadata",
                    "type": "string"
                }
            }
        },
        "blockatlas.TxOutput": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "blockatlas.TxPage": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "block": {
                        "description": "Height of the block the transaction was included in",
                        "type": "integer"
                    },
                    "coin": {
                        "description": "SLIP-44 coin index of the platform",
                        "type": "integer"
                    },
                    "date": {
                        "description": "Unix timestamp of the block the transaction was included in",
                        "type": "integer"
                    },
                    "direction": {
                        "description": "Transaction Direction",
                        "type": "string"
                    },
                    "error": {
                        "description": "Empty if the transaction was successful,\nelse error explaining why the transaction failed (optional)",
                        "type": "string"
                    },
                    "fee": {
                        "description": "Transaction fee (native currency)",
                        "type": "string"
                    },
                    "from": {
                        "description": "Address of the transaction sender",
                        "type": "string"
                    },
                    "id": {
                        "description": "Unique identifier",
                        "type": "string"
                    },
                    "inputs": {
                        "description": "Input addresses",
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/blockatlas.TxOutput"
                        }
                    },
                    "memo": {
                        "description": "Meta data object",
                        "type": "string"
                    },
                    "metadata": {
                        "type": "object"
                    },
                    "outputs": {
                        "description": "Output addresses",
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/blockatlas.TxOutput"
                        }
                    },
                    "sequence": {
                        "description": "Transaction nonce or sequence",
                        "type": "integer"
                    },
                    "status": {
                        "description": "Status of the transaction",
                        "type": "string"
                    },
                    "to": {
                        "description": "Address of the transaction recipient",
                        "type": "string"
                    },
                    "type": {
                        "description": "Type of metadata",
                        "type": "string"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
