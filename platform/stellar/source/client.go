package source

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/xdr"
	"sync"
	"time"
)

type Client struct {
	horizon.Client
}

type Tx struct {
	Tx       horizon.Transaction
	Envelope xdr.TransactionEnvelope
	Payment  *xdr.PaymentOp
}

type txStream struct {
	sync.Mutex
	Txs []Tx
	Err error
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	var ctxt context.Context
	var cancel context.CancelFunc
	ctxt = context.Background()
	ctxt, _ = context.WithTimeout(context.Background(), time.Second)
	ctxt, cancel = context.WithCancel(ctxt)

	var stream txStream
	go c.streamTransactions(ctxt, cancel, address, &stream)

	// Wait for transaction stream to finish
	<-ctxt.Done()

	stream.Lock()
	defer stream.Unlock()
	if stream.Err != nil {
		return nil, stream.Err
	}

	return stream.Txs, nil
}

func (c *Client) streamTransactions(ctxt context.Context, cancel context.CancelFunc, address string, stream *txStream) {
	defer cancel()

	err := c.StreamTransactions(ctxt, address, nil, func(tx horizon.Transaction) {
		if tx.ResultXdr == "" {
			return
		}
		if !tx.Successful {
			return
		}

		var envelope xdr.TransactionEnvelope
		err := xdr.SafeUnmarshalBase64(tx.EnvelopeXdr, &envelope)
		if err != nil {
			return
		}

		stream.Lock()
		defer stream.Unlock()

		for _, op := range envelope.Tx.Operations {
			payment := op.Body.PaymentOp
			if payment == nil {
				continue
			}
			spew.Dump(payment)
			if payment.Asset.Type != xdr.AssetTypeAssetTypeNative {
				continue
			}
			stream.Txs = append(stream.Txs, Tx{
				Tx:       tx,
				Envelope: envelope,
				Payment:  payment,
			})
			if len(stream.Txs) >= 25 {
				return
			}
		}
	})
	if err != nil {
		stream.Lock()
		stream.Err = err
		stream.Unlock()
	}
}
