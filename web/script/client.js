function Client() {

}
Client.GetAuthToken = function() {
    var i = document.cookie.split(';');
    if (!i[0]) {
        return null;
    }

    t = i[0].split('=')

    if (t[0] === 'token') {
        return t[1];
    }
}

Client.SaveInterview = function(i, success, error) {
    var s = JSON.stringify(i);
    $.ajax({
        url : 'https://localhost:8443/interview',
        type: 'POST',
        data : s,
        headers: {
            "Api-Version":  1.0,
            "Authorization": this.GetAuthToken()
        },
        success: success,
        error: error,
    });
}

Client.GetInterview = function(id, cname, success, error) {
    $.ajax({
        url : 'https://localhost:8443/interview?id=' + id,
        type: 'GET',
        headers: {
            "Api-Version":  1.0,
            "Authorization": this.GetAuthToken()
        },
        success: success,
        error: error,
    });
}

Client.GetToken = function(uname, pwd, success, error) {
    $.ajax({
        url : 'https://localhost:8443/token',
        type: 'POST',
        data: 'uname=' + uname + '&pwd=' + pwd,
        headers: {
            "Api-Version":  1.0,
            "Authorization": this.GetAuthToken()
        },
        success: success,
        error: error,
    });
}
