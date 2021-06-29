import React from 'react';
import { Themeable2, withTheme2 } from '@grafana/ui';
import { ServiceComponent } from 'app/types/catalog';
import Pods from './Pods';
import Labels from './ServiceLabels';

interface Props extends Themeable2 {
  component: ServiceComponent;
}

const component = (props: Props) => {
  const { component } = props;
  return (
    <>
      <span>{component.name}</span>
      <Labels labels={component.labels}></Labels>
      <span>{component.address}</span>
      <span>{component.external}</span>
      <span>{component.namespace}</span>
      <span>{component.teams}</span>
      <Pods pods={component.pods}></Pods>
    </>
  );
};

// Theming
export default withTheme2(component);
