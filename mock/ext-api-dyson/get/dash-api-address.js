/// Mock for external Dash API
/// See:
/// curl "http://{dash rpc}/api/v2/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG?details=txs"
/// curl "http://localhost:3347/dash-api/v2/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG?details=txs"
/// curl "http://localhost:8437/v1/dash/address/XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"

module.exports = {
    path: '/dash-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG':
                return JSON.parse(`
                {
                    "page": 1,
                    "totalPages": 333,
                    "itemsOnPage": 2,
                    "address": "XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG",
                    "balance": "24107591472",
                    "totalReceived": "137513751052",
                    "totalSent": "113406159580",
                    "unconfirmedBalance": "0",
                    "unconfirmedTxs": 0,
                    "txs": 665,
                    "transactions": [
                        {
                            "txid": "8a1859bb849e207b7771c0301a6109a039eec955234b7848715b150f20fabeca",
                            "version": 3,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "isAddress": false,
                                    "coinbase": "035c33131a4d696e656420627920416e74506f6f6c347c000b02203c94a613416a0000be010000"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "155331342",
                                    "n": 0,
                                    "hex": "76a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac",
                                    "addresses": [
                                        "XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "155331350",
                                    "n": 1,
                                    "hex": "76a91404dfe99f6dda028e4e6a90595096e63787a4a01c88ac",
                                    "addresses": [
                                        "Xb8cmjtK67y9T2haqrcDicoAMByDesLcAe"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "000000000000000f2a4362c05fd3aeca12640061c9c119862cf20b2cc95560a8",
                            "blockHeight": 1258332,
                            "confirmations": 1181,
                            "blockTime": 1587515977,
                            "value": "310662692",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "03000500010000000000000000000000000000000000000000000000000000000000000000ffffffff27035c33131a4d696e656420627920416e74506f6f6c347c000b02203c94a613416a0000be010000ffffffff020e2b4209000000001976a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac162b4209000000001976a91404dfe99f6dda028e4e6a90595096e63787a4a01c88ac000000004602005c3313008c104ddcb4ef489cb517b4ae6aeec851ab8010748b00d20c7bb10b2a86c9ed6478777117ab77017feac02c03ef31f19a0ddfa3e1039f8d364a79cb22052233e3"
                        },
                        {
                            "txid": "bc017866f1eb2118c59d3c7b6127b96ecce595062abc626f468345782f0bba9f",
                            "version": 3,
                            "vin": [
                                {
                                    "sequence": 4294967295,
                                    "n": 0,
                                    "isAddress": false,
                                    "coinbase": "03033313042c599f5e0100048e1c1d000000000000"
                                }
                            ],
                            "vout": [
                                {
                                    "value": "155462330",
                                    "n": 0,
                                    "hex": "76a914aeb4b16eb331e7be66082f1dc132ef245e722d7188ac",
                                    "addresses": [
                                        "XrcbsQdrFYEzbqA9nCJi8zDtnRZzNKkCtG"
                                    ],
                                    "isAddress": true
                                },
                                {
                                    "value": "155462346",
                                    "n": 1,
                                    "spent": true,
                                    "hex": "76a91417158751f496276811c148495b28915314dd9fc088ac",
                                    "addresses": [
                                        "XcnuAVsG4pRCqxQmQB6PeTX8FJo2ssYPyB"
                                    ],
                                    "isAddress": true
                                }
                            ],
                            "blockHash": "000000000000001c141317365929a3d7937f21be525873142bdaebbdcebe7a1d",
                            "blockHeight": 1258243,
                            "confirmations": 1270,
                            "blockTime": 1587501361,
                            "value": "310924676",
                            "valueIn": "0",
                            "fees": "0",
                            "hex": "03000500010000000000000000000000000000000000000000000000000000000000000000ffffffff1503033313042c599f5e0100048e1c1d000000000000ffffffff02ba2a4409000000001976a914aeb4b16eb331e7be66082f1dc132ef245e722d7188acca2a4409000000001976a91417158751f496276811c148495b28915314dd9fc088ac000000004602000333130024b54c2354b15fb05d02604e811682ea82b3e81abf3610a4019a9fde7963be6ef022b80b9ef1524c1629323e53296cebee22b80dde230ed8eaee578d5a0f4457"
                        }
                    ]
                }
                `);
        }
        return {error: "Not implemented"};
    }
}
