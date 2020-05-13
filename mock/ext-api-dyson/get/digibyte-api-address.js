/// Mock for external Digibyte API
/// See:
/// curl "http://{digibyte rpc}/api/v2/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi?details=txs"
/// curl "http://localhost:3347/mock/digibyte-api/v2/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi?details=txs"
/// curl "http://localhost:8437/v1/digibyte/address/DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi"

module.exports = {
    path: '/mock/digibyte-api/v2/address/:address?',
    template: function(params, query, body) {
        switch (params.address) {
            case 'DEs1RJKuASSjfphFJdxX9eidrjWewMZgAi':
                var fn = "../../ext-api-data/get/" +
                    "mock%2Fdigibyte-api%2Fv2%2Faddress%2FDEs1RJKuASSjfphFJdxX9eidrjWewMZgAi%3Fdetails%3Dtxs.json";
                var json = require(fn);
                return json;
        }
        return {error: "Not implemented"};
    }
}
