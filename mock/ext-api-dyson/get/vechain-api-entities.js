/// Mock for external coin API, transactions
/// See
/// curl https://vethor-pubnode.digonchain.com/blocks/best
/// curl https://vethor-pubnode.digonchain.com/transactions/0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7
/// curl http://localhost:3347/mock/vechain-api/blocks/best
/// curl http://localhost:3347/mock/vechain-api/transactions/0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7
/// curl "http://localhost:8437/v1/vechain/0xB5e883349e68aB59307d1604555AC890fAC47128"

module.exports = {
    path: '/mock/vechain-api/:entity/:id?',
    template: function(params, query, body) {
        //console.log(params)
        if (params.entity === 'blocks') {
            if (params.id === 'best') {
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fvechain-api%2Fblocks%2Fbest.json";
                var json = require(fn);
                return json;
            }
        }
        if (params.entity === 'transactions') {
            switch (params.id) {
                case '0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7':
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fvechain-api%2Ftransactions%2F0x702edd54bd4e13e0012798cc8b2dfa52f7150173945103d203fae26b8e3d2ed7.json";
                    var json = require(fn);
                    return json;

                case '0x004aa0448e458105b098aea2a764a1d54ab95451bee488869f417b351857c3c5':
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fvechain-api%2Ftransactions%2F0x004aa0448e458105b098aea2a764a1d54ab95451bee488869f417b351857c3c5.json";
                    var json = require(fn);
                    return json;
            }
        }
        
        return {error: "Not implemented"};
    }
};
