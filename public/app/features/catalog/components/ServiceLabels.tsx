import React from 'react';
import { Themeable2, withTheme2 } from '@grafana/ui';
import { CatalogLabels } from 'app/types/catalog';

interface Props extends Themeable2 {
  labels?: CatalogLabels;
}

const labels = (props: Props) => {
  const { labels } = props;
  return (
    <div className="service-label">
      {Object.keys(labels || {}).map((key) => {
        return (
          <span className="service-label">
            {key} : {labels![key]}
          </span>
        );
      })}
    </div>
  );
};

// Theming
export default withTheme2(labels);
