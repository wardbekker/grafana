import { ThunkResult } from 'app/types';
import { getCatalog } from './api';
import { catalogLoaded } from './reducers';

export function loadCatalog(): ThunkResult<void> {
  return async (dispatch) => {
    const catalog = await getCatalog();

    dispatch(catalogLoaded(catalog));
  };
}
