/// Theta API Mock
/// See:
/// curl "http://localhost:3347/mock/theta-api/accounttx/0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f?type=2&pageNumber=1&limitNumber=100&isEqualType=true"
/// curl "https://explorer.thetatoken.org:9000/api/accounttx/0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f?type=2&pageNumber=1&limitNumber=100&isEqualType=true"
/// curl http://localhost:8437/v1/theta/0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f
module.exports = {
    path: "/mock/theta-api/accounttx/:address?",
    template: function(params, query, body) {
        //console.log(params)
        if (params.address === '0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f') {
            var fn = "../../ext-api-data/get/" +
                "mock%2Ftheta-api%2Faccounttx%2F0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f%3FisEqualType%3Dtrue%26limitNumber%3D100%26pageNumber%3D1%26type%3D2.json";
            var json = require(fn);
            return json;
        }
        
        return {error: "Not implemented"};
    }
};
