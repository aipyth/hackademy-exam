import React, { useState } from "react"
import router from "next/router"
import { SignUpCallbackProps } from "./types"

export default function SignUpForm({ callback }: { callback: (data: SignUpCallbackProps) => Promise<void> }): JSX.Element {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [confirmPassword, setConfirmPassword] = useState('')

    const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        setEmail(event.target.value)
    }
    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        setPassword(event.target.value)
    }
    const handleConfirmPasswordChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
        setConfirmPassword(event.target.value)
    }

    const handleSubmit = async (event: React.SyntheticEvent) => {
        event.preventDefault()
        await callback({ email, password })
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
            <input
                type="password"
                value={confirmPassword}
                placeholder="Confirm password"
                onChange={handleConfirmPasswordChange}
            />

            <div>
                <button onClick={() => router.back()}>Back</button>
                <button
                    className="signin-submit-button"
                    type="submit">
                    Sign up
                </button>
            </div>
        </form>
    )
}
