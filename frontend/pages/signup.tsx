import useToken from "../hooks/useToken"
import { useRouter } from "next/router"
import OpenwareLogo from "../containers/OpenwareLogo"
import SignUpForm from "../components/SignUpForm"
import { SignUpCallbackProps } from "../components/types"
import {useEffect} from "react"
import requests from "../utils/requests"

function SignUp(): JSX.Element {
    const [token, setToken] = useToken()
    const router = useRouter()

    useEffect((): void => {
        if (token != '') {
            console.log("Not empty token. Redirecting to /")
            router.push('/')
        }
    })

    const formCallback = async (data: SignUpCallbackProps) => {
        const created = await requests.createAccount(data)
        
        if (created) {
            const token = await requests.requestToken(data)
            if (token === '') {
                console.log('Wrong email or password')
            } else {
                setToken(token)
            }
        }
    }

    return (
        <div className="signin-container">
            <div className="signin-card">
                <div>
                    <OpenwareLogo/>
                    <span className="">Todo</span>
                </div>
                <h4>Sign up</h4>
                <SignUpForm callback={formCallback}/>
            </div>
        </div>
    )
}

export default SignUp
