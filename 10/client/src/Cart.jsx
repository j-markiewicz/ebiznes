import { useNavigate } from "react-router";
import PropTypes from "prop-types";
import back from "./assets/back.svg";

function Cart({ cartItems }) {
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
							[{i.amount}] {i.product.name} |{" "}
							{i.product.price * i.amount}¤ ({i.product.price}
							¤/szt.)
						</h2>
						<p>{i.product.description}</p>
					</section>
				))
			)}
		</>
	);
}

Cart.propTypes = {
	cartItems: PropTypes.arrayOf(
		PropTypes.exact({
			id: PropTypes.string.isRequired,
			product: PropTypes.exact({
				name: PropTypes.string.isRequired,
				price: PropTypes.number.isRequired,
				description: PropTypes.string.isRequired,
			}),
			amount: PropTypes.number.isRequired,
		}),
	).isRequired,
};

export default Cart;
