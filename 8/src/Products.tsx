import { useNavigate } from "react-router";
import { useCart, useProducts } from "./state.js";
import cart from "./assets/cart.svg";

function Products() {
	const [cartItems, addToCart] = useCart();
	const [products, setProducts] = useProducts();
	const navigate = useNavigate();

	return (
		<>
			<h1>
				<button
					title="Przejdź do koszyka"
					className="cart"
					onClick={() => navigate("/cart")}
				>
					<img src={cart} alt="Koszyk" />
				</button>{" "}
				Produkty:
			</h1>

			{products === null ? (
				<h2>Ładowanie...</h2>
			) : (
				Object.entries(products).map(([i, p]) => (
					<section key={i}>
						<h2>
							{p.name} | {p.price}¤{" "}
							<button
								title="Dodaj do koszyka"
								className="cart"
								onClick={() => addToCart(i)}
							>
								<img src={cart} alt="Koszyk" />
							</button>
							{cartItems.findIndex((ci) => ci.id == i) !== -1 ? (
								<sub>
									{cartItems.find((ci) => ci.id == i)?.amount}
								</sub>
							) : null}
						</h2>
						<p>{p.description}</p>
					</section>
				))
			)}
		</>
	);
}

export default Products;
