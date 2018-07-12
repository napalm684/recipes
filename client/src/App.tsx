import * as React from 'react';
import Recipes from './components/recipes';

export default class App extends React.Component {
  public render() {
    return (
      <div>
        <Recipes />
      </div>
    );
  }
}
