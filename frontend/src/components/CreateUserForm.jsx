import { signal } from "@preact/signals-react";
import { setToken } from "../services/Auth";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

const creds = signal({
	email: "",
	username: "",
	password: "",
	confirmPassword: "",
});

function handleInput(event) {
	creds.value = {
		...creds.value,
		[event.target.name]: event.target.value,
	};
}

const errorMessage = signal("");

export default function CreateUserForm() {
	let navigate = useNavigate()
	
	async function handleCreateUser(event) {
		let res;
		try {
			errorMessage.value = "";
			res = await axios.post("http://localhost:3001/api/auth/createuser", {
				email: creds.value.email,
				username: creds.value.username,
				password: creds.value.password,
				confirmPassword: creds.value.confirmPassword,
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
				<label htmlFor="email">Email</label>
				<input
					type="text"
					id="email"
					name="email"
					placeholder="email"
					defaultValue={creds.value.email}
					onChange={handleInput}
				/>
			</div>
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
			<div className="flex flex-col">
				<label htmlFor="confirmPassword">Confirm Password</label>
				<input
					type="password"
					id="confirmPassword"
					name="confirmPassword"
					placeholder="Confirm Password"
					defaultValue={creds.value.confirmPassword}
					onChange={handleInput}
				/>
			</div>
			<div className="error-message">{errorMessage}</div>
			<div className="flex flex-row">
				<Link
					className="mt-4 p-1 mx-auto"
					to="../login"
				>
					Login instead
				</Link>
				<button className="btn bg-zinc-600 round mt-4 p-1 mx-auto" onClick={handleCreateUser}>
					Create Account
				</button>
			</div>
		</form>
	);
}
