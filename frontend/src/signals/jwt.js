import { createSignal } from 'solid-js';
import Cookies from 'js-cookie';

const [jwt, setter] = createSignal(Cookies.get('Auth') || '');

const setJwt = (jwt) => {
	Cookies.set('Auth', token, { expires: 7, path: '', secure: true, sameSite: 'strict' })
	setter(jwt);
};

export default [jwt, setJwt];