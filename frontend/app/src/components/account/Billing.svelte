<div class="row mb-5">
  <div class="col-12 mb-4">
    <div class="card border-light text-center py-4 mb-4">
      <div class="card-body">
        <h2 class="display-3 mb-3">
          Pixel<span class="pixel-pro-badge subscription-badge bg-tertiary
              font-weight-bolder text-white">PRO</span>
        </h2>
        <p class="my-4">
          Next payment of <span class="font-weight-bold">$36 (yearly)</span> occurs
          on August 13, 2020.
        </p>
        <a class="btn btn-sm btn-dark" href="#">Cancel subscription</a>
      </div>
    </div>
    <form action="#" method="post" class="card border-light p-3 mb-4">
      <div class="card-header bg-white border-light p-3 mb-4 mb-md-0">
        <h3 class="h5 mb-0">Billing details</h3>
      </div>
      <div class="card-body p-0 p-md-4">
        <div class="row justify-content-center">
          <div class="col-12 col-lg-6">
            <!-- Form -->
            <div class="mb-4">
              <label for="cartInputAddress1">Address</label>
              <input
                type="text"
                placeholder="123 4th St"
                class="form-control"
                id="cartInputAddress1" />
            </div>
            <!-- End of Form -->
          </div>
          <div class="col-12 col-lg-6">
            <!-- Form -->
            <div class="mb-4">
              <label for="cartInputCity1">City</label>
              <input
                type="text"
                placeholder="Dallas"
                class="form-control"
                id="cartInputCity1" />
            </div>
            <!-- End of Form -->
          </div>
          <div class="col-12 col-lg-6">
            <!-- Form -->
            <div class="mb-4">
              <label for="cartInputZip1">Zip/Postal code</label>
              <input
                type="number"
                placeholder="123456"
                class="form-control"
                id="cartInputZip1" />
            </div>
            <!-- End of Form -->
          </div>
          <div class="col-12 col-lg-6">
            <!-- Form -->
            <div class="mb-4">
              <label for="cartInputCompany1">Company Name</label>
              <input
                type="text"
                placeholder="Company LLC"
                class="form-control"
                id="cartInputCompany1"
                required />
            </div>
            <!-- End of Form -->
          </div>
          <div class="col-12 col-lg-6 mb-4">
            <!--Form-->
            <label
              class="my-1 mr-2"
              for="inlineFormCustomSelectBilling">Country</label>
            <select
              class="form-select"
              id="inlineFormCustomSelectBilling"
              aria-label="Messages select example">
              <option selected>Choose...</option>
              <option value="1">United States</option>
              <option value="2">Germany</option>
              <option value="3">Canada</option>
            </select>
            <!-- End of Form -->
          </div>
          <div class="col-12 col-lg-6">
            <!-- Form -->
            <div class="mb-4">
              <label for="cartInputPhone1">Phone Number</label>
              <input
                type="number"
                placeholder="+(12)345 6789"
                class="form-control"
                id="cartInputPhone1" />
            </div>
            <!-- End of Form -->
          </div>
          <div class="col-12">
            <button
              class="btn btn-primary btn-dark mt-2 animate-up-2"
              type="submit">Update</button>
          </div>
        </div>
      </div>
    </form>
    <form action="#" method="post" class="card border-light p-3 mb-4">
      <div class="card-header bg-white border-light p-3 mb-4 mb-md-0">
        <h3 class="h5 mb-0">Card details</h3>
      </div>
      <div class="card-body p-0 p-md-4">
        <div class="row justify-content-center">
          <div class="col-12">
            <div class="mb-4">
              <label class="form-label" for="cardNameLabel"><span
                  class="small text-dark">(Full name as displayed on card)</span></label>
              <div class="input-group mb-4">
                <input
                  class="form-control"
                  id="cardNameLabel"
                  placeholder="Name on Card *"
                  type="text"
                  required="" />
              </div>
            </div>
          </div>
          <div class="col-12">
            <div class="mb-4">
              <label class="form-label h6" for="cardNumberLabel">Card Number <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text"><span
                    class="fas fa-credit-card" /></span>
                <input
                  type="text"
                  class="form-control"
                  id="cardNumberLabel"
                  placeholder="0000 0000 0000 0000"
                  aria-label="Card number"
                  required="" />
              </div>
            </div>
          </div>
          <div class="col-12 col-md-6">
            <div class="mb-4">
              <label class="form-label h6" for="cardCVCLabel">CVC <span class="text-danger">*</span></label>
              <input
                class="form-control"
                id="cardCVCLabel"
                placeholder="CVC"
                type="number"
                required="" />
            </div>
          </div>
          <div class="col-12 col-md-6">
            <div class="mb-4">
              <label class="form-label h6" for="cardExpiryLabel">Card Expiry <span class="text-danger">*</span></label>
              <div class="input-group">
                <span class="input-group-text"><span
                    class="fas fa-credit-card" /></span>
                <input
                  type="text"
                  class="form-control"
                  id="cardExpiryLabel"
                  placeholder="MM / YY"
                  aria-label="Card expiry"
                  required="" />
              </div>
            </div>
          </div>
          <div class="col-12">
            <button
              class="btn btn-primary btn-dark mt-2 animate-up-2"
              type="submit">Update</button>
          </div>
        </div>
      </div>
    </form>
    <div class="card card-body bg-white border-light p-0 p-md-4">
      <div class="card-header bg-white border-0 p-3">
        <h3 class="h5">Order History</h3>
        <p class="small pr-lg-10">
          This is a list of devices that have logged into your account. Revoke
          any sessions that you do not recognize.
        </p>
      </div>
      <div class="card-body px-0 pt-0">
        <ul class="list-group list-group-flush">
          <li class="list-group-item border-bottom py-3">
            <div class="row align-items-center">
              <div class="col">
                <h3 class="h6 mb-1">
                  <a href="./invoice.html">Invoice #120345</a>
                </h3>
                <!-- Text -->
                <small class="text-gray-700"> Billed August 21, 2019 </small>
              </div>
              <div class="col-auto">
                <button class="btn btn-xs btn-outline-dark"> Pay now </button>
              </div>
            </div>
          </li>
          <li class="list-group-item border-bottom py-3">
            <div class="row align-items-center">
              <div class="col">
                <h3 class="h6 mb-1">
                  <a href="./invoice.html">Invoice #120344</a>
                </h3>
                <!-- Text -->
                <small class="text-gray-700"> Billed July 21, 2019 </small>
              </div>
              <div class="col-auto">
                <span class="badge badge-pill badge-success">
                  <span class="text-uppercase font-weight-bold">Paid</span>
                </span>
              </div>
            </div>
          </li>
          <li class="list-group-item border-bottom py-3">
            <div class="row align-items-center">
              <div class="col">
                <h3 class="h6 mb-1">
                  <a href="./invoice.html">Invoice #120343</a>
                </h3>
                <!-- Text -->
                <small class="text-gray-700"> Billed June 21, 2019 </small>
              </div>
              <div class="col-auto">
                <span class="badge badge-pill badge-success">
                  <span class="text-uppercase font-weight-bold">Paid</span>
                </span>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</div>
