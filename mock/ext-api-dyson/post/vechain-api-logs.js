/// Vechain RPC API Mock
/// See:
/// curl -H "Content-Type: application/json" -d '{"options": {"offset": 0, "limit": 15 }, "criteriaSet": [{"sender": "0xB5e883349e68aB59307d1604555AC890fAC47128"},{"recipient": "0xB5e883349e68aB59307d1604555AC890fAC47128"}], "range": {"unit": "block", "from": 0, "to": 5466405 }, "order": "desc"}' https://vethor-pubnode.digonchain.com/logs/transfer
/// curl -H "Content-Type: application/json" -d '{"options": {"offset": 0, "limit": 15 }, "criteriaSet": [{"sender": "0xB5e883349e68aB59307d1604555AC890fAC47128"},{"recipient": "0xB5e883349e68aB59307d1604555AC890fAC47128"}], "range": {"unit": "block", "from": 0, "to": 5466405 }, "order": "desc"}' http://localhost:3347/mock/vechain-api/logs/transfer
/// curl "http://localhost:8437/v1/vechain/0xB5e883349e68aB59307d1604555AC890fAC47128"

module.exports = {
    path: '/mock/vechain-api/logs/:entity',
    template: function(params, query, body) {
        //console.log(params);
        //console.log(body);
        if (params.entity === 'transfer') {
            // TODO check sender/recipient
            if (body["criteriaSet"][0]["sender"] === '0xB5e883349e68aB59307d1604555AC890fAC47128' ||
                body["criteriaSet"][1]["recipient"] === '0xB5e883349e68aB59307d1604555AC890fAC47128') {
                var fn = "../../ext-api-data/post/" +
                    "mock%2Fvechain-api%2Flogs%2Ftransfer.json";
                var json = require(fn);
                return json;
            }
        }
        return {error: "Not implemented"};
    }
};
