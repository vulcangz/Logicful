<script lang="typescript">
    import type {IField} from "@app/models/IField";
    import {subscribeFieldChange} from "@app/event/FieldEvent";
    import Label from "./Label.svelte";
    import {onMount} from "svelte";
    import formStore from "@app/store/FormStore";
    import {debounce} from "@app/util/Debounce";
    import InputErrorIcon from "./validations/InputErrorIcon.svelte";
    import InputErrorText from "./validations/InputErrorText.svelte";
    import FieldHelperText from "./FieldHelperText.svelte";
    import {validateField} from "@app/services/FieldValidator";

    export let field: IField;
    export let value = "";
    export let type = "text";
    let debouncedOnChange: any;

    subscribeFieldChange(onMount, (newField) => {
        if (newField.id === field.id) {
            value = newField.value ?? value ?? "";
        }
    });

    onMount(() => {
        debouncedOnChange = debounce((e: any) => {
            field.value = e.target.value ?? "";
            formStore.set(field, {
                fromUser: true,
                field: "value",
                value: field.value,
            });
            validateField(field, field.value);
            field.onChange?.(field.value);
        }, 500);

        value =
            field.value ?? formStore.getValue(field.configTarget ?? field.id) ?? "";
    });
</script>

<div>
    <Label {field}/>
    <div class="mt-1 relative rounded-md shadow-sm">
        {#if field.rows && field.rows > 1}
      <textarea
              rows={field.rows}
              data-test-id={"text-input-" +
          field.label +
          "-" +
          (field.required ? "required" : "optional")}
              on:click|stopPropagation
              on:input={debouncedOnChange}
              class:input-error={field.error}
              class={"shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"}
              id={field.id}
              {value}
              placeholder={field.placeholder ?? ""}
              name={field.name}
              {type}
      />
        {:else}
            <input
                    on:click|stopPropagation
                    on:input={debouncedOnChange}
                    data-test-id={"text-input-" +
          field.label +
          "-" +
          (field.required ? "required" : "optional")}
                    class:input-error={field.error}
                    class={"shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"}
                    id={field.id}
                    {value}
                    placeholder={field.placeholder ?? ""}
                    name={field.name}
                    {type}
            />
        {/if}
        <InputErrorIcon {field}/>
    </div>
    <InputErrorText {field}/>
    <FieldHelperText {field}/>
</div>
