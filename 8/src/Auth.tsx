import { useState } from "react";
import { API_ROOT } from "./main";
import axios, { toFormData } from "axios";

function Auth() {
	const [state, setState] = useState("initial");
	const [loginError, setLoginError] = useState("");
	const [signupError, setSignupError] = useState("");
	const [loginUsername, setLoginUsername] = useState("");
	const [loginPassword, setLoginPassword] = useState("");
	const [signupUsername, setSignupUsername] = useState("");
	const [signupPassword, setSignupPassword] = useState("");

	const submitLogin = () => {
		setState("busy");

		axios
			.postForm(API_ROOT + "/login", { loginUsername, loginPassword })
			.then(() => {
				setState("done");
				setSignupError("");
				setSignupUsername("");
				setSignupPassword("");
			})
			.catch((err) => {
				if (err.response) {
					setLoginError(
						`Błąd: ${err.response.data} (${err.response.status})`,
					);
				} else {
					setLoginError(`Błąd: ${err} (${err.status})`);
				}

				setState("error");
			});
	};

	const submitSignup = () => {
		setState("busy");

		axios
			.postForm(API_ROOT + "/signup", { signupUsername, signupPassword })
			.then(() => {
				setState("intial");
				setSignupError("");
				setSignupUsername("");
				setSignupPassword("");
			})
			.catch((err) => {
				if (err.response) {
					setSignupError(
						`Błąd: ${err.response.data} (${err.response.status})`,
					);
				} else {
					setSignupError(`Błąd: ${err} (${err.status})`);
				}

				setState("error");
			});
	};

	return (
		<>
			{state === "done" ? (
				<p>Zalogowane</p>
			) : (
				<div className="forms">
					<form
						onSubmit={(e) => {
							e.preventDefault();
							submitLogin();
						}}
					>
						<h2>Zaloguj</h2>
						{state === "error" ? <p>{loginError}</p> : null}
						<label>
							Nazwa użytkownika:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={1}
								type="text"
								required
								value={loginUsername}
								onChange={(e) =>
									setLoginUsername(e.currentTarget.value)
								}
							/>
						</label>
						<label>
							Hasło:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={1}
								maxLength={128}
								type="password"
								required
								value={loginPassword}
								onChange={(e) =>
									setLoginPassword(e.currentTarget.value)
								}
							/>
						</label>

						<button disabled={state === "busy"} type="submit">
							Zaloguj
						</button>
					</form>
					<form
						onSubmit={(e) => {
							e.preventDefault();
							submitSignup();
						}}
					>
						<h2>Utwórz konto</h2>
						{state === "error" ? <p>{signupError}</p> : null}
						<label>
							Nazwa użytkownika:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={1}
								type="text"
								required
								value={signupUsername}
								onChange={(e) =>
									setSignupUsername(e.currentTarget.value)
								}
							/>
						</label>
						<label>
							Hasło:{" "}
							<input
								disabled={state === "busy"}
								size={25}
								minLength={1}
								maxLength={128}
								type="password"
								required
								value={signupPassword}
								onChange={(e) =>
									setSignupPassword(e.currentTarget.value)
								}
							/>
						</label>

						<button disabled={state === "busy"} type="submit">
							Utwórz konto
						</button>
					</form>
				</div>
			)}
		</>
	);
}

export default Auth;
