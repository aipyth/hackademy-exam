const token_key = "auth-token"
const Authed = {
    saveToken(token: string): void {
        localStorage.setItem(token_key, token)
    },

    tokenPresent(): boolean {
        const token: string | null = localStorage.getItem(token_key)
        return token !== null
    },

    getToken(): string {
        const token: string | null = localStorage.getItem(token_key)
        return token || ''
    },

    deleteToken(): void {
        localStorage.removeItem(token_key)
    },
}

export default Authed
