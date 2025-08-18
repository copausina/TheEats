import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import RestaurantIndex from "./pages/RestaurantIndex";
import RestaurantPage from "./pages/RestaurantPage";
import Navbar from "./components/Navbar";
import './App.css'


const App = () => {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/restaurants" element={<RestaurantIndex/>} />
        <Route path="/restaurants/:id" element={<RestaurantPage />} />
      </Routes>
    </Router>
  );
};

export default App;