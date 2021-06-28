import { Catalog } from 'app/types/catalog';
import { getBackendSrv } from '../../../core/services/backend_srv';

export async function getCatalog(): Promise<Catalog> {
  const result = await getBackendSrv().get<Catalog>(`/api/catalog`);

  console.log(result);
  return result;
}
