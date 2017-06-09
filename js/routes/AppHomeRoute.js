import Relay from 'react-relay';

// Define a root GraphQL query into which
// your containers' query fragments will be composed

export default class extends Relay.Route {
  static queries = {
    allOperations: () => Relay.QL`
      query {
        allOperations
      }
    `,
  };
  static routeName = 'AppHomeRoute';
}
