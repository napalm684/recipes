import * as React from 'react';
import Header from './components/header';
import Main from './components/main';

export default class App extends React.Component {
  public render() {
    return (
      <div>
        <Header />
        <Main />
      </div>
    );
  }
}
