import { signal } from "@preact/signals-react";
import axios from "axios";

const token = signal('');

export function getToken() {
    return token;
}

export function setToken(value) {
    token.value = value;
}

/**
 * Asynchronously generates a token for authentication by sending a POST request to the server.
 * 
 * @async
 * @function generateToken
 * @param {signal} username - A Preact signal containing the username.
 * @param {signal} password - A Preact signal containing the password.
 * @returns {Promise<void>} Returns a promise that resolves once the token has been generated and assigned to the token signal.
 * 
 * @example
 * import { signal } from '@preact/signals-react'
 * const username = signal("exampleUser");
 * const password = signal("examplePassword");
 * generateToken(username, password);
 * console.log(getToken().value)
 */
export async function generateToken(username, password) {
    let res;
    try {
        res = await axios.post('/api/auth/login', {username: username.value, password: password.value}) 
    } catch (error) {
        console.log(error)
    }

    setToken(res?.data)
    return getToken()
}