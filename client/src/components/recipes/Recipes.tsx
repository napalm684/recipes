import * as React from 'react';
import ListItem from '../../shared/components/list-item';
import RecipeSearch from './recipe-search';
import './Recipes.css';

// export interface IRecipeListProps {
// }

const test: (s: string) => void = (s: string) => {
  alert("What up");
};

export default class Recipes extends React.Component<any, any> {
  public render() {
    return (
      <div className="recipes">
        <h1>Recipes</h1>
        <RecipeSearch handleSearch={test} />
        <ListItem label="Pizza" />
        <ListItem label="Slum Gaullion" />
      </div>
    );
  }
}
