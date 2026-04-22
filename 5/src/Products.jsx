import { useEffect, useState } from "react";
import { API_ROOT } from "./main.jsx";

function Products() {
	const [products, setProducts] = useState(null);
	useEffect(() => {
		fetch(API_ROOT + "/products")
			.then((res) => res.json())
			.then((products) => setProducts(products));
	}, []);

	return (
		<>
			<h1>Produkty:</h1>

			{products === null ? (
				<h2>Ładowanie...</h2>
			) : (
				Object.entries(products).map(([i, p]) => (
					<section key={i}>
						<h2>
							{p.name} | {p.price}¤
						</h2>
						<p>{p.description}</p>
					</section>
				))
			)}
		</>
	);
}

export default Products;
