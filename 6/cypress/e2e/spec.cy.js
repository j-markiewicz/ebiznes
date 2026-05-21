describe("api", () => {
	it("works", () => {
		cy.request("http://localhost:8000/products")
			.its("status")
			.should("equal", 200);

		cy.request("http://localhost:8000/products/1")
			.its("status")
			.should("equal", 200);
	});

	it("returns products", () => {
		cy.request("http://localhost:8000/products").then(
			(res) =>
				expect(res.body["1"].name).to.equal("Truskawki") &&
				expect(res.body["1"].description).to.equal("opakowanie 500g") &&
				expect(res.body["1"].price).to.equal(699) &&
				expect(res.body["2"].name).to.equal("Ogórki") &&
				expect(res.body["2"].description).to.equal("1kg, świeże") &&
				expect(res.body["2"].price).to.equal(899) &&
				expect(res.body["3"].name).to.equal("Marchewki") &&
				expect(res.body["3"].description).to.equal("1kg luzem") &&
				expect(res.body["3"].price).to.equal(399),
		);
	});

	it("returns a product", () => {
		cy.request("http://localhost:8000/products/1").then(
			(res) =>
				expect(res.body.name).to.equal("Truskawki") &&
				expect(res.body.description).to.equal("opakowanie 500g") &&
				expect(res.body.price).to.equal(699),
		);

		cy.request("http://localhost:8000/products/2").then(
			(res) =>
				expect(res.body.name).to.equal("Ogórki") &&
				expect(res.body.description).to.equal("1kg, świeże") &&
				expect(res.body.price).to.equal(899),
		);

		cy.request("http://localhost:8000/products/3").then(
			(res) =>
				expect(res.body.name).to.equal("Marchewki") &&
				expect(res.body.description).to.equal("1kg luzem") &&
				expect(res.body.price).to.equal(399),
		);
	});
});

describe("product list page", () => {
	it("shows products", () => {
		cy.visit("http://localhost:4173");

		cy.contains("Truskawki");
		cy.contains("Ogórki");
		cy.contains("Marchewki");
	});

	it("supports tab navigation", () => {
		cy.visit("http://localhost:4173");

		cy.get("button").first().focus();
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.SPACE);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.SPACE);
		cy.press(Cypress.Keyboard.Keys.SPACE);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.SPACE);
		cy.press(Cypress.Keyboard.Keys.SPACE);
		cy.press(Cypress.Keyboard.Keys.SPACE);

		cy.get('button[title = "Przejdź do koszyka"]').click();
		cy.contains("[1] Truskawki");
		cy.contains("[2] Ogórki");
		cy.contains("[3] Marchewki");
	});

	it("has correct strawberry information", () => {
		cy.visit("http://localhost:4173");

		cy.contains("Truskawki");
		cy.contains("699¤");
		cy.contains("opakowanie 500g");
	});

	it("has correct cucumber information", () => {
		cy.visit("http://localhost:4173");

		cy.contains("Ogórki");
		cy.contains("899¤");
		cy.contains("1kg, świeże");
	});

	it("has correct carrot information", () => {
		cy.visit("http://localhost:4173");

		cy.contains("Marchewki");
		cy.contains("399¤");
		cy.contains("1kg luzem");
	});

	it("can add a product to cart", () => {
		cy.visit("http://localhost:4173");

		for (let i = 1; i <= 10; i++) {
			cy.get('button[title = "Dodaj do koszyka"]').first().click();
			cy.contains(`${i}`);
		}
	});

	it("can add multiple products to cart", () => {
		cy.visit("http://localhost:4173");

		for (let i = 1; i <= 10; i++) {
			cy.get('button[title = "Dodaj do koszyka"]').first().click();
			cy.contains(`${i}`);
		}

		for (let i = 1; i <= 10; i++) {
			cy.get('button[title = "Dodaj do koszyka"]').click({
				multiple: true,
			});
			cy.contains(`${i}`);
		}
	});

	it("can navigate to cart", () => {
		cy.visit("http://localhost:4173");

		for (let i = 1; i <= 10; i++) {
			cy.get('button[title = "Dodaj do koszyka"]').first().click();
		}

		cy.get('button[title = "Przejdź do koszyka"]').click();
		cy.contains("[10] Truskawki");
		cy.contains("Ogórki").should("not.exist");
		cy.contains("Marchewki").should("not.exist");
	});

	it("multiple products show in cart", () => {
		cy.visit("http://localhost:4173");

		for (let i = 1; i <= 10; i++) {
			cy.get('button[title = "Dodaj do koszyka"]').first().click();
			cy.contains(`${i}`);
		}

		for (let i = 1; i <= 10; i++) {
			cy.get('button[title = "Dodaj do koszyka"]').click({
				multiple: true,
			});
			cy.contains(`${i}`);
			cy.contains(`${i + 10}`);
		}

		cy.get('button[title = "Przejdź do koszyka"]').click();
		cy.contains("[20] Truskawki");
		cy.contains("[10] Ogórki");
		cy.contains("[10] Marchewki");
	});
});

describe("cart page", () => {
	it("shows an empty cart", () => {
		cy.visit("http://localhost:4173/cart");
		cy.contains("Koszyk pusty");
	});

	it("can navigate back", () => {
		cy.visit("http://localhost:4173/cart");
		cy.contains("Koszyk pusty");
		cy.get('button[title = "Wróć do strony głównej"]').click();
		cy.contains("Truskawki");
	});

	it("can pay", () => {
		cy.visit("http://localhost:4173");

		for (let i = 1; i <= 10; i++) {
			cy.get('button[title = "Dodaj do koszyka"]').first().click();
		}

		cy.get('button[title = "Przejdź do koszyka"]').click();

		cy.get("input").first().type("Imie Nazwisko");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().type("1234567891011121");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().click();

		cy.contains("Płatność zakończona");
	});

	it("disallows invalid month", () => {
		cy.visit("http://localhost:4173/cart");

		cy.get("input").first().type("Imie Nazwisko");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().type("1234567891011121");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().type("100");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().click();

		cy.wait(5000);
		cy.contains("Płatność zakończona").should("not.exist");
	});

	it("disallows invalid year", () => {
		cy.visit("http://localhost:4173/cart");

		cy.get("input").first().type("Imie Nazwisko");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().type("1234567891011121");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().type("10000");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().click();

		cy.wait(5000);
		cy.contains("Płatność zakończona").should("not.exist");
	});

	it("requires a name", () => {
		cy.visit("http://localhost:4173/cart");

		cy.get("input").first().focus();
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().type("1234567891011121");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().click();

		cy.wait(5000);
		cy.contains("Płatność zakończona").should("not.exist");
	});

	it("requires a card number", () => {
		cy.visit("http://localhost:4173/cart");

		cy.get("input").first().type("Imie Nazwisko");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().click();

		cy.wait(5000);
		cy.contains("Płatność zakończona").should("not.exist");
	});

	it("disables the form after submitting", () => {
		cy.visit("http://localhost:4173/cart");

		cy.get("input").first().type("Imie Nazwisko");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().type("1234567891011121");
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.press(Cypress.Keyboard.Keys.TAB);
		cy.focused().click();
		cy.get("input").first().should("be.disabled");

		cy.contains("Płatność zakończona");
	});
});
