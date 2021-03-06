import * as React from 'react';
import { Route, Switch } from 'react-router-dom';

import RecipeDetail from '../recipeDetail';
import Recipes from '../recipes/Recipes';

type MainFunction = () => JSX.Element;

const Main : MainFunction = () => (
    <main>
      <Switch>
        <Route exact={true} path="/" component={Recipes} />
        <Route path="/recipes" component={RecipeDetail} />
      </Switch>
    </main>
);

export default Main;
