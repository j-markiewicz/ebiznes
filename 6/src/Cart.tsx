import { useNavigate } from "react-router";
import { useCart, useProducts } from "./state.js";
import back from "./assets/back.svg";

function Products() {
	const [cartItems, addToCart] = useCart();
	const [products, setProducts] = useProducts();
	const navigate = useNavigate();

	return (
		<>
			<h1>
				<button
					title="Wróć do strony głównej"
					className="back"
					onClick={() => navigate("/")}
				>
					<img src={back} alt="Wróć" />
				</button>{" "}
				Koszyk:
			</h1>

			{(cartItems?.length ?? 0) === 0 ? (
				<p>Koszyk pusty</p>
			) : (
				cartItems.map((i) => (
					<section key={i.id}>
						<h2>
							[{i.amount}] {products?.[i.id]?.name} |{" "}
							{products?.[i.id]?.price ?? NaN * i.amount}¤ (
							{products?.[i.id]?.price}
							¤/szt.)
						</h2>
						<p>{products?.[i.id]?.description}</p>
					</section>
				))
			)}
		</>
	);
}

export default Products;
