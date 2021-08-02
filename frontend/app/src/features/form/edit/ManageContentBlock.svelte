<script lang="typescript">
  import { onMount } from "svelte";
  import type { ContentBlock } from "@app/models/ContentBlock";
  import { getUrlParameter } from "@app/util/Http";
  import DropdownButton from "@app/components/DropdownButton.svelte";
  import ContentBlockList from "./ContentBlockList.svelte";
  import { dispatch } from "@app/event/EventBus";
  import { randomString } from "@app/util/Generate";
  import { getApi, postApi, putApi } from "@app/services/ApiService";
  import ConfigField from "./ConfigField.svelte";

  let block: ContentBlock;
  let data: any;
  let loading = false;
  let errored = false;
  export let id: string;
  export let isNew: boolean;

  onMount(async () => {
    if (!isNew) {
      await load();
    } else {
      block = {};
    }
  });

  async function load() {
    loading = true;
    const blocks: ContentBlock[] = await getApi("content-block");
    const result = blocks.find((w) => w.id === id);
    if (!result) {
      return;
    }
    data = await convertUrlToLocal(result);
    block = result;
    loading = false;
  }

  async function convertUrlToLocal(block: ContentBlock): Promise<any> {
    loading = true;
    try {
      const url = block.value;
      if (!url) {
        return [];
      }
      const response = await fetch(url);
      return await response.json();
    } catch (ex) {
      errored = true;
      return [];
    } finally {
      loading = false;
    }
  }

  async function save() {
    block.value = await generateInlineUrl();
    await postApi("content-block", block);
    dispatch("option_set_modified", block);
    dispatch("dialog_show", {
      child: ContentBlockList,
      closeOnOutsideClick: false,
      confirmCloseOnDirty: true,
      title: "Manage Content Blocks",
    });
  }

  async function generateInlineUrl(): Promise<string> {
    const saveId = getUrlParameter("id", block.value);
    const qs = saveId ? `?id=${saveId}` : "";
    const { message } = qs
      ? await putApi<any>(`s3/json?${qs}`, data)
      : await postApi<any>(`s3/json?${qs}`, data);
    return message;
  }

  function onChange(blockData: any) {
    data = blockData;
  }
</script>

<div>
  {#if block == null}
    <div style="text-align: center; padding-top: 1em; padding-bottom: 1em;">
      <div class="spinner-border text-secondary" role="status">
        <span class="sr-only">Loading...</span>
      </div>
    </div>
  {:else}
    <div style="margin-top: 1em">
      <ConfigField
        field={{ id: `${block.id}-name`, type: 'string', required: true, name: 'name', label: 'Name', placeholder: 'Name', value: block.name, onChange: (value) => {
            block.name = value;
          } }} />
      <ConfigField
        config={{ onChange }}
        field={{ value: data, id: randomString(), name: block.name, type: 'block-editor' }} />
    </div>
  {/if}
  <div class="float-right">
    <DropdownButton
      label={'Save'}
      processingLabel={'Saving...'}
      actions={[{ label: 'Save as Draft', onClick: save }, { label: 'Save and Publish', onClick: save }]} />
  </div>
</div>
