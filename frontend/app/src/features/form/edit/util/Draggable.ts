export function transformDraggedElement(el: any, data: any, index: number) {
  if (data.name === "string") {
    el.innerHTML = "<label>New Text Input</label><input class='form-control'/>";
  }
  if (data.name === "combobox") {
    el.innerHTML =
      "<label>New Dropdown</label><select class='form-control'><option>Dropdown Value</option></select>";
  }
  if (data.name === "switch") {
    el.innerHTML = `<div class="form-check form-switch">
            <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
            <label class="form-check-label" for="flexSwitchCheckDefault">New Toggle</label>
            </div>`;
  }
  if (data.name === "date") {
    el.innerHTML =
      "<label>New Date</label><input type='date' class='form-control shadow'></input>";
  }
  if (data.name === "block") {
    el.innerHTML =
      "<div style='background-color: #f0f0f0;margin-top:5px;padding: 10px 10px 3px;border-radius: 0.45em'><h5>New Content Block</h5><p>Content will display here.</p></div>";
  }
  if (data.name === "file") {
    el.innerHTML = `<label>New File</label><div class="form-file">
        <input type="file" class="form-file-input" />
        <label class="form-file-label">
          <span class="form-file-text">Choose a file...</span>
          <span class="form-file-button">Browse</span>
        </label>
      </div>`;
  }
  if (data.name === "checkbox-group") {
    el.innerHTML = `<label>New Checkbox Group</label><div class="form-check">
    <input class="form-check-input" type="checkbox" checked="true" value="" id="checkboxOne" />
    <label class="form-check-label" for="defaultCheck1">Option 1</label>
  </div>
  <div class="form-check">
    <input class="form-check-input" type="checkbox" value="" id="checkboxTwo" />
    <label class="form-check-label" for="defaultCheck2">Option 2</label>
  </div>`;
  }
  if (data.name === "radio-group") {
    el.innerHTML = `<label>New Radio Group</label><div class="form-check">
    <input class="form-check-input" type="radio" value="" checked="true" id="checkboxOne" />
    <label class="form-check-label" for="defaultCheck1">Option One</label>
  </div>
  <div class="form-check">
    <input class="form-check-input" type="radio" value="" id="checkboxTwo" />
    <label class="form-check-label" for="defaultCheck2">Option Two</label>
  </div>`;
  }
  if (data.name === "address") {
    el.innerHTML = `<div class="row">
    <div class="col">
      <label class="form-check-label" for="defaultCheck1">Address</label>
      <div class="row">
        <input class="form-control" />
        <small class="form-text text-muted">Address Line 1</small>
      </div>
      <div class="row" style="padding-top: 0.8em;">
        <input class="form-control" />
        <small class="form-text text-muted">Address Line 2</small>
      </div>
      <div class="row" style="padding-top: 0.8em;">
        <div class="col-5">
          <input class="form-control" />
          <small class="form-text text-muted">City</small>
        </div>
        <div class="col-3">
          <select class="form-control shadow">
            <option>Select State</option>
          </select>
        </div>
        <div class="col-4">
          <input class="form-control" />
          <small class="form-text text-muted">Zip Code</small>
        </div>
      </div>
    </div>
  </div>`;
  }
  return el;
}
