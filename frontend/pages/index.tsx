import {useRouter} from 'next/router';
import {useEffect} from 'react';
import ListsColumn from '../components/ListsColumn';
import LogOut from '../containers/LogOut';
import OpenwareLogo2 from '../containers/OpenwareLogo2';
import auth_token from "../utils/auth-token"

function Homepage(): JSX.Element {
    const router = useRouter()

    useEffect((): void => {
        if (!auth_token.tokenPresent()) {
            router.push('/signin')
        }
    })

    return (
        <div className="w-full h-screen flex flex-row">
            <div className="w-1/6 h-screen bg-yellow-400 p-8 flex flex-col justify-between">
                <div>
                    <OpenwareLogo2/>
                    <ListsColumn/> 
                </div>
                <LogOut/>
            </div>

            
        </div>
    )
}

export default Homepage
