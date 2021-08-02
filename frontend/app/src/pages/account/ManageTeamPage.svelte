<script lang="typescript">
  import AccountSidebar from "../../components/account/AccountSidebar.svelte";
  import ManageTeam from "../../components/account/ManageTeam.svelte";
</script>

<div>
  <div class="section section-lg pt-5 pt-md-6 bg-soft">
    <div class="container">
      <div class="row pt-5 pt-md-0">
        <div class="col-12 col-md-4 d-none d-lg-block">
          <AccountSidebar page={'security'} />
        </div>
        <div class="col-12 col-lg-8">
          <ManageTeam />
        </div>
      </div>
    </div>
  </div>
</div>
