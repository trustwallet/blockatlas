/// Kusama RPC API Mock
/// See:
/// curl -H "Content-Type: application/json" -d '{"address": "HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK", "row": 25}' https://kusama.subscan.io/api/scan/transfers
/// curl -H "Content-Type: application/json" -d '{"address": "HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK", "row": 25}' http://localhost:3000/kusama-rpc/scan/transfers
/// curl "http://localhost:8420/v1/kusama/HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK"

module.exports = {
    path: '/kusama-rpc/scan/transfers',
    template: function(params, query, body) {
        //console.log(body);
        switch (body.address) {
            case 'HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK':
                return JSON.parse(`
                    {
                        "code": 0,
                        "message": "Success",
                        "ttl": 1,
                        "data": {
                            "count": 3,
                            "transfers": [
                                {
                                    "from": "EonK7NScfhd7ZRfgnLhm4cKRFJWK1z59zPximUZRg8VjHQj",
                                    "to": "HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK",
                                    "module": "balances",
                                    "amount": "30",
                                    "hash": "0x9908579abffb409aef95394128f9097f0a21cfdf16675588423e31ab0d3c58e2",
                                    "block_timestamp": 1583243826,
                                    "block_num": 1291387,
                                    "extrinsic_index": "1291387-3",
                                    "success": true,
                                    "fee": "10000000000"
                                },
                                {
                                    "from": "EEWyMLHgwtemr48spFNnS3U2XjaYswqAYAbadx2jr9ppp4X",
                                    "to": "HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK",
                                    "module": "balances",
                                    "amount": "7.333459",
                                    "hash": "0xab5dbf2d38f249785c260c3fca9e19a17df298c5294a99eaec2baab29970eb85",
                                    "block_timestamp": 1581154290,
                                    "block_num": 961987,
                                    "extrinsic_index": "961987-2",
                                    "success": true,
                                    "fee": "10000000000"
                                },
                                {
                                    "from": "EEWyMLHgwtemr48spFNnS3U2XjaYswqAYAbadx2jr9ppp4X",
                                    "to": "HZaz6cUo8wJ9zjwoDtA3ZzYkrCLfDYU8b3uPmYttKFFvvRK",
                                    "module": "balances",
                                    "amount": "233.193373",
                                    "hash": "0x08aad18956058ee385387da2f7d613ed17ac5c98ad23683cfea855b0fd529a63",
                                    "block_timestamp": 1579722240,
                                    "block_num": 738814,
                                    "extrinsic_index": "738814-3",
                                    "success": true,
                                    "fee": "10000000000"
                                }
                            ]
                        }
                    }
                `);
        }
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
