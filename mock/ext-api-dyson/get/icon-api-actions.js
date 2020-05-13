/// Icon API Mock
/// See:
/// curl "http://localhost:3347/mock/icon-api/address/txList?address=hxee691e7bccc4eb11fee922896e9f51490e62b12e&count=25"
/// curl "https://tracker.icon.foundation/v3/address/txList?address=hxee691e7bccc4eb11fee922896e9f51490e62b12e&count=25"
/// curl http://localhost:8437/v1/icon/hxee691e7bccc4eb11fee922896e9f51490e62b12e
module.exports = {
    path: "/mock/icon-api/address/:action?",
    template: function(params, query, body) {
        //console.log(query)
        if (params.action === 'txList') {
            if (query.address === 'hxee691e7bccc4eb11fee922896e9f51490e62b12e') {
                var fn = "../../ext-api-data/get/" +
                    "mock%2Ficon-api%2Faddress%2FtxList%3Faddress%3Dhxee691e7bccc4eb11fee922896e9f51490e62b12e%26count%3D25.json";
                var json = require(fn);
                return json;
            }
        }
        
        return {error: "Not implemented"};
    }
};
