import { A, useNavigate } from '@solidjs/router';
import TextInput from './TextInput';
import { createStore, produce } from 'solid-js/store';
import Form, { setError } from './Form';
import { setJwt } from '../signals/jwt';

export const [form, setForm] = createStore({
	username: '',
	password: '',
});

export default function LoginForm(props) {
	const navigate = useNavigate();
	setError('');

	return (
		<Form
			title="Login"
			onsubmit={async (e) => {
				e.preventDefault();
				console.log(form);

				let response = await fetch(
					'http://localhost:3001/api/users/token',
					{
						method: 'POST',
						body: JSON.stringify({
							username: form.username,
							password: form.password,
						}),
						headers: {
							'Content-Type': 'application/json',
						},
					}
				);

				if (response.status !== 201) {
					setError((await response.json()).message);
					return;
				}

				setError('');

				let token = await response.text();

				setJwt(token);
				navigate('/', { replace: true });
			}}
		>
			<TextInput
				id="loginUsernameInput"
				placeholder="Enter Username"
				name="username"
				autocomplete="username"
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
				id="loginPasswordInput"
				type="password"
				name="password"
				autocomplete="pass"
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
			<div className="row justify-content-center">
				<div className="col-sm-6 col-8 mt-2 justify-content-center d-flex">
					<button type="submit" class="btn btn-primary">
						Login
					</button>
				</div>
				<div className="col-sm-6 col-8 align-self-center justify-content-center d-flex">
					<A href="/createuser" class="">
						Create a user
					</A>
				</div>
			</div>
		</Form>
	);
}
