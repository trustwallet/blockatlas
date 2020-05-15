/// Thundertoken API Mock
/// See:
/// curl "http://localhost:3347/mock/thundertoken-api/transactions?address=0x0b230def08139f18a86536d9cfa150f04435414c"
/// curl "http://localhost:3347/mock/thundertoken-api/tokens?address=0x0b230def08139f18a86536d9cfa150f04435414c"
/// curl "https://{Thundertoken rpc}/transactions?address=0x0b230def08139f18a86536d9cfa150f04435414c"
/// curl "https://{Thundertoken rpc}/tokens?address=0x0b230def08139f18a86536d9cfa150f04435414c"
/// curl "http://localhost:8437/v1/thundertoken/0x0b230def08139f18a86536d9cfa150f04435414c"
/// curl "http://localhost:8437/v2/thundertoken/tokens/0x0b230def08139f18a86536d9cfa150f04435414c?Authorization=Bearer"

module.exports = {
    path: '/mock/thundertoken-api/:command1?',
    template: function(params, query, body) {
        //console.log(params);
        //console.log(query);
        switch (params.command1) {
            case 'transactions':
                if (query.address === '0x0b230def08139f18a86536d9cfa150f04435414c') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fthundertoken-api%2Ftransactions%3Faddress%3D0x0b230def08139f18a86536d9cfa150f04435414c.json";
                    var json = require(fn);
                    return json;
                }
                break;

            case 'tokens':
                if (query.address === '0x0b230def08139f18a86536d9cfa150f04435414c') {
                    var fn = "../../ext-api-data/get/" +
                        "mock%2Fthundertoken-api%2Ftokens%3Faddress%3D0x0b230def08139f18a86536d9cfa150f04435414c.json";
                    var json = require(fn);
                    return json;
                }
                break;
        }

        return {error: "Not implemented"};
    }
};
