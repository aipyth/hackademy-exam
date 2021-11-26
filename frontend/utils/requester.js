import axios from "axios";

const instance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_PROTOCOL + "://" + process.env.NEXT_PUBLIC_API_HOST + ":" + process.env.NEXT_PUBLIC_API_PORT,
    timeout: 1000,
    headers: {

    },
});

export default instance
