// Header Component //
var Header = React.createClass({displayName: 'Header',
    newHandler : function(e) {
        this.setDialog('new', 'Create New Interview', 'Save');
    },
    findHandler : function(e) {
        this.setDialog('find', 'Find Interview', 'Go');
    },
    logoutHandler : function(e) {
        deleteCookie();
    },
    setDialog : function(type, title, buttonText) {
        var e = this.props.dialog;
        if (e) {
            e.setState({ contentType: type });
            e.setState({ showDialog: true });
            e.setState({ title: title });
            e.setState({ buttonText: buttonText });
        } else {
            console.log('no dialog object');
        }
    },
    render: function() {
        return (
            <div className="header" >
                { this.props.show === 'true' ?
                    <div>
                    <MenuItem label="New" handler={this.newHandler} />
                    <MenuItem label="Find" handler={this.findHandler} />
                    <MenuItem label="Logout" handler={this.logoutHandler} />
                    </div>
                : null }
            </div>
        );
    }
});
