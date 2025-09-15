import api from "./api";

export async function login(email, password) {
    const res = await api.post("/auth/login", {email, password});
    localStorage.setItem("token", res.data.token);
    return res.data;
}

export async function register(username, email, password) {
    return api.post("/auth/register", {username, email, password});
}

export function logout() {
    localStorage.removeItem("token");
}