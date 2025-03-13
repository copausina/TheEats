import { Link, useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { logout } from "../api/auth"; 
import { publicApi, authApi } from "../api/axiosInstance"
import "./Navbar.css";

const Navbar = () => {
  const navigate = useNavigate();
  const [isAuthenticated, setIsAuthenticated] = useState(localStorage.getItem("isAuthenticated") === "true");

  useEffect(() => {
    // Function to check auth status
    const checkAuth = () => {
      setIsAuthenticated(localStorage.getItem("isAuthenticated") === "true");
    };
  
    checkAuth(); // Run once on mount
  
    // Listen for storage changes (i.e. login from Login.jsx page)
    window.addEventListener("storage", checkAuth);
  
    return () => window.removeEventListener("storage", checkAuth); // Cleanup
  }, []);

  const handleLogout = async () => {
    await logout();
    setIsAuthenticated(false); // Update state
    localStorage.setItem("isAuthenticated", "false"); // For conditional rendering
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