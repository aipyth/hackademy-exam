import Image from 'next/image'
import OpenwareLogo from '../public/logo2.svg'

export default function (): JSX.Element {
    return <span>
        <Image
            src={OpenwareLogo}
            alt="Openware Logo"
            width={35}
            height={24}
        />
    </span>
}
