import { createSignal } from 'solid-js';
export const [jwt, setter] = createSignal(localStorage.getItem('jwt') || '');
export const setJwt = (jwt) => {
	localStorage.setItem('jwt', jwt);
	setter(jwt);
};
