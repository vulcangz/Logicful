<script lang="typescript">
  import type { IField } from "@app/models/IField";
  import { subscribeFieldChange } from "@app/event/FieldEvent";
  import { onMount } from "svelte";

  import formStore from "@app/store/FormStore";
  let element: any;

  export let field: IField;
  export let value = { blocks: [] };
  export let onChange: (value: any) => any;
  let editor: any;

  function onFieldChange(data: any) {
    onChange?.(data);
  }

  subscribeFieldChange(onMount, (newField) => {
    if (newField.id === field.id) {
      value = newField.value ?? "";
    }
  });

  onMount(async () => {
    //@ts-ignore
    import("quill/dist/quill.js");
    //@ts-ignore
    import("quill/dist/quill.snow.css");
    const Quill = (await import("quill")).default;

    value = formStore.getValue(field.configTarget ?? field.id) ?? "";

    var toolbarOptions = [
      ["bold", "italic", "underline", "strike"],
      [{ list: "ordered" }, { list: "bullet" }],
      [{ header: [1, 2, 3, 4, 5, 6, false] }, { color: [] }, { align: [] }],
      ["clean"],
    ];
    let quill = new Quill(element, {
      theme: "snow",
      placeholder: "Start typing and see the preview on the left side.",
      modules: {
        toolbar: toolbarOptions,
      },
    });

    //@ts-ignore
    quill.container.firstChild.innerHTML = value;

    quill.on("text-change", function (delta: any, oldDelta: any, source: any) {
      //@ts-ignore
      field.value = quill.container.firstChild.innerHTML;
      formStore.set(field, {
        fromUser: true,
        value: field.value,
        field: "value",
      });
    });
  });
</script>

<div>
  <div bind:this={element} />
</div>
