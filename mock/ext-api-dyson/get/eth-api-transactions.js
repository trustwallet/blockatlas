/// <coin> RPC Mock
/// See:
/// curl "http://localhost:3000/eth-api/transactions?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl "https://trust-wallet.herokuapp.com/transactions?address=0x0875BCab22dE3d02402bc38aEe4104e1239374a7"
/// curl http://localhost:8420/v1/ethereum/0x0875BCab22dE3d02402bc38aEe4104e1239374a7

module.exports = {
    path: '/eth-api/transactions',
    template: function(params, query, body) {
        //console.log(query)
        if (query.address === '0x0875BCab22dE3d02402bc38aEe4104e1239374a7') {
            return JSON.parse(`
            {
                "docs": [
                    {
                        "operations": [],
                        "contract": null,
                        "_id": "0x2a9fd94735e273526a2cde57c6a19b9d488e0c9a960565c3e19be5e12d4b4b47",
                        "blockNumber": 9551915,
                        "time": 1582624428,
                        "nonce": 227,
                        "from": "0x0875BCab22dE3d02402bc38aEe4104e1239374a7",
                        "to": "0x1717f94202c126ef71d6C562de253Fe95eEbDD5f",
                        "value": "17635730000000000",
                        "gas": "21000",
                        "gasPrice": "4320000000",
                        "gasUsed": "21000",
                        "input": "0x",
                        "error": "",
                        "id": "0x2a9fd94735e273526a2cde57c6a19b9d488e0c9a960565c3e19be5e12d4b4b47",
                        "timeStamp": "1582624428"
                    }
                ],
                "total": 1
            }
            `);
        }
        // fallback
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
