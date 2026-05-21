import { useEffect, useState } from "react";
import axios from "axios";
import { BrowserRouter, Route, Routes } from "react-router";
import { AppState, CartItem, Product } from "./state";
import { API_ROOT } from "./main";
import Payments from "./Payments";
import Cart from "./Cart";
import Products from "./Products";

function App() {
	const [products, setProducts] = useState<{ [key: string]: Product } | null>(
		null,
	);
	const [cartItems, setCartItems] = useState<CartItem[]>([]);

	useEffect(() => {
		axios
			.get(API_ROOT + "/products")
			.then((res) => res.data)
			.then((products) => setProducts(products));
	}, []);

	return (
		<AppState value={{ products, setProducts, cartItems, setCartItems }}>
			<BrowserRouter>
				<Routes>
					<Route index element={<Products />} />

					<Route
						path="cart"
						element={
							<>
								<Cart />
								<Payments />
							</>
						}
					/>
				</Routes>
			</BrowserRouter>
		</AppState>
	);
}

export default App;
