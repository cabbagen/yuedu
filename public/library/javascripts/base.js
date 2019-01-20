
var cbTool = {
    getTocken: function() {
        return window.sessionStorage.getItem('tocken') || '';
    },
    request: function (url, method, params) {
        return $.ajax({
            url: url,
            method: method,
            data: params,
            headers: {
                'Content-Type': 'application/json',
                'token': cbTool.getTocken(),
            },
        });
    }
};

