import { IRecipe } from "../entities/Recipe";
import FetchError from "../exceptions/FetchError";

const baseUrl: string = '/api/recipes';

export default class RecipeService {
     public async getAll() : Promise<IRecipe[]> {
        const response : Response = await fetch(
            baseUrl,
            { method: 'GET' }
        );

        if(response.ok) {
            return response.json();
        }

        throw new FetchError(`RecipeService.getAll() failure -- \ 
            ${response.status}: ${response.statusText}`);
    }
}
