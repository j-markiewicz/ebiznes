import { useState } from "react";
import { BrowserRouter, Route, Routes } from "react-router";
import "./App.css";
import Payments from "./Payments";
import Cart from "./Cart";
import Products from "./Products";

function App() {
	const [cartItems, setCartItems] = useState([]);
	const addToCart = (id, product) => {
		let items = [...cartItems];
		const existing = cartItems.findIndex((i) => i.id === id);

		if (existing === -1) {
			items.push({ id, product, amount: 1 });
		} else {
			items[existing].amount += 1;
		}

		setCartItems(items);
	};

	return (
		<BrowserRouter>
			<Routes>
				<Route index element={<Products addToCart={addToCart} />} />

				<Route
					path="cart"
					element={
						<>
							<Cart cartItems={cartItems} />
							<Payments />
						</>
					}
				/>
			</Routes>
		</BrowserRouter>
	);
}

export default App;
