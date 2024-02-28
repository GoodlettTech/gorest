import { effect, signal } from "@preact/signals-react"
import { generateToken, setToken } from "../services/Auth"

export const username = signal('')
export const password = signal('')

effect(() => { username.value || password.value ? setToken(`${username}:${password}`) : console.log('running once')})

export default function AuthPage() {
    return (
        <div className="px-6 py-6 max-w-sm mx-auto bg-zinc-800 text-zinc-400 rounded">
            <form className="flex flex-col" method="POST" action="/api/auth/login">
                <div className="flex flex-col">
                    <label htmlFor="username">Username</label>
                    <input type="text" id="username" name="username" placeholder="username" defaultValue={username} onChange={(e) => username.value = e.target.value}/>
                </div>
                <div className="flex flex-col">
                    <label htmlFor="password">Password</label>
                    <input type="password" id="password" name="password" placeholder="password" defaultValue={password} onChange={(e) => password.value = e.target.value}/>         
                </div>
                <div className="flex-flex-row">
                    <button className="btn round bg-zinc-600 round mt-4 p-1" type="button" onClick={() => generateToken(username, password)}>
                        Login
                    </button>
                </div>            
            </form>
        </div>
    )
}