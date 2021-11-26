import React, { useState } from "react"
import Link from "next/link"
import { SignInCallbackProps } from "./types"


export default function SignInForm({ callback }: { callback: (data: SignInCallbackProps) => void }): JSX.Element {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        setEmail(event.target.value)
    }
    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        setPassword(event.target.value)
    }

    const handleSubmit = (event: React.SyntheticEvent): void => {
        event.preventDefault()
        callback({ email, password })
    }

    return ( 
        <form onSubmit={handleSubmit}
            className=".signin-form">
            <input
                className=".signin-input"
                type="email"
                value={email}
                placeholder="Email"
                onChange={handleEmailChange}
            />
            <input
                className=".signin-input"
                type="password"
                value={password}
                placeholder="Password"
                onChange={handlePasswordChange}
            />

            <div>
                No account?
                <span>
                    <Link href="/signup">Create one!</Link>
                </span>
            </div>
            <div >Forgot password?</div>
            <div>
                <button
                    className="signin-submit-button"
                    type="submit">
                    Sign in
                </button>
            </div>
        </form>
    )
}
