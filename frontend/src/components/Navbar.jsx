import { Link, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { logout } from "../api/auth"; 
import axiosInstance from "../api/axiosInstance"
import "./Navbar.css";

const Navbar = () => {
  const navigate = useNavigate();
  const [isAuthenticated, setIsAuthenticated] = useState(true);

  useEffect(() => {
    // Call backend to check authentication
    axiosInstance.get("http://localhost:8080/auth/check", { withCredentials: true })
      .then(response => {
        setIsAuthenticated(response.data.authenticated);
      })
      .catch(() => {
        setIsAuthenticated(false);
      });
  }, []);

  const handleLogout = async () => {
    await logout();
    setIsAuthenticated(false); // Update state
    navigate("/login"); // Redirect to login
  };
  
  return (
    <nav className="navbar">
      <div className="nav-container">
        <h1 className="logo">TheEats</h1>
        <ul className="nav-links">
          <li><Link to="/">Home</Link></li>
          {!isAuthenticated ? (
            <>
              <li><Link to="/login">Login</Link></li>
              <li><Link to="/register">Register</Link></li>
            </>
          ) : (
            <li><button onClick={handleLogout} style={{ background: "red", color: "white" }}>Logout</button></li>
          )}
        </ul>
      </div>
    </nav>
  );
};

export default Navbar;