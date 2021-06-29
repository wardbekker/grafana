import React from 'react';
import { Themeable2, withTheme2 } from '@grafana/ui';
import { CatalogPod } from 'app/types/catalog';

interface Props extends Themeable2 {
  pods?: CatalogPod[];
}

const pod = (props: Props) => {
  const { pods } = props;

  return (
    <div className="pod-list">
      {pods?.map((pod, i) => {
        return (
          <span className="pod" key={i}>
            {pod.name}
          </span>
        );
      })}
    </div>
  );
};

// Theming
export default withTheme2(pod);
