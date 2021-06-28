import { AnyAction, createAction } from '@reduxjs/toolkit';
import { Catalog, CatalogState } from 'app/types/catalog';

export const catalogLoaded = createAction<Catalog>('catalog/catalogLoaded');

export const catalogReducers = (state: CatalogState = initialCatalogState, action: AnyAction): CatalogState => {
  if (catalogLoaded.match(action)) {
    return {
      ...state,
      catalog: action.payload,
    };
  }

  return state;
};

export const initialCatalogState: CatalogState = {
  catalog: [],
};

export default {
  catalog: catalogReducers,
};
