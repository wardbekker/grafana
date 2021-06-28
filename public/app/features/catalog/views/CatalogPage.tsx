import React, { useEffect } from 'react';
import { hot } from 'react-hot-loader';
import { connect } from 'react-redux';
import { Themeable2, withTheme2 } from '@grafana/ui';
import { StoreState } from 'app/types';
import { loadCatalog } from '../state/actions';
import { Catalog } from 'app/types/catalog';

interface Props extends Themeable2 {
  name: string;
  catalog: Catalog;
  loadCatalog: typeof loadCatalog;
}

export const UnthemedCatalogPage = (props: Props) => {
  const { loadCatalog } = props;
  useEffect(() => {
    loadCatalog();
  }, [loadCatalog]);

  console.log(props.catalog);
  return (
    <>
      <h2>catalog</h2>
      {props.catalog.map((svc, key) => {
        return (
          <div key={key}>
            <h2>Service: {svc.name}</h2>
            {svc.components.map((cmp, i) => {
              return <h3 key={i}>Component: {cmp.name}</h3>;
            })}
          </div>
        );
      })}
    </>
  );
};

// Global state stuff, redux
const mapDispatchToProps = {
  loadCatalog,
};

export const mapStateToProps = (state: StoreState) => ({
  catalog: state.catalog.catalog,
});

// Theming
export const CatalogPage = withTheme2(UnthemedCatalogPage);

// Hot module reloading ?
export default hot(module)(connect(mapStateToProps, mapDispatchToProps)(CatalogPage));
