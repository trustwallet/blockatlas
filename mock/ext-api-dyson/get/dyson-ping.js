/// Health-check for Dyson

module.exports = {
    path: '/dyson-ping/ping',
    template: function(params, query) {
        return {
            status: true,
            msg: 'Dyson is alive'
        };
    }
}
