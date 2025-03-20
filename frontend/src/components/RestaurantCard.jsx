import "./RestaurantCard.css";

const RestaurantCard = ({ name, imageUrl, cuisine, location, rating}) => {
  return (
    <div className="restaurant-card">
      <img src={imageUrl} alt={name} />
      <div className="info">
        <h3>{name}</h3>
        <p>{cuisine}</p>
        <p>{location}</p>
        <p>⭐{rating}</p>
      </div>
    </div>
  );
};

export default RestaurantCard;