import axios from "axios";

// Base instance for all requests
const baseURL = "http://localhost:8080";

// Public API instance
export const publicApi = axios.create({
  baseURL,
  withCredentials: true, // Ensures cookies are sent
});

// Authenticated API instance (with token interceptor)
export const authApi = axios.create({
  baseURL,
  withCredentials: true,
});

// Request Interceptor
authApi.interceptors.request.use(
  async (config) => {
    const accessToken = localStorage.getItem("access_token");
    if (accessToken) {
      config.headers["Authorization"] = `Bearer ${accessToken}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response Interceptor (Handles Expired Tokens)
authApi.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response.status === 401) { // if unauthorized, assume token is expired
      try {
        // Request a new access token using the refresh token
        const res = await authApi.post("/auth/refresh");
        localStorage.setItem("access_token", res.data.access_token);

        // Retry the original request with the new token
        error.config.headers["Authorization"] = `Bearer ${res.data.access_token}`;
        return authApi(error.config);
      } catch (refreshError) {
        console.error("Session expired. Please log in again.");
        window.location.href = "/login"; // Redirect to login
      }
    }
    return Promise.reject(error);
  }
);

export default authApi;