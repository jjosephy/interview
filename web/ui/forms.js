// New Form //
var NewForm = React.createClass({
    render: function() {
        return (
            <div>
                <div>Candidate Name</div>
                <input type="text" className="iText" ref="cname"/>
                <div className="interviewers">
                    <span className="lInterview">Interviewers</span>
                </div>
                <div>
                    &nbsp;
                </div>
                <div>
                    <div className="addlabel">
                        <label>Interviewer One</label>
                        <input type="text" className="iText" ref="ic1" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Two</label>
                        <input type="text" className="iText" ref="ic2" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Three</label>
                        <input type="text" className="iText" ref="ic3" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Four</label>
                        <input type="text" className="iText" ref="ic4" />
                    </div>
                    <div className="addlabel">
                        <label>Interviewer Five</label>
                        <input type="text" className="iText" ref="ic5" />
                    </div>
                </div>
            </div>
        );
    }
});

// Find Form//
var FindForm = React.createClass({
    render: function() {
        return (
            <div>
                <div className="addlabel">
                    <label>Interviewer Id</label>
                    <input type="text" className="iText" ref="ic1" defaultValue="56b7aaa479ad484c3c000001"/>
                </div>
                <div className="addlabel">
                    <label>Candidate Name</label>
                    <input type="text" className="iText" ref="ic2" />
                </div>
            </div>
        );
    }
});
