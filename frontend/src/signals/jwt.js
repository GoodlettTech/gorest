import { createSignal } from 'solid-js';

let existingToken = localStorage.getItem('jwt');

export const [jwt, setter] = createSignal(existingToken || '');

export const setJwt = (jwt) => {
	localStorage.setItem('jwt', jwt);
	setter(jwt);
};
