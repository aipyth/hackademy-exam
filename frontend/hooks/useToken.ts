import { useEffect, useState } from 'react'

import auth_token from '../utils/auth-token'

export default function (): [string, (token: string) => void] {
    const [stateToken, setStateToken] = useState('')
    useEffect(() => {
        if (auth_token.tokenPresent() &&
            auth_token.getToken() !== stateToken) {
            // console.log("resetting token in hook", {
            //     at: auth_token.getToken(),
            //     ht: stateToken,
            // })
            setStateToken(auth_token.getToken())
            // console.log("reseted token in hook", {
            //     at: auth_token.getToken(),
            //     ht: stateToken,
            // })
        }
    })
    const setToken = (token: string): void => {
        setStateToken(token)
        auth_token.saveToken(token)
    }
    return [stateToken, setToken]
}


