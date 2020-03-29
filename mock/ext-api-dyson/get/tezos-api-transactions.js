/// Mock for external Tezos API, transactions
/// See
/// curl "http://api.tezos.id/mooncake/mainnet/v1/transactions?account=tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8&n=50&p=0"
/// curl "http://localhost:3000/tezos-api/v1/transactions?account=tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8&n=50&p=0"
/// curl "http://localhost:8420/v1/tezos/tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8"

module.exports = {
    path: '/tezos-api/v1/transactions?',
    template: function(params, query, body) {
        //console.log(query)
        if (query.account === 'tz1foWxaV3VQyWqFbWTERS6YDJjPT6C7jPp8') {
            return JSON.parse(`
                [
                    {
                        "tx": {
                            "kind": "transaction",
                            "status": "applied"
                        },
                        "op": {
                            "opHash": "ooLrNAP233Qvoz3AGvjRjhk1fjG7z19UfyLEGBm1rwfHn4NSVhd"
                        }
                    }
                ]
            `)
        }
        // fallback
        return {error: "Not implemented"}
    }
};
