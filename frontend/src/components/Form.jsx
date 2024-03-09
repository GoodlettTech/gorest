import { Show, createSignal } from 'solid-js';

export const [error, setError] = createSignal('');

export default function Form(props) {
	return (
		<div class="card">
			<h2 class="card-title text-center mt-4">{props.title}</h2>
			<div class="card-body">
				<Show when={error().length !== 0}>
					<div class="row">
						<div class="col">
							<p class="text-center text-danger text-xl font-weight-bold">
								{error()}
							</p>
						</div>
					</div>
				</Show>
				<form onsubmit={props.onsubmit}>{props.children}</form>
			</div>
		</div>
	);
}
