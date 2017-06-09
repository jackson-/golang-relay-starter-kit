import React from 'react';
import Relay from 'react-relay';

class App extends React.Component {

  createOpList(ops){
    const display = []
    ops.forEach((o) => {
      display.push(<li key={o.id}>{o.title}</li>)
    })
    return display
  }

  render() {
    const operations = this.props.allOperations ? this.createOpList(this.props.allOperations.operations) : []
    return (
      <div>
        <h1>Operations</h1>
        <ul>
          {operations}
        </ul>
      </div>
    );
  }
}

export default Relay.createContainer(App, {
  fragments: {
    allOperations: () => Relay.QL`
      fragments on OperationList {
        id
        operations{
          id
          title
        }
      }
    `,
  },
});
