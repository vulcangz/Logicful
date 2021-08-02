<script lang="typescript">
  import Dialog from "@app/components/layout/Dialog.svelte";
  import type { ButtonAction } from "@app/components/models/ComponentProps";
  import Repeater from "@app/components/Repeater.svelte";
  import { isEmail } from "@app/guards/Guard";
  import type { Team } from "@app/models/IForm";
  import { postApi } from "@app/services/ApiService";

  export let team: Team;
  export let onClose: () => any;
  let emails: string[] = [];

  function onChange(values: any[]) {
    emails = values;
  }

  async function add() {
    const error = emails.find((e) => !isEmail(e));
    if (error) {
      throw new Error(
        "All inputs must be a valid email. " + error + " is not valid."
      );
    }
    await postApi(`team/members/invite`, emails);
  }

  let actions: ButtonAction[] = [
    {
      label: "Send Invite",
      type: "primary",
      focus: true,
      onClick: add,
    },
    {
      label: "Cancel",
      type: "secondary",
    },
  ];
</script>

<Dialog isOpen={true} {actions} title={`Add New Teammate(s)`} {onClose}>
  <Repeater
    inputType="email"
    onlyLabel={true}
    {onChange}
    label={'Email Addresses'}
    placeholder={'user@example.com'} />
</Dialog>
