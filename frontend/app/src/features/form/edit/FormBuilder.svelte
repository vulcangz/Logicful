<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { randomStringSmall, randomString } from "@app/util/Generate";
  import type { IForm } from "@app/models/IForm";
  import { onMount, tick } from "svelte";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import DynamicForm from "./DynamicForm.svelte";
  import formStore from "@app/store/FormStore";
  import { fastClone } from "@app/util/Compare";
  import { saveForm, saveToLocalStorage } from "./services/SaveForm";
  import ToastManager from "@app/components/ToastManager.svelte";
  import { setFieldDefaults } from "@app/features/form/edit/services/DefaultFieldValueFactory";
  import { getApi } from "@app/services/ApiService";
  import { getUrlParameter } from "@app/util/Http";
  import { startPreviewSaver } from "./services/PreviewSaver";
  import { DynamicFormMode } from "@app/components/models/ComponentProps";
  import { dispatch, subscribeComponent } from "@app/event/EventBus";
  import { navigate } from "svelte-routing";

  let dropped = false;
  let loadingActive: boolean = false;
  let order = [];
  let dragForm: IForm | undefined;
  let demo = false;
  let lastLength = 0;
  let dirty = false;

  let form: IForm;

  async function loadForm() {
    loadingActive = true;
    const formId = getUrlParameter("formId");
    if (!formId) {
      navigate("/form/create", { replace: true });
      return;
    }
    demo = formId === "demo";
    try {
      await loadFromApi(formId);

      if (!form) {
        return;
      }

      if (!form.fields) {
        form.fields = [];
      }

      form.fields = form.fields.map((w: IField) => {
        w.selected = false;
        return w;
      });

      form.loaded = true;
      formStore.setForm(form);
      dispatch("form_loaded", {
        form,
      });
      saveToLocalStorage(form);
      startPreviewSaver();
      addPlaceHolder();
    } finally {
      loadingActive = false;
    }
  }

  async function loadFromApi(formId: string) {
    if (!formId) {
      return;
    }
    if (demo) {
      form = {
        groups: [],
        submissionConfig: undefined,
        url: "",
        fields: [
          {
            id: "bbe7e84e-3099-4e68-8349-0acba49ce595",
            type: "section-header",
            name: "new-field-1squ",
            label: "New Field nkjq",
            selected: false,
            helperText:
              "This is a demo form that you can use to experiment with the editor. You will not be able to save the form in demo mode. You may still preview the live form by clicking the Preview button.",
            header: "Demo Form",
            updated: false,
          },
          {
            id: "f3e0c82f-29a2-4f1c-b68e-a7c99c28c33e",
            type: "section-header",
            name: "new-field-0z4o",
            label: "New Field 78oa",
            selected: false,
            helperText:
              "We have added a few fields to get you started, feel free to delete them.",
            header: "",
            updated: true,
          },
          {
            id: "c7bfb17d-7069-43cb-b9d1-f36dbb178733",
            type: "string",
            name: "new-field-8gj1",
            label: "First Name",
            selected: false,
            updated: false,
            required: true,
            rows: 1,
          },
          {
            id: "ccf6eff1-a153-469c-8c5d-f7a1f2d82bb8",
            type: "file",
            name: "new-field-k9ko",
            label: "Your Resume",
            selected: false,
            updated: true,
            required: true,
          },
          {
            id: "dcbc8276-58c7-4fc8-87c2-34367853b5dd",
            type: "checkbox-group",
            name: "new-field-wfuy",
            label: "Which positions would you like to apply for?",
            selected: false,
            options: ["Assisant Manager", "Manager", "Executive Assistant"],
            updated: true,
            required: true,
            includeOther: true,
            helperText: "",
          },
          {
            id: "1a36dbf2-ca8c-4040-86c8-4f5e71259354",
            type: "switch",
            name: "new-field-l5mk",
            label: "Subscribe to receive application updates by email",
            selected: false,
            updated: false,
            defaultValue: true,
            logic: {},
          },
          {
            id: "880d85ce-1892-4fe9-ad8c-6b672bb30943",
            type: "string",
            name: "new-field-pdu9",
            label: "Email Address",
            selected: true,
            updated: true,
            rows: 1,
            logic: {
              action: "show-all-match",
              rules: [
                {
                  field: "1a36dbf2-ca8c-4040-86c8-4f5e71259354",
                  value: "",
                  condition: "isTrue",
                },
              ],
            },
          },
        ],
        id: "demo",
        title: "My New Form",
        description: "Your Form Description (Optional)",
        changeDate: "2020-10-15T21:26:53Z",
        changeBy: "maddox2",
        creationDate: "2020-10-15T20:18:14Z",
        createBy: "maddox2",
        submissionCount: 5,
        unreadSubmissions: 0,
        folder: "42feae3b-3219-45c6-ac59-24a011a096e6",
        workflow: { integrations: [] },
        loaded: true,
        initialized: true,
        enableLogic: true
      };
    } else {
      form = await getApi(`form/${formId}`);
    }
  }

  function removePlaceHolder() {
    const placeholder = form.fields.findIndex((w) => w.type === "placeholder");
    if (placeholder !== -1) {
      const temp = fastClone(form.fields);
      temp.splice(placeholder, 1);
      form.fields = temp;
      dispatch("form_placeholder_changed", {
        added: false,
      });
    }
  }

  function addPlaceHolder() {
    if (form.fields.filter((w) => w.type !== "placeholder").length !== 0) {
      removePlaceHolder();
      return;
    }
    if (form.fields.find((w) => w.type === "placeholder")) {
      return;
    }
    form.fields = form.fields.concat([
      {
        name: "placeholder-field",
        label: "You have no fields",
        type: "placeholder",
        id: "placeholder",
      },
    ]);
    dispatch("form_placeholder_changed", {
      added: true,
    });
  }

  subscribeComponent("form_updated", (props) => {
    form = props;
    addPlaceHolder();
    dirty = true;
  });

  subscribeComponent("field_delete", (params) => {
    const index = form.fields.findIndex((w) => w.id === params.field.id);
    const temp = [...form.fields];
    temp.splice(index, 1);
    form.fields = temp;
    formStore.setForm(form);
    dispatch("unselect_field", {});
  });

  subscribeComponent("right_sidebar_loaded", () => {
    form &&
      dispatch("form_loaded", {
        form,
      });
  });

  subscribeComponent("add_field", (params) => {
    form.fields = form.fields.map((w) => {
      w.selected = false;
      return w;
    });
    const id = randomString();
    let field: IField = {
      name: "new-field-" + randomStringSmall(),
      label: "New Field " + randomStringSmall(),
      type: params.type,
      id,
      selected: true,
      value: undefined,
      expanded: true,
    };
    field = setFieldDefaults(field);
    form.fields = form.fields.concat(field);
    removePlaceHolder();
    formStore.setForm(form);
  });

  subscribeComponent("field_clone", (params) => {
    const index = form.fields.findIndex((w) => w.id === params.field.id);
    const copy = fastClone(form.fields[index]);
    copy.name = copy.name + "-" + randomStringSmall();
    copy.label = copy.label + " Copy";
    copy.id = randomString();
    copy.selected = true;
    const temp = fastClone(form.fields);
    temp.splice(index + 1, 0, copy);
    form.fields = temp;
    formStore.set(copy);
  });

  subscribeComponent("save_form", async (params) => {
    await saveForm(params);
  });

  subscribeComponent("get_form_fields", () => {
    return form.fields;
  });

  subscribeComponent("drag_over", () => {
    removePlaceHolder();
  });

  subscribeComponent("drag_finished", async (elements) => {
    removePlaceHolder();

    let fields: IField[] = elements
      .filter((w: any) => w)
      .map((e: Element) => {
        if (e.id === "form-field-placeholder") {
          return undefined;
        }
        if (e.id.startsWith("form-field-")) {
          const field = form.fields.find(
            (w) => w.id === e.id.replace("form-field-", "")
          );
          if (field) {
            field.selected = false;
          }
          return field;
        }
        if (e.id.startsWith("sidebar-block-")) {
          const type = e.id.replace("sidebar-block-", "");
          let field: IField = {
            id: randomString(),
            type: type,
            name: "new-field-" + randomStringSmall(),
            label: "New Field " + randomStringSmall(),
            selected: true,
            value: undefined,
          };
          field = setFieldDefaults(field);
          return field;
        }
      });
    fields = fields.filter((w) => w != null);
    form.fields = fastClone(fields);
    await dispatch("destroy_dragula", {});
    await tick();
    dragForm = fastClone(form);
    formStore.setForm(form);
    await tick();
    dragForm = undefined;
    if (form.fields.length === 0) {
      addPlaceHolder();
    }
    await tick();
    dispatch("reload_dragula", {});
  });

  subscribeFieldChange(onMount, (newField: IField) => {
    if (!newField.selected) {
      return;
    }
    form.fields = form.fields.map((f) => {
      if (f.id !== newField.id && f.selected) {
        f.selected = false;
        formStore.set(f);
      }
      return f;
    });
  });

  subscribeComponent("form_saved", (f) => {
    dirty = false;
  });

  subscribeComponent("form_updated", (params) => {
    form = params;
  });

  subscribeComponent("unselect_field", () => {
    form.fields = form.fields.map((f) => {
      if (f.selected) {
        f.selected = false;
        formStore.set(f);
      }
      return f;
    });
  });

  subscribeFieldChange(onMount, async (field: IField) => {
    if (!form || !form.fields) {
      return;
    }
    dirty = true;
    const index = form.fields.findIndex((w) => w.id === field.id);
    if (index !== -1) {
      form.fields[index] = field;
    }
  });

  onMount(async () => {
    loadForm();
  });
</script>

<div>
  <ToastManager />
  {#if form == null || loadingActive}
    <div class="flex-column justify-content-center align-items-center">
      <div class="d-flex justify-content-center">
        <div
          class="spinner-border"
          style="width: 3rem; height: 3rem; margin-top: 2em"
          role="status">
          <span class="sr-only">Loading...</span>
        </div>
      </div>
    </div>
  {:else}
    <div>
      <div>
        {#if dragForm}
          <div>
            <DynamicForm form={dragForm} mode={DynamicFormMode.Preview} />
          </div>
        {:else}
          <div>
            <DynamicForm {form} mode={DynamicFormMode.Preview} />
          </div>
        {/if}
      </div>
    </div>
  {/if}
</div>
