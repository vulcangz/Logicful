import type { Group } from "./Group";
import type { IField } from "./IField";

export type IForm = {
  fields: IField[];
  title?: string;
  description?: string;
  enableLogic?: boolean;
  id?: string;
  changeDate?: string;
  creationDate?: string;
  createBy?: string;
  changeBy?: string;
  groups?: Group[];
  submissionCount?: number;
  unreadSubmissions?: number;
  loaded?: boolean;
  url?: string;
  initialized?: boolean;
  folder?: string;
  workflow?: Workflow;
  submissionConfig: SubmissionConfig;
};

export type Team = {
  id: string;
  name: string;
};

export type SubmissionConfig = {
  afterSubmitAction: any;
  afterSubmitConfig: any;
};

export type Workflow = {
  integrations: Integration[];
};

export type Integration = {
  name: string;
  label: string;
  config: any;
  enabled?: boolean;
  description: string;
  editor: any;
};

export type ISubmission = {
  id: string;
  formId: string;
  details: { [key: string]: any };
  meta: {
    env: any;
    ipAddress: string;
  };
  isUnread: boolean;
  [key: string]: any;
};
