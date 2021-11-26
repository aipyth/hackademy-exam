import useToken from "../hooks/useToken"
import { useRouter } from "next/router"
import OpenwareLogo from "../containers/OpenwareLogo"
import SignInForm from "../components/SignInForm"
import { SignInCallbackProps } from "../components/types"
import {useEffect} from "react"
import requests from "../utils/requests"

function SignIn(): JSX.Element {
    const [token, setToken] = useToken()
    const router = useRouter()

    useEffect((): void => {
        if (token != '') {
            console.log("Not empty token. Redirecting to /")
            router.push('/')
        }
    })

    const formCallback = (data: SignInCallbackProps): void => {
        requests.requestToken(data).then((token: string) => {
            if (token === '') {
                console.log('Wrong email or password')
            } else {
                setToken(token)
            }
        })
    }

    return (
        <div className="signin-container">
            <div className="signin-card">
                <div>
                    <OpenwareLogo/>
                    <span className="text-gray-500 text-2xl font-semibold px-1">Todo</span>
                </div>
                <h4 className="text-black text-xl mb-2">Sign in</h4>
                <SignInForm callback={formCallback}/>
            </div>
        </div>
    )
}

export default SignIn
