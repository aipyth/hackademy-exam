import auth_token from "../utils/auth-token"

const host = process.env.NEXT_PUBLIC_API_PROTOCOL + "://" + process.env.NEXT_PUBLIC_API_HOST + ":" + process.env.NEXT_PUBLIC_API_PORT

export default {
    async requestToken(data) {
        const url = host + '/user/signin'
        const resp = await fetch(url, {
            method: 'POST',
            body: JSON.stringify(data),
        })
        return resp.ok ? await resp.text() : ''
    },

    async createAccount(data) {
        const url = host + '/user/signup'
        const resp = await fetch(url, {
            method: 'POST',
            body: JSON.stringify(data),
        })
        return resp.status === 201
    },

    async createList(data) {
        const url = host + '/todo/lists'
        console.log('Bearer ' + auth_token.getToken())
        const resp = await fetch(url, {
            method: 'POST',
            mode: 'no-cors',
            body: JSON.stringify(data),
            headers: {
                'Authorization': 'Bearer ' + auth_token.getToken(),
            }
        })
        if (resp.status === 201) {
            return await resp.body()
        } else {
            return {'error': await resp.text()}
        }
    },
}
