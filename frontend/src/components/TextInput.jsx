export default function TextInput(props) {
	return (
		<div class="form-group py-2">
			<label for={props.id}>{props.children}</label>
			<input
				type={props.type || 'text'}
				id={props.id}
				class="form-control"
				minLength={props.minLength}
				maxLength={props.maxLength}
				required={props.required}
				autofocus={props.autofocus}
				placeholder={props.placeholder}
				onInput={props.onInput}
				name={props.name}
				autocomplete={props.autocomplete}
			/>
		</div>
	);
}
