import type { LabelValue } from "./IField";

export interface LogicRule {
  field?: string;
  condition: string;
  value: any;
}

export interface LogicConditional {
  value: string;
  label: string;
  helper?: string;
  placeholder?: string;
  valueInput?: string;
  options?: (index: number) => Promise<LabelValue[]>;
}

export interface LogicRuleOptions {
  valueType: string;
  showValue: boolean;
  options: () => Promise<LabelValue[]>;
  helperText: string;
  placeholder: string;
}
