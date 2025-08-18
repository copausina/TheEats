import { Link } from "react-router-dom";
import "./RestaurantCard.css";

const RestaurantCard = ({ id, name, imageurl, cuisine, address, rating}) => { //these names are NOT case-sensitive
  return (
    <Link to={`/restaurants/${id}`} className="restaurant-card">
      <img src={imageurl} alt={name} />
      <div className="info">
        <h3>{name}</h3>
        <p>{cuisine}</p>
        <p>{address}</p>
        <p>⭐ {rating}</p>
      </div>
    </Link>
  );
};

export default RestaurantCard;