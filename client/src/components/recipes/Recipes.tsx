import * as React from 'react';

import './Recipes.css';

import { IRecipe } from '../../entities/Recipe';
import RecipeService from '../../services/RecipeService';
import ListItem from '../../shared/components/list-item';
import RecipeSearch from './recipe-search';

// export interface IRecipesProps {
// }

export interface IRecipesState {
  recipes : IRecipe[];
}

const test: (s: string) => void = (s: string) => {
  alert("What up");
};

export default class Recipes extends React.Component<any, IRecipesState> {
  constructor(props: any, state: IRecipesState){
    super(props, state);
    this.state = {
      recipes: []
    };
  }

  public async componentDidMount() {
    const svc = new RecipeService();
    const initState = { recipes: await svc.getAll() };
    this.setState(initState);
  }

  public render() {
    return (
      <div className="recipes">
        <h1>Recipes</h1>
        <RecipeSearch handleSearch={test} />
        {this.buildListItems()}
      </div>
    );
  }

  public buildListItems() : JSX.Element[] {
    const listItems : JSX.Element[] = [];

    if(this.state.recipes.length > 0) {
      this.state.recipes.forEach(recipe => {
        listItems.push(<ListItem key={recipe.ID} label={recipe.Name} />)
      });
    }

    return listItems;
  }
}
