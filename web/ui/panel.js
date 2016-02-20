// Component Panel //
var CommentPanel = React.createClass({
    render: function() {
        return (
            <div className="commentPanel">
                <div className="cpLabel">
                    <label>Interviewer Name:</label>
                    <input type="text" className="iText" defaultValue={this.props.body.interviewer}/>
                </div>
                <div>
                    <textarea className="txt">
                    </textarea>
                </div>
                <div className="select"><select>
                        <option>No Hire</option>
                        <option>Hire</option>
                    </select>
                    <button className="buttonRight">Save</button>
                </div>
                <hr/>
            </div>
        );
    }
});

// Loading Panel //
var LoadingPanel = React.createClass({
    render: function() {
        return (
            <div className="loading">
                <img src="images/loading.gif" />
            </div>
        );
    }
});
