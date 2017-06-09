import React from 'react';
import Relay from 'react-relay';

class App extends React.Component {

  createPostList(posts){
    const display = []
    posts.forEach((p) => {
      display.push(<li key={p.id}><a key={p.id} href={p.id}>{p.title}</a> By: {p.author}</li>)
    })
    return display
  }

  render() {
    const posts = this.props.allPosts ? this.createPostList(this.props.allPosts.posts) : []
    return (
      <div>
        <h1>Articles</h1>
        <ul>
          {posts}
        </ul>
      </div>
    );
  }
}

export default Relay.createContainer(App, {
  fragments: {
    latestPost: () => Relay.QL`
      fragment on Post {
        id
        title
        text
        author
      }
    `,
    allPosts: () => Relay.QL`
      fragments on PostList {
        id
        posts{
          id
          title
          author
        }
      }
    `,
  },
});
