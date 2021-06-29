import React, { useEffect } from 'react';
import { hot } from 'react-hot-loader';
import { connect } from 'react-redux';
import { CustomScrollbar, Themeable2, withTheme2 } from '@grafana/ui';
import { StoreState } from 'app/types';
import { loadCatalog } from '../state/actions';
import { Catalog } from 'app/types/catalog';
import ServiceComponent from '../components/ServiceComponent';

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
    <CustomScrollbar autoHeightMin={'100%'}>
      <h2>catalog</h2>
      {props.catalog.map((svc, key) => {
        return (
          <div key={key}>
            <h2>Service: {svc.name}</h2>
            {svc.components.map((cmp, i) => {
              return <ServiceComponent component={cmp} key={i} />;
            })}
          </div>
        );
      })}
    </CustomScrollbar>
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
