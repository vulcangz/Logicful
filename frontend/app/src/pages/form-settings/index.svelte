<script lang="typescript">
    import {onMount} from "svelte";
    import type {IForm} from "@app/models/IForm";
    import formStore from "@app/store/FormStore";
    import FormSettings from "@app/components/form_settings/FormSettings.svelte";
    import {getApi} from "@app/services/ApiService";
    import {getUrlParameter} from "@app/util/Http";

    export let form: IForm;
    let selected: string = "general";

    onMount(async () => {
        const formId = getUrlParameter("formId");
        form = await getApi(`form/${formId}`);
        form.id = formId;
        formStore.setForm(form);
    });
</script>

<div>
    <FormSettings {form} {selected}/>
</div>
