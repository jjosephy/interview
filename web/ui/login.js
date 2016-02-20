var Login = React.createClass({displayName: 'Login',
    setCookie: function () {
        var t = Client.GetToken(
            'w',
            'w',
            function(data, textStatus, jqXHR) {
                t = data;
                var ex = new Date();
                var date = new Date();
                date.setDate(ex.getDate() + 1);
                document.cookie = 'token=' + t;
                document.cookie = 'Path=/';
                document.cookie = 'Expires=' + date;
                alertCookie();
                renderAuthencticated();
            },
            function (jqXHR, textStatus, errorThrown) {
                var res = JSON.parse(jqXHR.responseText);
                var msg = "Error Code: " + res.ErrorCode + " Message: " + res.Message;
                console.log(msg);
                document.getElementById('footer').innerText = msg;
            }
        );
    },
    render: function() {
        return (
            <div className="login">
                <div className="label">User Name</div>
                <input type="text" />
                <div className="label">Password</div>
                <input type="text" />
                <div><input className="button" type="button" value="Log In" onClick={this.setCookie}/></div>
            </div>
        );
    }
});
