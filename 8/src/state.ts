import { Context, createContext, useContext } from "react";

export type State = {
	cartItems: CartItem[];
	setCartItems: (newCartItems: CartItem[]) => void;
	products: {
		[id: string]: Product;
	} | null;
	setProducts: (
		newProducts: {
			[id: string]: Product;
		} | null,
	) => void;
};

export type Product = {
	name: string;
	description: string;
	price: number;
};

export type CartItem = {
	id: string;
	amount: number;
};

export const AppState: Context<State> = createContext({
	cartItems: [],
	setCartItems: () => {
		console.error("AppState is not initialized");
	},
	products: null,
	setProducts: () => {
		console.error("AppState is not initialized");
	},
} as State);

export const useProducts = (): [
	products: { [id: string]: Product } | null,
	setProducts: (
		newProducts: {
			[id: string]: Product;
		} | null,
	) => void,
] => {
	const appState = useContext(AppState);
	return [appState.products, appState.setProducts];
};

export const useCart = (): [
	cart: CartItem[],
	addToCart: (id: string) => void,
] => {
	const appState = useContext(AppState);

	return [
		appState.cartItems,
		(id: string) => {
			let items = [...appState.cartItems];
			const existing = items.findIndex((i) => i.id === id);
			if (existing === -1) {
				items.push({ id, amount: 1 });
			} else {
				items[existing].amount += 1;
			}
			appState.setCartItems(items);
		},
	];
};
