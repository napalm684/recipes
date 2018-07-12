import * as React from 'react';
import './RecipeSearch.css';

export interface IRecipeSearchProps{
    handleSearch: (searchQry: string) => void;
}

export default function RecipeSearch(props: IRecipeSearchProps){
    return(
        <div className="recipe-search__container">
            <input type="text" className="recipe-search__input" placeholder="Search" />
        </div>
    );
}
