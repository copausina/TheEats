import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
  withCredentials: true, // Ensures cookies (i.e. JWT) are sent
  headers: {
    "Content-Type": "application/json",
  },
});

export default api;