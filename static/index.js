new Vue({
    el: '#app',
    data: {
        longUrl: '',
        result: ''
    },
    methods: {
        tiny: function () {
            var self = this;
            axios.post('/v1/tiny', {
                longUrl: this.longUrl
            }).then(function (result) {
                data = result.data;
                if (location.port)
                    return self.result = location.protocol + '//' + document.domain + ':' + location.port + '/v1/r/' + data.tinyUrl;
                return self.result = location.protocol + '//' + document.domain + '/v1/r/' + data.tinyUrl;
            })
        }
    }
});
