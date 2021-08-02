import {getByTestId, getFormInput} from "../extension/extensions";
import '@testing-library/cypress/add-commands'
import 'cypress-file-upload'


describe("Form Submission Validations", () => {

    it("expects field with validation should show error if validation fails", () => {
        cy.visit(`${Cypress.env('HOST')}form/preview?formId=d43ca729-576b-488e-8f03-0bceed2a7b91&cypress_test=true`)
        cy.findByText('Validation Test Form').should('exist');
        cy.findByText('Validation Test Form Description').should("exist")
        assertOptionalFieldValidatedIfHasValue();
        assertOptionalFieldNotValidated();
        assertDateTimeOnlyValidation();
        assertDateAndTimeValidation()
        assertDateOnlyValidation();
        assertCheckboxTwoRequired();
        assertRadioOtherInput();
        assertFileIsPdfOrText();
        assertSubmission();
    });

    function assertOptionalFieldValidatedIfHasValue() {
        getByTestId("submit-form-button").should("be.enabled");
        getFormInput("text-input", "Optional Field With Email Address Validation", false).type("hello");
        cy.findByText("Please enter a valid email address.").should('exist');
        getByTestId("submit-form-button").should("be.disabled");
        getFormInput("text-input", "Optional Field With Email Address Validation", false).clear();
        cy.findByText("Please enter a valid email address.").should('not.exist');
        getFormInput("text-input", "Optional Field With Email Address Validation", false).type("test@logicful.org");
        cy.findByText("Please enter a valid email address.").should('not.exist');
        getByTestId("submit-form-button").should("be.enabled");
    }

    function assertOptionalFieldNotValidated() {
        getFormInput("text-input", "Optional Field With Email Address Validation", false)
            .clear().type("value").clear();
        getByTestId("submit-form-button").should("be.enabled");
        getFormInput("text-input", "Optional Field With Email Address Validation", false).should("exist")
    }

    function assertDateTimeOnlyValidation() {

        getFormInput("date-input", "Date With Time Only Validation 9-5", false).type("2020-06-20")
            .should("not.have.class", "input-error");

        getFormInput("time-input", "Date With Time Only Validation 9-5", false).wait(300).type("03:00")
            .should("have.class", "input-error").clear().wait(300).type("11:00")
            .should("not.have.class", "input-error");

        cy.findAllByText("Time must be between 9:00 AM - 5:00 PM.").should("not.exist")
    }

    function assertDateAndTimeValidation() {
        getFormInput("date-input", "Date With Date and Time Validation 3-14 - 3-16", false).type("2020-06-20")
            .should("have.class", "input-error")

        cy.findAllByText("Date must be between Mar 14, 2020, 5:00 AM - Mar 16, 2020, 8:00 AM.").should("exist")

        getFormInput("date-input", "Date With Date and Time Validation 3-14 - 3-16", false)
            .clear().type("2020-03-16").should("not.have.class", "input-error");

        cy.findAllByText("Date must be between Mar 14, 2020, 5:00 AM - Mar 16, 2020, 8:00 AM.").should("not.exist")

        getFormInput("time-input", "Date With Date and Time Validation 3-14 - 3-16", false).wait(300).type("09:00")
            .should("have.class", "input-error");

        cy.findAllByText("Date must be between Mar 14, 2020, 5:00 AM - Mar 16, 2020, 8:00 AM.").should("exist")

        getFormInput("time-input", "Date With Date and Time Validation 3-14 - 3-16", false).wait(300).clear()
            .type("07:00")
            .should("not.have.class", "input-error");

        cy.findAllByText("Date must be between Mar 14, 2020, 5:00 AM - Mar 16, 2020, 8:00 AM.").should("not.exist")
    }

    function assertDateOnlyValidation() {
        getFormInput("date-input", "Date With Date Only Validation", false).type("2020-06-20")
            .should("have.class", "input-error")

        cy.findAllByText("Date must be between Mar 14, 2020, 12:00 AM - Mar 16, 2020, 12:00 AM.").should("exist")

        getFormInput("date-input", "Date With Date Only Validation", false)
            .clear().type("2020-03-16").should("not.have.class", "input-error");
    }

    function assertSubmission() {
        getByTestId("submit-form-button").should("be.enabled");
        getByTestId("submit-form-button").click();
        cy.findByText("Validation test has been completed.").should('exist');
    }

    function assertCheckboxTwoRequired() {
        getFormInput("checkbox-input", "Checkbox Two Values Required-Option 1", true).click()
            .should("have.class", "border-red-300")
        cy.findByText("You must select at-least two values.").should("exist")
        getFormInput("checkbox-input", "Checkbox Two Values Required-Option 2", true).click()
            .should("not.have.class", "border-red-300")
        getFormInput("checkbox-input", "Checkbox Two Values Required-Option 1", true)
            .should("not.have.class", "border-red-300")
        cy.findByText("You must select at-least two values.").should("not.exist")
    }

    function assertRadioOtherInput() {
        getFormInput("radio-input", "Radio Must Type Other-Option 1", true).click()
        cy.findByText("Must type in the other block.").should("exist")
        getFormInput("radio-input", "Radio Must Type Other-other-input", true).type("This is in the other block.")
        cy.findByText("Must type in the other block.").should("not.exist")
    }

    function assertFileIsPdfOrText() {
        getFormInput("file-input-hidden", "File Must Be Pdf Or Txt", true)
            .attachFile('./../fixtures/dummy.png');
        cy.findByText("Must be a pdf or a txt file.").should("exist")
        getFormInput("file-input", "File Must Be Pdf Or Txt", true)
            .should("have.class", "input-error").should("have.value", "dummy.png")
        getFormInput("file-input-clear", "File Must Be Pdf Or Txt", true).click();
        getFormInput("file-input", "File Must Be Pdf Or Txt", true)
            .should("have.class", "input-error").should("have.value", "")
        getFormInput("file-input-hidden", "File Must Be Pdf Or Txt", true)
            .attachFile('./../fixtures/dummy.txt');
        cy.findByText("Must be a pdf or a txt file.").should("not.exist")
        getFormInput("file-input-clear", "File Must Be Pdf Or Txt", true).click();
        getFormInput("file-input-hidden", "File Must Be Pdf Or Txt", true)
            .attachFile('./../fixtures/dummy.pdf');
        cy.findByText("Must be a pdf or a txt file.").should("not.exist")
        getFormInput("file-input", "File Must Be Pdf Or Txt", true)
            .should("not.have.class", "input-error").should("have.value", "dummy.pdf")
    }
});
