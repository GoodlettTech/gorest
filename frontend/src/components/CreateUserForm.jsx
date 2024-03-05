import { A, useNavigate } from '@solidjs/router';
import TextInput from './TextInput';
import { createStore, produce } from 'solid-js/store';
import Form, { setError } from './Form';
import { setJwt } from '../signals/jwt';

export const [form, setForm] = createStore({
	email: '',
	username: '',
	password: '',
	confirm: '',
});

const navigate = useNavigate();
async function handleSubmit(e) {
	e.preventDefault();
	console.log(form);

	if (form.password !== form.confirm) {
		setError('Passwords must match');
		return;
	}

	let response = await fetch('http://localhost:3001/api/auth/createuser', {
		method: 'POST',
		body: JSON.stringify({
			email: form.email,
			username: form.username,
			password: form.password,
			confirm: form.confirm,
		}),
		headers: {
			'Content-Type': 'application/json',
		},
	});

	if (response.status !== 201) {
		setError((await response.json()).message);
		return;
	}

	setError('');

	let token = await response.text();

	setJwt(token);
	navigate('/', { replace: true });
}

export default function CreateUserForm() {
	setError('');

	return (
		<Form title="Create User" onsubmit={handleSubmit}>
			<TextInput
				id="emailInput"
				placeholder="Enter Email"
				type="email"
				required={true}
				onInput={(e) =>
					setForm(
						produce((currentForm) => {
							currentForm.email = e.target.value;
						})
					)
				}
				value={form.email}
			>
				Email
			</TextInput>
			<TextInput
				id="usernameInput"
				placeholder="Enter Username"
				required={true}
				minLength={4}
				maxLength={16}
				onInput={(e) =>
					setForm(
						produce((currentForm) => {
							currentForm.username = e.target.value;
						})
					)
				}
				value={form.username}
			>
				Username
			</TextInput>
			<TextInput
				id="passwordInput"
				type="password"
				placeholder="Enter Password"
				required={true}
				minLength={8}
				maxLength={32}
				onInput={(e) =>
					setForm(
						produce((currentForm) => {
							currentForm.password = e.target.value;
						})
					)
				}
				value={form.password}
			>
				Password
			</TextInput>
			<TextInput
				id="confirmPasswordInput"
				type="password"
				placeholder="Confirm Password"
				required={true}
				minLength={8}
				maxLength={32}
				onInput={(e) =>
					setForm(
						produce((currentForm) => {
							currentForm.confirm = e.target.value;
						})
					)
				}
				value={form.confirm}
			>
				Password
			</TextInput>
			<div className="row justify-content-center">
				<div className="col-sm-6 col-8 mt-2 justify-content-center d-flex">
					<button type="submit" class="btn btn-primary">
						Create User
					</button>
				</div>
				<div className="col-sm-6 col-8 align-self-center justify-content-center d-flex">
					<A href="/login">Already a user? Log In</A>
				</div>
			</div>
		</Form>
	);
}
