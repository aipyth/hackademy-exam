import Head from 'next/head'
import type { AppProps } from 'next/app'

import '../styles/base.css'

function App({ Component, pageProps }: AppProps) {
    return (
        <>
            <Head>
                <title>Hackademy Todo</title>
                <meta charSet="utf-8" />
                <meta name="viewport" content="initial-scale=1.0, width=device-width" />
            </Head>
            <Component {...pageProps} />
        </>
    )
}

export default App
