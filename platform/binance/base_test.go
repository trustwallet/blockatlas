package binance

import (
	"fmt"
	"net/http"
)

const (
	wantedBlock              = `{"number":104867508,"txs":[{"id":"4CD5BAA433BABA63D862141A4A2F9235B0BA5CBAB8114C93A0556ECA4EC7A68A","coin":714,"from":"bnb1l83kstts7lt9dpgawzechnrgjq54dql36dyspc","to":"","fee":"0","date":1596472337,"block":104867508,"status":"completed","sequence":1023322,"type":"any_action","direction":"outgoing","memo":"","metadata":{"coin":714,"title":"Cancel Order","key":"cancel_order","token_id":"AVA-645","name":"SELL","symbol":"AVA","decimals":8,"value":"40"}},{"id":"9B87D17581F2AC73D2999EDE56535E50D9D4DB75150A92A90122190F77D47755","coin":714,"from":"bnb1c4czpzvn0ttdcpnv3cy2858l2m9frxdfgk4jr0","to":"bnb1g2ukzn702napq3levm54m2z3p2gam7upern9aq","fee":"37500","date":1596472337,"block":104867508,"status":"completed","sequence":6,"type":"transfer","memo":"","metadata":{"value":"24481570","symbol":"BNB","decimals":8}},{"id":"5C0580AC983C1CF36F1656D9E8B062CD0578839BEFC839EFD6720FF315B45EEB","coin":714,"from":"bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg","to":"","fee":"0","date":1596472337,"block":104867508,"status":"completed","sequence":1509154,"type":"any_action","direction":"outgoing","memo":"","metadata":{"coin":714,"title":"Place Order","key":"place_order","token_id":"AVA-645","name":"BUY","symbol":"AVA","decimals":8,"value":"105"}}]}`
	wantedBlockMulti         = `{"number":105529271,"txs":[{"id":"432FF828B1DC1C4DAF13A51B6AE7ADD9932B7FC6526C28F7FE9C905F95472820","coin":714,"from":"bnb1z35wusfv8twfele77vddclka9z84ugywug48gn","to":"","fee":"0","date":1596746326,"block":105529271,"status":"completed","sequence":11317170,"type":"any_action","direction":"outgoing","memo":"","metadata":{"coin":714,"title":"Cancel Order","key":"cancel_order","token_id":"","name":"","symbol":"","decimals":8,"value":""}},{"id":"CC7C3EF1407373FDA74B005E64683AB5865126DE93A5FAF755FF5CC948992067","coin":714,"from":"bnb15qced76xere38hmmpe644u5kd8v4lzl9gsex9w","to":"bnb15qced76xere38hmmpe644u5kd8v4lzl9gsex9w","fee":"60000","date":1596746326,"block":105529271,"status":"completed","sequence":2300,"type":"transfer","memo":"0","metadata":{"value":"1","symbol":"BNB","decimals":8}},{"id":"CC7C3EF1407373FDA74B005E64683AB5865126DE93A5FAF755FF5CC948992067","coin":714,"from":"bnb1t38ccns9var4ac4yj2ylmu99r9ecmggr8ye5e5","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"0","date":1596746326,"block":105529271,"status":"completed","sequence":2300,"type":"transfer","memo":"0","metadata":{"value":"39421249","symbol":"BNB","decimals":8}}]}`
	wantedBlockResponseMulti = `{"blockHeight":105529271,"tx":[{"txHash":"432FF828B1DC1C4DAF13A51B6AE7ADD9932B7FC6526C28F7FE9C905F95472820","blockHeight":105529271,"txType":"CANCEL_ORDER","timeStamp":"2020-08-06T20:38:46.583Z","fromAddr":"bnb1z35wusfv8twfele77vddclka9z84ugywug48gn","toAddr":null,"value":null,"txAsset":null,"txFee":null,"code":0,"data":"{\"orderData\":{\"orderId\":\"1468EE412C3ADC9CFF3EF31ADC7EDD288F5E208E-11315084\"}}","memo":"","source":0,"sequence":11317170},{"txHash":"CC7C3EF1407373FDA74B005E64683AB5865126DE93A5FAF755FF5CC948992067","blockHeight":105529271,"txType":"TRANSFER","timeStamp":"2020-08-06T20:38:46.583Z","fromAddr":null,"toAddr":null,"value":null,"txAsset":null,"txFee":null,"code":0,"data":null,"memo":"0","source":1,"sequence":2300,"subTransactions":[{"txHash":"CC7C3EF1407373FDA74B005E64683AB5865126DE93A5FAF755FF5CC948992067","blockHeight":105529271,"txType":"TRANSFER","fromAddr":"bnb15qced76xere38hmmpe644u5kd8v4lzl9gsex9w","toAddr":"bnb15qced76xere38hmmpe644u5kd8v4lzl9gsex9w","txAsset":"BNB","txFee":"0.00060000","value":"0.00000001"},{"txHash":"CC7C3EF1407373FDA74B005E64683AB5865126DE93A5FAF755FF5CC948992067","blockHeight":105529271,"txType":"TRANSFER","fromAddr":"bnb1t38ccns9var4ac4yj2ylmu99r9ecmggr8ye5e5","toAddr":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","txAsset":"BNB","txFee":null,"value":"0.39421249"}]}]}`
	mockedBlockResponse      = `{"blockHeight":104867508,"tx":[{"txHash":"4CD5BAA433BABA63D862141A4A2F9235B0BA5CBAB8114C93A0556ECA4EC7A68A","blockHeight":104867508,"txType":"CANCEL_ORDER","timeStamp":"2020-08-03T16:32:17.963Z","fromAddr":"bnb1l83kstts7lt9dpgawzechnrgjq54dql36dyspc","toAddr":null,"value":null,"txAsset":null,"txFee":null,"code":0,"data":"{\"orderData\":{\"symbol\":\"AVA-645_BUSD-BD1\",\"orderType\":\"LIMIT\",\"side\":\"SELL\",\"price\":\"1.94064\",\"quantity\":\"40\",\"timeInForce\":\"GTE\",\"orderId\":\"F9E3682D70F7D656851D70B38BCC6890295683F1-1023254\"}}","memo":"","source":0,"sequence":1023322},{"txHash":"9B87D17581F2AC73D2999EDE56535E50D9D4DB75150A92A90122190F77D47755","blockHeight":104867508,"txType":"TRANSFER","timeStamp":"2020-08-03T16:32:17.963Z","fromAddr":"bnb1c4czpzvn0ttdcpnv3cy2858l2m9frxdfgk4jr0","toAddr":"bnb1g2ukzn702napq3levm54m2z3p2gam7upern9aq","value":"0.24481570","txAsset":"BNB","txFee":"0.00037500","code":0,"data":null,"memo":"","source":0,"sequence":6},{"txHash":"5C0580AC983C1CF36F1656D9E8B062CD0578839BEFC839EFD6720FF315B45EEB","blockHeight":104867508,"txType":"NEW_ORDER","timeStamp":"2020-08-03T16:32:17.963Z","fromAddr":"bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg","toAddr":null,"value":"169.20330000","txAsset":"AVA-645","txFee":"0.00000000","orderId":"7783C148DC7D2CBC504C0CC569B57A593FE53E70-1509155","code":0,"data":"{\"orderData\":{\"symbol\":\"AVA-645_BUSD-BD1\",\"orderType\":\"LIMIT\",\"side\":\"BUY\",\"price\":\"1.61146\",\"quantity\":\"105\",\"timeInForce\":\"GTE\",\"orderId\":\"7783C148DC7D2CBC504C0CC569B57A593FE53E70-1509155\"}}","memo":"","source":0,"sequence":1509154}]}`
	mockedNodeInfo           = `{"node_info":{"protocol_version":{"p2p":7,"block":10,"app":0},"id":"46ba46d5b6fcb61b7839881a75b081123297f7cf","listen_addr":"10.212.32.84:27146","network":"Binance-Chain-Tigris","version":"0.32.3","channels":"3640202122233038","moniker":"Ararat","other":{"tx_index":"on","rpc_address":"tcp://0.0.0.0:27147"}},"sync_info":{"latest_block_hash":"507BB016F306906569F12883617A4231AB51DAF5FA5004C8F70B17CDF73A8B40","latest_app_hash":"A96FA3DB1FAC12D325845FFEE679EC52CB944BE4B343BC016CE4707FA63EE2BE","latest_block_height":104867535,"latest_block_time":"2020-08-03T16:32:29.834625465Z","catching_up":false},"validator_info":{"address":"B7707D9F593C62E85BB9E1A2366D12A97CD5DFF2","pub_key":[113,242,215,184,236,28,139,153,166,83,66,155,1,24,205,32,31,121,79,64,157,15,234,77,101,177,182,98,242,176,0,99],"voting_power":1000000000000}}`
)

func createMockedAPI() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/v1/node-info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedNodeInfo); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/v2/transactions-in-block/104867508", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedBlockResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/v2/transactions-in-block/105529271", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, wantedBlockResponseMulti); err != nil {
			panic(err)
		}
	})

	return r
}
