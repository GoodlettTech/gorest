import { signal } from "@preact/signals-react";

const token = signal("");

export function getToken() {
	return token;
}

export function setToken(value) {
	token.value = value;
}