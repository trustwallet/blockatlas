/// Nebulas API Mock
/// See:
/// curl "http://localhost:3347/mock/nebulas-api/tx?a=n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a&p=0"
/// curl "https://explorer-backend.nebulas.io/api/tx?a=n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a&p=0"
/// curl http://localhost:8437/v1/nebulas/n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a
module.exports = {
    path: "/mock/nebulas-api/tx?",
    template: function(params, query, body) {
        //console.log(query)
        if (query.a === 'n1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Fnebulas-api%2Ftx%3Fa%3Dn1RCYwrpLMpSpUCQ8QUDzGRg6B2PnY8R94a%26p%3D0.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
