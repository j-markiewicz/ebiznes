import { useState } from "react";
import { API_ROOT } from "./main";
import axios from "axios";

function Payments() {
	const [state, setState] = useState("initial");
	const [name, setName] = useState("");
	const [card, setCard] = useState("");
	const [month, setMonth] = useState(1);
	const [year, setYear] = useState(2026);

	const submit = () => {
		setState("busy");

		axios
			.post(API_ROOT + "/pay", {
				name,
				card,
				month,
				year,
			})
			.then(() => setState("done"));
	};

	return (
		<>
			<h1>Informacje płatnicze:</h1>
			{state === "done" ? (
				<p>Płatność zakończona</p>
			) : (
				<form
					onSubmit={(e) => {
						e.preventDefault();
						submit();
					}}
				>
					<label>
						Imie na karcie:{" "}
						<input
							disabled={state === "busy"}
							size={25}
							minLength={1}
							type="text"
							required
							value={name}
							onChange={(e) => setName(e.value)}
						/>
					</label>
					<label>
						Numer karty:{" "}
						<input
							disabled={state === "busy"}
							size={25}
							minLength={16}
							maxLength={16}
							type="text"
							required
							value={card}
							onChange={(e) => setCard(e.value)}
						/>
					</label>
					<label>
						Data ważności:{" "}
						<input
							disabled={state === "busy"}
							type="number"
							min="1"
							max="12"
							size={2}
							required
							value={month}
							onChange={(e) => setMonth(e.value)}
						/>
						.
						<input
							disabled={state === "busy"}
							type="number"
							min="2000"
							max="2100"
							size={4}
							required
							value={year}
							onChange={(e) => setYear(e.value)}
						/>
					</label>

					<button disabled={state === "busy"} type="submit">
						Zapłać
					</button>
				</form>
			)}
		</>
	);
}

export default Payments;
