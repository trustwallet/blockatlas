/// Vechain RPC API Mock
/// See:
/// curl -H "Content-Type: application/json" -d '{"options": {"offset": 0, "limit": 15 }, "criteriaSet": [{"sender": "0xB5e883349e68aB59307d1604555AC890fAC47128"},{"recipient": "0xB5e883349e68aB59307d1604555AC890fAC47128"}], "range": {"unit": "block", "from": 0, "to": 5466405 }, "order": "desc"}' https://vethor-pubnode.digonchain.com/logs/transfer
/// curl -H "Content-Type: application/json" -d '{"options": {"offset": 0, "limit": 15 }, "criteriaSet": [{"sender": "0xB5e883349e68aB59307d1604555AC890fAC47128"},{"recipient": "0xB5e883349e68aB59307d1604555AC890fAC47128"}], "range": {"unit": "block", "from": 0, "to": 5466405 }, "order": "desc"}' http://localhost:3000/vechain-api/logs/transfer
/// curl "http://localhost:8420/v1/vechain/0xB5e883349e68aB59307d1604555AC890fAC47128"

module.exports = {
    path: '/vechain-api/logs/:entity',
    template: function(params, query, body) {
        //console.log(params);
        //console.log(body);
        if (params.entity === 'transfer') {
            // TODO check sender/recipient
            if (body["criteriaSet"][0]["sender"] === '0xB5e883349e68aB59307d1604555AC890fAC47128' ||
                body["criteriaSet"][1]["recipient"] === '0xB5e883349e68aB59307d1604555AC890fAC47128') {
                return JSON.parse(`
                    [
                        {
                            "sender": "0xb5e883349e68ab59307d1604555ac890fac47128",
                            "recipient": "0x2c7a8d5cce0d5e6a8a31233b7dc3dae9aae4b405",
                            "amount": "0x12b1815d00738000",
                            "meta": {
                                "blockID": "0x004313a4bd4286e821b684cc1749deb3df12fa2a8114435fbd35baa155e82016",
                                "blockNumber": 4395940,
                                "blockTimestamp": 1574410670,
                                "txID": "0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7",
                                "txOrigin": "0xb5e883349e68ab59307d1604555ac890fac47128",
                                "clauseIndex": 0
                            }
                        },
                        {
                            "sender": "0xb5e883349e68ab59307d1604555ac890fac47128",
                            "recipient": "0x00bae5ed35736e4ef17af1be0c6f50e0fb73d685",
                            "amount": "0x38f6ea18e810b6f00",
                            "meta": {
                                "blockID": "0x0042249bee56223e0ed7a9c7fcfffe8e61b9fd95d29d24843c558ff2c46ea094",
                                "blockNumber": 4334747,
                                "blockTimestamp": 1573795570,
                                "txID": "0x004aa0448e458105b098aea2a764a1d54ab95451bee488869f417b351857c3c5",
                                "txOrigin": "0xb5e883349e68ab59307d1604555ac890fac47128",
                                "clauseIndex": 0
                            }
                        }
                    ]
                `);
            }
        }
        var return4Codacy = {error: "Not implemented"};
        return return4Codacy;
    }
};
