import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();
  
  return(
    <div>  
      <h1>Welcome to TheEats!</h1>
      <button onClick={() => navigate("/login")}>Login</button>
      <button onClick={() => navigate("/register")}>Register</button>
    </div>
  );

  
}
  export default Home;