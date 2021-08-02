<script lang="typescript">
  import { tick } from "svelte";
  import { LoadState } from "@app/models/LoadState";
  import RemoteTable from "@app/components/RemoteTable.svelte";
  import type { TableRow } from "@app/components/models/RemoteTableProps";
  import type { IForm, ISubmission } from "@app/models/IForm";
  import { deleteApi, getApi, putApi } from "@app/services/ApiService";
  import { getUrlParameter } from "@app/util/Http";
  import Dialog from "@app/components/layout/Dialog.svelte";
  import FormPreview from "@app/features/form/live/FormPreview.svelte";
  import { fastClone } from "@app/util/Compare";
  import type { Dictionary } from "@app/models/Utility";
  import Shell from "@app/components/Shell.svelte";
  import { SubmissionExcludedFields } from "@app/features/form/live/service/SubmitForm";
  import {formatSubmissionItem, types} from "../../services/SubmissionFormatter";

  export let formId = "";
  export let form: IForm | undefined = undefined;

  let state: LoadState = LoadState.NotStarted;
  let hidden = new Set([
    "submission_id",
    "full_submission_data",
    "meta_unread",
  ]);
  let preview: ISubmission | undefined = undefined;

  async function getRows(): Promise<TableRow[]> {
    formId = getUrlParameter("formId") ?? "";
    if (!formId) {
      return [];
    }
    form = await getApi<IForm>(`form/${formId}`);

    if (!form) {
      return [];
    }

    const url: { message: string } = await getApi(`form/${formId}/submission`);
    const response = await fetch(url.message);
    let submissions: ISubmission[];
    if (response.ok) {
      submissions = await response.json();
    } else {
      submissions = [];
    }

    const read = await getRead(form!);

    const labels: { [key: string]: string } = {};

    if (!form.fields || form.fields?.length === 0) {
      return [];
    }

    form.fields.forEach((f) => {
      if (f.name) {
        labels[f.name] = f.label ?? f.name;
      }
    });

    const deleted = JSON.parse(
      sessionStorage.getItem("deleted_submissions") || JSON.stringify({})
    );

    setTimeout(() => {
      clearUnreadSubmissions();
    }, 2000);

    return submissions
      .filter((d) => !deleted[d.id])
      .map((d) => {
        d.isUnread = read[d.id] !== true;
        Object.keys(d.details).forEach((key) => {
          if (labels[key]) {
            const label = labels[key];
            d.details[label] = d.details[key];
            types[label] =
              form!.fields.find((w) => w.label === label)?.type ?? "";
            delete d.details[key];
          } else {
            const fieldByName = form!.fields.find((w) => w.name === key)?.type;
            if (fieldByName) {
              types[key] = fieldByName;
            }
          }
        });
        d.details["Submission Date"] = new Date(
          d.creationDate
        ).toLocaleString();
        d.details["submission_id"] = d.id;
        d.details["full_submission_data"] = JSON.stringify(d);
        d.details["meta_unread"] = d.isUnread;
        return d.details;
      });
  }

  function sortColumns(columns: string[]) {
    const used = new Set<string>();
    const results: string[] = [];
    results.push("Submission Date");
    used.add("Submission Date");
    form!.fields.forEach((f) => {
      if (SubmissionExcludedFields.includes(f.type)) {
        return;
      }
      results.push(f.label!);
      used.add(f.label!);
    });
    columns.forEach((c) => {
      if (!used.has(c)) {
        results.push(c);
        hidden.add(c);
      }
    });
    return results;
  }

  function onRowClick(row: any) {
    const submission: ISubmission = JSON.parse(row["full_submission_data"]);
    preview = fastClone(submission);
    markRead(submission.formId, [submission]);
  }

  async function onDelete(rows: any[]) {
    const ids = rows.map((r) => r["submission_id"]).filter((r) => r != null);
    await deleteApi(`form/${formId}/submission`, ids);
    const deleted = JSON.parse(
      sessionStorage.getItem("deleted_submissions") || JSON.stringify({})
    );
    ids.forEach((i) => {
      deleted[i] = true;
    });
    sessionStorage.setItem("deleted_submissions", JSON.stringify(deleted));
  }

  async function onMarkRead(rows: any[], value: boolean) {
    const result: Dictionary<boolean> = {};
    rows.forEach((r) => {
      const id = r["submission_id"];
      if (!id) {
        return;
      }
      const unread = Boolean(r["meta_unread"]);
      if (unread === !value) {
        return;
      }
      result[id] = value;
    });
    if (Object.keys(result).length === 0) {
      return;
    }
    await putApi(`form/${formId}/submission/mark/read`, result);
  }

  async function clearUnreadSubmissions() {
    if (form!.unreadSubmissions === 0) {
      return;
    }
    form!.unreadSubmissions = 0;
    putApi(`form/${form!.id}`, form);
  }

  async function markRead(formId: string, submissions: ISubmission[]) {
    submissions = submissions.filter((f) => f.isUnread);
    if (submissions.length === 0) {
      return;
    }
    const result: Dictionary<boolean> = {};
    submissions.forEach((s) => (result[s.id] = true));
    await putApi(`form/${formId}/submission/mark/read`, result);
  }

  async function getRead(form: IForm): Promise<Dictionary<boolean>> {
    return await getApi<Dictionary<boolean>>(`form/${form.id}/submission/read`);
  }
</script>

<Shell header="Form Submissions">
  <RemoteTable
    defaultSortColumn={'Submission Date'}
    searchPlaceHolder={'Search Anything...'}
    columnMeta={{ 'Submission Date': { type: 'date' } }}
    {getRows}
    {sortColumns}
    {onDelete}
    {onRowClick}
    onRead={onMarkRead}
    onFormat={formatSubmissionItem}
    {hidden} />
</Shell>

{#if preview && form}
  <Dialog
    isOpen={true}
    title={`Viewing Submission Details`}
    onClose={async () => {
      preview = undefined;
      await tick();
    }}>
    <div>
      <FormPreview submission={preview} {form} mode={'submission_preview'} />
    </div>
  </Dialog>
{/if}
