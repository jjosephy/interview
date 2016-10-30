var Login = React.createClass({displayName: 'Login',
	setCookie: function () {
		var n = React.findDOMNode(this.refs.uname).value;
		var p = React.findDOMNode(this.refs.pwd).value;
		var t = Client.GetToken(
			n,
			p,
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
			<div className='login'>
				<div className='label'>User Name</div>
				<input type='text' ref='uname'/>
				<div className="label">Password</div>
				<input type='text' ref='pwd' />
				<div><input className='button' type='button' value='Log In' onClick={this.setCookie}/></div>
			</div>
		);
	}
});
