// Dialog //
var Dialog = React.createClass({
    getInitialState: function() {
        return {
            showDialog: false,
            contentType: 'none',
            title: 'title',
            buttonText: 'Save',
        };
    },
    saveInterview : function() {
        var cname = this.refs.body.refs.cname.getDOMNode().value;
        if (!cname || cname.trim().length == 0) {
            alert('Must provide candidate name');
            return false;
        }

        var i = {
            candidate: cname,
            complete: false,
            comments: new Array(),
        };

        var isEmpty = true;
        for (var x = 1; x < 6; x++) {
            var z = 'ic' + x.toString();
            var n = this.refs.body.refs[z].getDOMNode().value;
            if (n && n.length > 0) {
                isEmpty = false;
                var c = {
                    content: '',
                    interviewer: n
                };
                i.comments.push(c);
            }
        }

        if (isEmpty) {
            alert('Need to supply at least one interviewer.')
            return false;
        }

        var that = this;
        Client.SaveInterview(
            i,
            function(data, textStatus, jqXHR) {
                var res = JSON.parse(jqXHR.responseText);
                React.render(
                    <div>
                        <div> Candidate Name: {res.candidate}</div>
                        {
                            res.comments.map(function(c) {
                                return <CommentPanel body={c}></CommentPanel>
                            })
                        }
                    </div>,
                    document.getElementById('content')
                );
                document.getElementById('footer').innerText = 'Interview ' + res.id + ' saved successfully';
            },
            function (jqXHR, textStatus, errorThrown) {
                var res = JSON.parse(jqXHR.responseText);
                var msg = "Error Code: " + res.ErrorCode + " Message: " + res.Message;
                console.log(msg);
                document.getElementById('footer').innerText = msg;
                that.showErrorContent(res);
            }
        );
        this.showLoadingPanel();
        return true;
    },
    showLoadingPanel() {
        React.render(
            <LoadingPanel />,
            document.getElementById('content')
        );
    },
    showErrorContent(res) {
        React.render(
            <div>
                {res.Message}
            </div>,
            document.getElementById('content')
        );
    },
    getInterview() {
        var id = this.refs.body.refs.ic1.getDOMNode().value;
        var that = this;
        Client.GetInterview(
            id,
            '',
            function(data, textStatus, jqXHR) {
                var res = JSON.parse(jqXHR.responseText);
                React.render(
                    <div>
                        <div> Candidate Name: {res.candidate}</div>
                        {
                            res.comments.map(function(c) {
                                return <CommentPanel body={c}></CommentPanel>
                            })
                        }
                    </div>,
                    document.getElementById('content')
                );
                document.getElementById('footer').innerText = 'Interview ' + id + ' retrieved successfully';
            },
            function (jqXHR, textStatus, errorThrown) {
                var res = JSON.parse(jqXHR.responseText);
                var msg = "Error Code: " + res.ErrorCode + " Message: " + res.Message;
                console.log(msg);
                document.getElementById('footer').innerText = msg;
                that.showErrorContent(res);
            }
        );
        this.showLoadingPanel();
    },
    handleSave : function(e) {
        var success = false;
        switch (this.state.contentType) {
            case 'new':
                success = this.saveInterview();
                break;
            case 'find':
                this.getInterview();
                success = true;
                break;
        }

        if (success === true) {
            this.setState({ showDialog: false });
        }
    },
    handleCancel : function(e) {
        switch (this.state.contentType) {
            case 'new':
                for (var x = 1; x < 6; x++) {
                    var i = 'ic' + x.toString();
                    this.refs.body.refs[i].getDOMNode().value = '';
                }
            case 'find':
        }

        this.setState({ showDialog: false });
    },
    render: function() {
        return (
            <div>
                { this.state.showDialog ?
                    <div className="parentDialog">
                        <div className="newDialog">
                            <div className="dialogHeader">
                                {this.state.title}
                            </div>
                            <div className="dialogBody">
                                {(() => {
                                    switch (this.state.contentType) {
                                        case 'none':    return 'none';
                                        case 'new':     return <NewForm ref="body"/>;
                                        case 'find':    return <FindForm ref="body"/>
                                        default:        return 'empty';
                                    }
                                })()}
                            </div>
                            <div className="dialogFooter">
                                <button onClick={this.handleSave} className="dialogButton">{this.state.buttonText}</button>
                                <button onClick={this.handleCancel} className="dialogButton">Cancel</button>
                            </div>
                        </div>
                    </div>
                : null }
            </div>
        );
    }
});
