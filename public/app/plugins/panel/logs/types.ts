import { LogsSortOrder, LogsDedupStrategy } from '@grafana/data';

export interface Options {
  showLabels: boolean;
  showTime: boolean;
  wrapLogMessage: boolean;
  prettifyLogMessage: boolean;
  enableLogDetails: boolean;
  sortOrder: LogsSortOrder;
  dedupStrategy: LogsDedupStrategy;
}
