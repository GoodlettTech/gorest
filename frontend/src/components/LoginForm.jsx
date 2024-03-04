import { signal } from "@preact/signals-react";
import { setToken } from "../services/Auth";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

const errorMessage = signal("");
const creds = signal({
	username: "",
	password: "",
});

function handleInput(event) {
	creds.value = {
		...creds.value,
		[event.target.name]: event.target.value,
	};
}

export default function LoginForm() {
	let navigate = useNavigate();
	async function handleLogin(event) {
		let res;
		try {
			errorMessage.value = "";
			res = await axios.post("http://localhost:3001/api/auth/login", {
				username: creds.value.username,
				password: creds.value.password,
			});
		} catch (error) {
			errorMessage.value = error.response.data.message;
			return;
		}

		setToken(res?.data);
		navigate('/')
	}

	return (
		<form className="flex flex-col" onSubmit={(e) => e.preventDefault()}>
			<div className="flex flex-col">
				<label htmlFor="username">Username</label>
				<input
					type="text"
					id="username"
					name="username"
					placeholder="username"
					defaultValue={creds.value.username}
					onChange={handleInput}
				/>
			</div>
			<div className="flex flex-col">
				<label htmlFor="password">Password</label>
				<input
					type="password"
					id="password"
					name="password"
					placeholder="password"
					defaultValue={creds.value.password}
					onChange={handleInput}
				/>
			</div>
			<div className="error-messages">{errorMessage}</div>
			<div className="flex flex-row">
				<button
					className="btn bg-zinc-600 round mt-4 p-1 mx-auto"
					onClick={handleLogin}
				>
					Login
				</button>
				<Link className="mt-4 p-1 mx-auto" to="../createuser">
					Create User instead
				</Link>
			</div>
		</form>
	);
}
