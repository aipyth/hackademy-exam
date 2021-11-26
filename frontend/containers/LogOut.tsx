import auth_token from "../utils/auth-token"
import { useRouter } from "next/router"

export default function() {
    const router = useRouter()

    const logOut = () => {
        auth_token.deleteToken()
        router.push('/signin')
    }

    return (
        <div
            className="cursor-pointer"
            onClick={logOut}
        >
            Log Out
        </div>
    )
}
