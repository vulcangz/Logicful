<script lang="typescript">
  // @ts-nocheck
  import { dispatch, subscribeComponent } from "@app/event/EventBus";
  import { onMount, tick } from "svelte";
  import { debounce } from "@app/util/Debounce";
  import Button from "@app/components/Button.svelte";
  import { navigate } from "svelte-routing";
  import LogoFullAlternative from "@app/components/LogoFullAlternative.svelte";
  import { emptyUser, me } from "@app/services/AuthService";
  import { nameToInitials } from "@app/util/Nav";

  let saving = false;
  let saved = false;
  let loaded = false;
  let drake: any = null;
  let dragula: any;
  let shouldReload = true;
  let user = emptyUser;

  function defaultBlocks() {
    return [
      {
        name: "string",
        width: 256,
        height: 512,
        icon:
          "M256 52.048V12.065C256 5.496 250.726.148 244.158.066 211.621-.344 166.469.011 128 37.959 90.266.736 46.979-.114 11.913.114 5.318.157 0 5.519 0 12.114v39.645c0 6.687 5.458 12.078 12.145 11.998C38.111 63.447 96 67.243 96 112.182V224H60c-6.627 0-12 5.373-12 12v40c0 6.627 5.373 12 12 12h36v112c0 44.932-56.075 48.031-83.95 47.959C5.404 447.942 0 453.306 0 459.952v39.983c0 6.569 5.274 11.917 11.842 11.999 32.537.409 77.689.054 116.158-37.894 37.734 37.223 81.021 38.073 116.087 37.845 6.595-.043 11.913-5.405 11.913-12V460.24c0-6.687-5.458-12.078-12.145-11.998C217.889 448.553 160 444.939 160 400V288h36c6.627 0 12-5.373 12-12v-40c0-6.627-5.373-12-12-12h-36V112.182c0-44.932 56.075-48.213 83.95-48.142 6.646.018 12.05-5.346 12.05-11.992z",
        text: "Text Input",
      },
      {
        name: "combobox",
        width: 448,
        height: 512,
        icon:
          "M125.1 208h197.8c10.7 0 16.1 13 8.5 20.5l-98.9 98.3c-4.7 4.7-12.2 4.7-16.9 0l-98.9-98.3c-7.7-7.5-2.3-20.5 8.4-20.5zM448 80v352c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V80c0-26.5 21.5-48 48-48h352c26.5 0 48 21.5 48 48zm-48 346V86c0-3.3-2.7-6-6-6H54c-3.3 0-6 2.7-6 6v340c0 3.3 2.7 6 6 6h340c3.3 0 6-2.7 6-6z",
        text: "Dropdown",
      },
      {
        name: "switch",
        width: 576,
        height: 512,
        icon:
          "M384 64H192C85.961 64 0 149.961 0 256s85.961 192 192 192h192c106.039 0 192-85.961 192-192S490.039 64 384 64zM64 256c0-70.741 57.249-128 128-128 70.741 0 128 57.249 128 128 0 70.741-57.249 128-128 128-70.741 0-128-57.249-128-128zm320 128h-48.905c65.217-72.858 65.236-183.12 0-256H384c70.741 0 128 57.249 128 128 0 70.74-57.249 128-128 128z",
        text: "Toggle",
      },
      {
        name: "spacer",
        width: 512,
        height: 512,
        icon:
          "M505.12019,19.09375c-1.18945-5.53125-6.65819-11-12.207-12.1875C460.716,0,435.507,0,410.40747,0,307.17523,0,245.26909,55.20312,199.05238,128H94.83772c-16.34763.01562-35.55658,11.875-42.88664,26.48438L2.51562,253.29688A28.4,28.4,0,0,0,0,264a24.00867,24.00867,0,0,0,24.00582,24H127.81618l-22.47457,22.46875c-11.36521,11.36133-12.99607,32.25781,0,45.25L156.24582,406.625c11.15623,11.1875,32.15619,13.15625,45.27726,0l22.47457-22.46875V488a24.00867,24.00867,0,0,0,24.00581,24,28.55934,28.55934,0,0,0,10.707-2.51562l98.72834-49.39063c14.62888-7.29687,26.50776-26.5,26.50776-42.85937V312.79688c72.59753-46.3125,128.03493-108.40626,128.03493-211.09376C512.07526,76.5,512.07526,51.29688,505.12019,19.09375ZM384.04033,168A40,40,0,1,1,424.05,128,40.02322,40.02322,0,0,1,384.04033,168Z",
        text: "Spacer",
      },
      {
        name: "date",
        width: 448,
        height: 512,
        icon:
          "M0 464c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192H0v272zm320-196c0-6.6 5.4-12 12-12h40c6.6 0 12 5.4 12 12v40c0 6.6-5.4 12-12 12h-40c-6.6 0-12-5.4-12-12v-40zm0 128c0-6.6 5.4-12 12-12h40c6.6 0 12 5.4 12 12v40c0 6.6-5.4 12-12 12h-40c-6.6 0-12-5.4-12-12v-40zM192 268c0-6.6 5.4-12 12-12h40c6.6 0 12 5.4 12 12v40c0 6.6-5.4 12-12 12h-40c-6.6 0-12-5.4-12-12v-40zm0 128c0-6.6 5.4-12 12-12h40c6.6 0 12 5.4 12 12v40c0 6.6-5.4 12-12 12h-40c-6.6 0-12-5.4-12-12v-40zM64 268c0-6.6 5.4-12 12-12h40c6.6 0 12 5.4 12 12v40c0 6.6-5.4 12-12 12H76c-6.6 0-12-5.4-12-12v-40zm0 128c0-6.6 5.4-12 12-12h40c6.6 0 12 5.4 12 12v40c0 6.6-5.4 12-12 12H76c-6.6 0-12-5.4-12-12v-40zM400 64h-48V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H160V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H48C21.5 64 0 85.5 0 112v48h448v-48c0-26.5-21.5-48-48-48z",
        text: "Date",
      },
      {
        name: "block",
        width: 448,
        height: 512,
        icon:
          "M27.31 363.3l96-96a16 16 0 0 0 0-22.62l-96-96C17.27 138.66 0 145.78 0 160v192c0 14.31 17.33 21.3 27.31 11.3zM432 416H16a16 16 0 0 0-16 16v32a16 16 0 0 0 16 16h416a16 16 0 0 0 16-16v-32a16 16 0 0 0-16-16zm3.17-128H204.83A12.82 12.82 0 0 0 192 300.83v38.34A12.82 12.82 0 0 0 204.83 352h230.34A12.82 12.82 0 0 0 448 339.17v-38.34A12.82 12.82 0 0 0 435.17 288zm0-128H204.83A12.82 12.82 0 0 0 192 172.83v38.34A12.82 12.82 0 0 0 204.83 224h230.34A12.82 12.82 0 0 0 448 211.17v-38.34A12.82 12.82 0 0 0 435.17 160zM432 32H16A16 16 0 0 0 0 48v32a16 16 0 0 0 16 16h416a16 16 0 0 0 16-16V48a16 16 0 0 0-16-16z",
        text: "Content Block",
      },
      {
        name: "file",
        width: 384,
        height: 512,
        icon:
          "M224 136V0H24C10.7 0 0 10.7 0 24v464c0 13.3 10.7 24 24 24h336c13.3 0 24-10.7 24-24V160H248c-13.2 0-24-10.8-24-24zm65.18 216.01H224v80c0 8.84-7.16 16-16 16h-32c-8.84 0-16-7.16-16-16v-80H94.82c-14.28 0-21.41-17.29-11.27-27.36l96.42-95.7c6.65-6.61 17.39-6.61 24.04 0l96.42 95.7c10.15 10.07 3.03 27.36-11.25 27.36zM377 105L279.1 7c-4.5-4.5-10.6-7-17-7H256v128h128v-6.1c0-6.3-2.5-12.4-7-16.9z",
        text: "File Input",
      },
      {
        name: "address",
        width: 576,
        height: 512,
        icon:
          "M528 32H48C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h480c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48zm-352 96c35.3 0 64 28.7 64 64s-28.7 64-64 64-64-28.7-64-64 28.7-64 64-64zm112 236.8c0 10.6-10 19.2-22.4 19.2H86.4C74 384 64 375.4 64 364.8v-19.2c0-31.8 30.1-57.6 67.2-57.6h5c12.3 5.1 25.7 8 39.8 8s27.6-2.9 39.8-8h5c37.1 0 67.2 25.8 67.2 57.6v19.2zM512 312c0 4.4-3.6 8-8 8H360c-4.4 0-8-3.6-8-8v-16c0-4.4 3.6-8 8-8h144c4.4 0 8 3.6 8 8v16zm0-64c0 4.4-3.6 8-8 8H360c-4.4 0-8-3.6-8-8v-16c0-4.4 3.6-8 8-8h144c4.4 0 8 3.6 8 8v16zm0-64c0 4.4-3.6 8-8 8H360c-4.4 0-8-3.6-8-8v-16c0-4.4 3.6-8 8-8h144c4.4 0 8 3.6 8 8v16z",
        text: "Address",
      },
      {
        name: "checkbox-group",
        width: 448,
        height: 512,
        icon:
          "M400 480H48c-26.51 0-48-21.49-48-48V80c0-26.51 21.49-48 48-48h352c26.51 0 48 21.49 48 48v352c0 26.51-21.49 48-48 48zm-204.686-98.059l184-184c6.248-6.248 6.248-16.379 0-22.627l-22.627-22.627c-6.248-6.248-16.379-6.249-22.628 0L184 302.745l-70.059-70.059c-6.248-6.248-16.379-6.248-22.628 0l-22.627 22.627c-6.248 6.248-6.248 16.379 0 22.627l104 104c6.249 6.25 16.379 6.25 22.628.001z",
        text: "Checkboxes",
      },
      {
        name: "radio-group",
        width: 512,
        height: 512,
        icon:
          "M256 8C119.033 8 8 119.033 8 256s111.033 248 248 248 248-111.033 248-248S392.967 8 256 8zm80 248c0 44.112-35.888 80-80 80s-80-35.888-80-80 35.888-80 80-80 80 35.888 80 80z",
        text: "Radio Buttons",
      },
      {
        name: "full-name",
        width: 448,
        height: 512,
        icon:
          "M436 160c6.6 0 12-5.4 12-12v-40c0-6.6-5.4-12-12-12h-20V48c0-26.5-21.5-48-48-48H48C21.5 0 0 21.5 0 48v416c0 26.5 21.5 48 48 48h320c26.5 0 48-21.5 48-48v-48h20c6.6 0 12-5.4 12-12v-40c0-6.6-5.4-12-12-12h-20v-64h20c6.6 0 12-5.4 12-12v-40c0-6.6-5.4-12-12-12h-20v-64h20zm-228-32c35.3 0 64 28.7 64 64s-28.7 64-64 64-64-28.7-64-64 28.7-64 64-64zm112 236.8c0 10.6-10 19.2-22.4 19.2H118.4C106 384 96 375.4 96 364.8v-19.2c0-31.8 30.1-57.6 67.2-57.6h5c12.3 5.1 25.7 8 39.8 8s27.6-2.9 39.8-8h5c37.1 0 67.2 25.8 67.2 57.6v19.2z",
        text: "Full Name",
      },
      {
        name: "section-header",
        width: 512,
        height: 512,
        icon:
          "M448 96v320h32a16 16 0 0 1 16 16v32a16 16 0 0 1-16 16H320a16 16 0 0 1-16-16v-32a16 16 0 0 1 16-16h32V288H160v128h32a16 16 0 0 1 16 16v32a16 16 0 0 1-16 16H32a16 16 0 0 1-16-16v-32a16 16 0 0 1 16-16h32V96H32a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16v32a16 16 0 0 1-16 16h-32v128h192V96h-32a16 16 0 0 1-16-16V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16v32a16 16 0 0 1-16 16z",
        text: "Section Header",
      },
    ];
  }

  let blocks = defaultBlocks();

  const loadDragula = debounce(async () => {
    shouldReload = true;
    if (!dragula) {
      dragula = (await import("dragula")).default;
      await tick();
    }
    if (drake) {
      drake.destroy();
    }
    await tick();
    drake = dragula(
      [
        document.querySelector("#block-container"),
        document.querySelector("#form-preview-fields"),
      ],
      {
        copy: function (el, source) {
          return source === document.getElementById("block-container");
        },
        accepts: function (el, target) {
          return target !== document.getElementById("block-container");
        },
      }
    )
      .on("drag", function (el) {
        if (el.id && el.id.startsWith("form-field-")) {
          return;
        }
        const container = document.getElementById("form-preview-fields");
        if (container && !container.className?.includes("ex-over")) {
          container.className += " ex-over";
        }
      })
      .on("over", function (el, container) {
        if (el.id && el.id.startsWith("form-field-")) {
          return;
        }
        if (
          container.id === "form-preview-fields" &&
          !container.className?.includes("ex-over")
        ) {
          container.className += " ex-over";
        }
        dispatch("drag_over", container);
      })
      .on("drop", function (el) {
        const container = document.getElementById("form-preview-fields");
        if (container) {
          container.className = container.className.replace("ex-over", "");
        }
        const fields = Array.from(
          document.querySelector("#form-preview-fields").childNodes
        ).filter(
          (w) =>
            w.id?.startsWith("sidebar-block") || w.id?.startsWith("form-field-")
        );
        el.remove();
        dispatch("drag_finished", fields);
      });
  }, 300);

  function addField(block: any) {
    dispatch("add_field", {
      type: block.name,
    });
  }

  async function saveDraft() {
    saving = true;
    await dispatch("save_form", {
      status: "draft",
      initiator: "user",
    });
    saving = false;
  }

  subscribeComponent("form_saved", (data) => {
    const initiator = data.options?.initiator;
    saved = false;
    if (initiator === "user") {
      dispatch("show_toast", {
        message: "Successfully saved your form.",
      });
    }
  });

  subscribeComponent("reload_dragula", () => {
    loadDragula();
  });

  subscribeComponent("destroy_dragula", () => {
    shouldReload = false;
    requestAnimationFrame(() => {
      if (drake) {
        drake.destroy();
      }
    });
  });

  subscribeComponent("form_updated", () => {
    if (shouldReload) {
      loadDragula();
    }
  });

  subscribeComponent("form_loaded", () => {
    loaded = true;
    loadDragula();
  });

  function onLogoClick() {
    if (
      !window.confirm(
        "Are you sure you want to leave the builder? Any unsaved changes will be lost."
      )
    ) {
      return;
    }
    navigate("/folder");
  }

  onMount(async () => {
    user = await me();
    import("dragula/dist/dragula.css");
  });
</script>

<div class="flex md:flex-shrink-0">
  <div class="flex flex-col w-52">
    <div class="flex flex-col h-0 flex-1 bg-indigo-800">
      <div class="flex-1 flex flex-col pt-5 pb-4 overflow-y-auto">
        <div class="h-12 ml-8 cursor-pointer" on:click={onLogoClick}>
          <a href="/folder">
            <LogoFullAlternative />
          </a>
        </div>
        <nav class="flex-1 px-2 bg-indigo-800 space-y-1">
          <div
            class="flex flex-col space-y-3 sm:space-y-0 sm:space-x-3 sm:flex-row
              xl:flex-col xl:space-x-0 xl:space-y-3">
            <span class="inline-flex rounded-md shadow-sm w-full">
              {#if !loaded}
                <button
                  class="w-full inline-flex items-center justify-center px-4
                    py-2 border border-transparent text-sm leading-5 font-medium
                    rounded-md text-white bg-indigo-600 hover:bg-indigo-500
                    focus:outline-none focus:border-indigo-700
                    focus:ring-indigo active:bg-indigo-700 transition
                    ease-in-out duration-150"
                  type="button"
                  disabled>Loading...</button>
              {:else if saved}
                <button
                  class="w-full inline-flex items-center justify-center px-4
                    py-2 border border-transparent text-sm leading-5 font-medium
                    rounded-md text-white bg-indigo-600 hover:bg-indigo-500
                    focus:outline-none focus:border-indigo-700
                    focus:ring-indigo active:bg-indigo-700 transition
                    ease-in-out duration-150"
                  type="button"
                  disabled>Saved</button>
              {:else if saving}
                <button
                  class="w-full inline-flex items-center justify-center px-4
                    py-2 border border-transparent text-sm leading-5 font-medium
                    rounded-md text-white bg-indigo-600 hover:bg-indigo-500
                    focus:outline-none focus:border-indigo-700
                    focus:ring-indigo active:bg-indigo-700 transition
                    ease-in-out duration-150"
                  type="button"
                  disabled>Saving...</button>
              {:else}
                <Button type="primary" width="full" onClick={saveDraft}>
                  Save Form
                </Button>
              {/if}
            </span>
          </div>
          <div id="block-container">
            {#each blocks as block}
              <div
                draggable={true}
                on:click={() => addField(block)}
                id={'sidebar-block-' + block.name}>
                <a
                  href="#"
                  class="group flex items-center px-2 py-2 text-sm leading-5
                    font-medium text-indigo-300 rounded-md hover:text-white
                    hover:bg-indigo-700 focus:outline-none focus:text-white
                    focus:bg-indigo-700 transition ease-in-out duration-150">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 {block.width} {block.height}"
                    class="h-7 w-7 pr-2 fill-current text-indigo-400"><path
                      d={block.icon} /></svg>
                  {block.text}
                </a>
              </div>
            {/each}
          </div>
        </nav>
      </div>
      <div class="flex-shrink-0 flex border-t border-indigo-700 p-4">
        <a
          href="/account/profile"
          target="_blank"
          class="flex-shrink-0 w-full group block">
          <div class="flex items-center">
            <div>
              <span
                class="inline-flex items-center justify-center h-10 w-10
                  rounded-full bg-indigo-500 cursor-pointer">
                <span
                  class="font-medium leading-none text-white">{nameToInitials(user.fullName) || 'DU'}</span>
              </span>
            </div>
            <div class="ml-3">
              <p class="text-sm leading-5 font-medium text-white">
                {user.displayName || 'Demo User'}
              </p>
              <a
                href="/account/profile"
                target="_blank"
                class="text-xs leading-4 font-medium text-indigo-300
                  group-hover:text-indigo-100 transition ease-in-out
                  duration-150">
                View profile
              </a>
            </div>
          </div>
        </a>
      </div>
    </div>
  </div>
</div>
