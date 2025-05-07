import "./RestaurantCard.css";

const RestaurantCard = ({ name, imageurl, cuisine, address, rating}) => { //these names are NOT case-sensitive
  return (
    <div className="restaurant-card">
      <img src={imageurl} alt={name} />
      <div className="info">
        <h3>{name}</h3>
        <p>{cuisine}</p>
        <p>{address}</p>
        <p>⭐{rating}</p>
      </div>
    </div>
  );
};

export default RestaurantCard;