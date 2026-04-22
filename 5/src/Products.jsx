import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import { API_ROOT } from "./main.jsx";
import cart from "./assets/cart.svg";

function Products({ addToCart }) {
	const [products, setProducts] = useState(null);
	const navigate = useNavigate();

	useEffect(() => {
		fetch(API_ROOT + "/products")
			.then((res) => res.json())
			.then((products) => setProducts(products));
	}, []);

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
								onClick={() => addToCart(i, p)}
							>
								<img src={cart} alt="Koszyk" />
							</button>
						</h2>
						<p>{p.description}</p>
					</section>
				))
			)}
		</>
	);
}

export default Products;
