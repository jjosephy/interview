// Menu Item Component //
var MenuItem  = React.createClass({
    handleClick : function(e) {
        this.props.handler();
    },
    render: function() {
        return (
            <span onClick={this.handleClick} className="menu_item">{this.props.label}</span>
        );
    }
});
