import type { IAddress } from "@app/models/IAddress";

export class AddressService {
  async normalize(address: IAddress) {
    const result = await fetch(
      "https://us-street.api.smartystreets.com/street-address?auth-id=666b1a30-4f4a-735c-8c34-1f957916e7aa&auth-token=63CwZKLIeojla8vjz43O",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify([
          {
            street: address.address1,
            city: address.city,
            state: address.state,
            zipcode: address.zip,
          },
        ]),
      }
    );
    const body = await result.json();
  }
}
