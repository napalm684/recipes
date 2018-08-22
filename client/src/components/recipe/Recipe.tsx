import * as React from 'react';
import { match } from 'react-router';

import { IRecipeRouterParams } from './RecipeRouterParams';

export interface IRecipeProps {
    match: match<IRecipeRouterParams>;
}

export default class Recipe extends React.Component<IRecipeProps> {
    public render() : JSX.Element {
        return <span>{this.props.match.params.recipeId}</span>;
    }
}
