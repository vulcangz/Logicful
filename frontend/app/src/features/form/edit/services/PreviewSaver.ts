import { dispatch, subscribe } from "@app/event/EventBus";
import type { IForm } from "@app/models/IForm";
import formStore from "@app/store/FormStore";
import { fastClone, fastEquals } from "@app/util/Compare";
import { debounce } from "@app/util/Debounce";
import { saveToLocalStorage } from "./SaveForm";

let initialized = false;
let lastForm: IForm | undefined = undefined;

const debounceSaveLocal = debounce(() => {
  const form = formStore.getForm();
  saveToLocalStorage(form);
}, 300);

const debounceSave = debounce(() => {
  const form = getForm();
  if (fastEquals(form, lastForm)) {
    return;
  }
  lastForm = form;
  dispatch("save_form", {
    status: "draft",
    initiator: "",
  });
}, 800);

function getForm(): IForm {
  const form: IForm = fastClone(formStore.getForm());
  form.fields.forEach((f) => {
    delete f.selected;
    delete f.updated;
  });
  return form;
}

export function startPreviewSaver() {
  if (initialized) {
    return;
  }
  initialized = true;
  subscribe("field_changed", () => {
    debounceSaveLocal();
    debounceSave();
  });
  subscribe("form_updated", () => {
    if (!lastForm) {
      lastForm = getForm();
    }
    debounceSaveLocal();
    debounceSave();
  });
  subscribe("form_loaded", () => {
    debounceSaveLocal();
  });
}

export function stopPreviewSaver() {}
