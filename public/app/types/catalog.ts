export type Catalog = CatalogService[];

export interface CatalogService {
  name: string;
  components: CatalogComponent[];
}

export type CatalogLabels = {
  [key: string]: string;
};

export enum CatalogPodStatus {
  Green = 0,
  Yellow = 1,
  Red = 2,
}

export interface CatalogPod {
  name: string;
  status: CatalogPodStatus;
}

export interface CatalogComponent {
  name: string;
  labels: CatalogLabels;
  teams: string[];
  namespace: string;
  address: string;
  catalogPod: CatalogPod[];
  external: boolean;
}

export interface CatalogState {
  catalog: Catalog;
}
