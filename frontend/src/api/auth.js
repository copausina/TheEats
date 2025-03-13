import { publicApi, authApi } from "../api/axiosInstance";

// Set default API base URL
const API_URL = "http://localhost:8080/auth";

// Register new user
export const register = async (email, password) => {
  try {
    const response = await publicApi.post(`${API_URL}/register`, {
      email,
      password,
    });
    return response.data;
  } catch (error) {
    throw error.response?.data || "Registration failed";
  }
};

// Login user
export const login = async (email, password) => {
  try {
    const response = await publicApi.post(`${API_URL}/login`, {
      email,
      password,
    }, { withCredentials: true }); // Send cookies

    return response.data;
  } catch (error) {
    throw error.response?.data || "Login failed";
  }
};

// Logout user
export const logout = async () => {
  try {
    await authApi.post(`${API_URL}/logout`, {}, { withCredentials: true });
  } catch (error) {
    console.error("Logout failed", error);
  }
};