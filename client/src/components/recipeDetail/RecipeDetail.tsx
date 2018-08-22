import * as React from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';

import Recipe from '../recipe';

type RecipeDetailFunction = () => JSX.Element;
type RedirectFunction = () => JSX.Element;

const goHome : RedirectFunction = () => <Redirect to='/' />;

const RecipeDetail : RecipeDetailFunction = () => (
    <Switch>
        <Route path='/recipes/:recipeId' component={Recipe} />
        <Route Path='/recipes' exact={true} render={goHome} />
    </Switch>
);

export default RecipeDetail;
