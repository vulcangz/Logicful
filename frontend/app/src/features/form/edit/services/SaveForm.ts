import type { IForm } from "@app/models/IForm";
import { postApi, putApi } from "@app/services/ApiService";
import formStore from "@app/store/FormStore";
import { fastClone } from "@app/util/Compare";
import { dispatch } from "@app/event/EventBus";
import {getUrlParameter} from "@app/util/Http";

export async function saveForm(
  options: {
    initiator: string;
  } = { initiator: "" },
  form?: IForm
) {
  if (!form) {
    form = formStore.getForm();
  }
  if (form.id === "demo" || getUrlParameter("cypress_test")) {
    setTimeout(() => {
      dispatch("form_saved", {
        form,
        options,
      });
    }, 1000);
    return;
  }
  const clone = fastClone(form);
  removeValues(clone);
  await save(clone);
  saveToLocalStorage(clone);
  dispatch("form_saved", {
    form,
    options,
  });
}

export function saveToLocalStorage(form: IForm) {
  let copy = fastClone(form);
  copy = removeValues(copy);
  localStorage.setItem("form", JSON.stringify(copy));
}

const skipRemoveValue = ["block"];

function removeValues(form: IForm): IForm {
  if (!form.fields) {
    return form;
  }
  form.fields = form.fields.map((f) => {
    if (skipRemoveValue.includes(f.type)) {
      return f;
    }
    delete f.value;
    delete f.error;
    return f;
  });
  return form;
}

async function save(form: IForm): Promise<IForm> {
  return await (form.id
    ? putApi(`form/${form.id}`, form)
    : postApi("form", form));
}
