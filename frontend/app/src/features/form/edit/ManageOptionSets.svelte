<script lang="typescript">
  import {onMount} from "svelte";
  import type {OptionSet} from "@app/models/OptionSet";
  import Repeater from "@app/components/Repeater.svelte";
  import type {LabelValue} from "@app/models/IField";
  import {dispatch} from "@app/event/EventBus";
  import {isString} from "@app/guards/Guard";
  import {getApi, postApi, putApi} from "@app/services/ApiService";
  import ConfigField from "./ConfigField.svelte";
  import Button from "@app/components/Button.svelte";
  import Loader from "../../../components/Loader.svelte";

  let sets: OptionSet[] = [];
  let loading = false;
  let errored = false;
  export let editingId: string;
  export let isNew: boolean;
  export let onSave: () => any;

  onMount(async () => {
    if (isNew) {
      sets = sets.concat([
        {
          value: [
            {
              label: "",
              value: "",
            },
          ],
          type: "local",
        },
      ]);
    } else {
      await load();
    }
  });

  async function load() {
    loading = true;
    const data: OptionSet[] = await getApi("option-set");
    const result = data.find((w) => w.id === editingId);
    if (!result) {
      return;
    }
    if (result.type === "local") {
      result.localSaveId = result.value as string;
      result.value = await convertUrlToLocal(result);
    }
    sets = [result];
    loading = false;
  }

  async function loadLocalOptions(index: number) {
    console.log("oload local", index);
    sets[index].value = await convertUrlToLocal(sets[index]);
  }

  async function convertUrlToLocal(set: OptionSet): Promise<LabelValue[]> {
    loading = true;
    try {
      const url = (set.value as string) ?? set.localSaveId;
      if (!url) {
        return [
          {
            label: "",
            value: "",
          },
        ];
      }
      const response = await fetch(url);
      const data = await response.json();
      const results: LabelValue[] = [];
      Object.keys(data).forEach((key) => {
        results.push({label: key, value: data[key]});
      });
      return results;
    } catch (ex) {
      errored = true;
      return [];
    } finally {
      loading = false;
    }
  }

  function onRepeaterChange(data: LabelValue[] | string[], index: number) {
    sets[index].value = data as LabelValue[];
  }

  async function save() {
    const promises = sets.map(async (s) => {
      if (s.type === "local") {
        s.value = await generateInlineUrl(s);
      }
      return s;
    });
    const toSave = await Promise.all(promises);
    editingId
            ? await putApi("option-set/" + editingId, toSave[0])
            : await postApi("option-set", toSave[0]);
    dispatch("option_set_modified", toSave[0]);
    onSave();
  }

  async function generateInlineUrl(set: OptionSet): Promise<string> {
    const body: any = {};
    const v = set.value as LabelValue[];
    v.forEach((s) => {
      body[s.label] = s.value;
    });

    const saveId =
            set.localSaveId?.replace("https://", "")?.split("/")[1] ?? "";
    const qs = saveId ? `?id=${saveId}` : "";
    const {message} = qs
            ? await putApi<any>(`s3/json${qs}`, body)
            : await postApi<any>(`s3/json${qs}`, body);
    return message;
  }
</script>

<div>
  {#if sets.length > 0}

  {:else}
    <div style="text-align: center; padding-top: 1em; padding-bottom: 1em;">
      <div class="spinner-border text-secondary" role="status">
        <span class="sr-only">Loading...</span>
      </div>
    </div>
  {/if}

  {#if Array.isArray(sets)}
    {#each sets as set, index}
      <div>
        <h2 class="ml-3">{set.name ?? ''}</h2>
        <div id={set.name ?? ''}>
          <ConfigField
                  field={{ isConfigField: true, id: `${set.id}-name`, type: 'string', required: true, name: 'name', label: 'Name', placeholder: 'Name', value: set.name, onChange: (value) => {
              set.name = value;
            } }} />
          <ConfigField
                  field={{ onChange: (value) => {
              console.log("on change", value);
              if (value === 'local') {
                set.remoteUrl = set.value ?? [];
                set.value = set.localOptions ?? [];
                if (!isNew && set.localOptions?.length === 0) {
                  set.value = undefined;
                  loadLocalOptions(index);
                }
              }
              if (value === 'remote') {
                set.localOptions = set.value ?? [];
                set.value = set.remoteUrl;
              }
              set.type = value;
            }, id: `${set.id}-type`, type: 'combobox', required: true, value: set.type, options: { type: 'local', value: [{ label: 'Inline', value: 'local' }, { label: 'Remote', value: 'remote' }] }, name: 'type', label: 'Type', helperText: 'Choose whether you want to automatically load options in from a remote url or manually specify them here.', isConfigField: true }} />

          {#if set.type === 'remote'}
            <ConfigField
                    field={{ isConfigField: true, helperText: 'See <a href="test" target="_blank">Remote Option Set Guide</a> for information on how to structure your endpoint response.', onChange: (value) => {
                set.value = value;
              }, id: `${set.name}-url`, type: 'string', value: set.value, name: 'url', label: 'Url', required: true }} />
          {:else if loading}
            <Loader/>
          {:else if errored}
            Failed to load, please try re-opening this dialog.
          {:else}
            <div class="ml-3">
              <Repeater
                      options={isString(set.value) ? [] : set.value}
                      onChange={(data) => {
                onRepeaterChange(data, index);
              }} />
            </div>
          {/if}
        </div>
      </div>
    {/each}
    {/if}

  <div class="ml-3 mt-3">
    <Button type="primary" onClick={save}>Save Option Set</Button>
  </div>
</div>
