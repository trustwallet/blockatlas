/// Tron API Mock
/// See:
/// curl "http://localhost:3347/mock/tron-api/v1/assets/1002798"
/// curl "https://{tron_rpc}/v1/assets/1002798"
/// curl "http://localhost:8437/v2/tron/tokens/TFFriedwRtWdFuzerDDtkoQTZ29smDZ1MB?Authorization=Bearer"

module.exports = {
    path: "/mock/tron-api/v1/assets/:arg2?",
    template: function(params, query, body) {
        //console.log(params)
        //console.log(query)
        switch (params.arg2) {
            case '1002000':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Ftron-api%2Fv1%2Fassets%2F1002000.json";
                var json = require(fn);
                return json;
                    
            case '1002798':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Ftron-api%2Fv1%2Fassets%2F1002798.json";
                var json = require(fn);
                return json;

            case '1002814':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Ftron-api%2Fv1%2Fassets%2F1002814.json";
                var json = require(fn);
                return json;
        }

        return {error: "Not implemented"};
    }
};
