export function getByTestId(id : string) {
    return cy.get(`[data-test-id="${id}"]`)
}

export function getFormInput(type : string, label : string, required : boolean) {
    return getByTestId(`${type}-${label}-${required ? 'required' : 'optional'}`);
}
