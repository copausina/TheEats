import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import RestaurantCard from "../components/RestaurantCard";
import "./RestaurantIndex.css";
import { publicApi, authApi} from "../api/axiosInstance";
import RestaurantFormModal from "../components/RestaurantFormModal";

const RestaurantIndex = () => {
  const [restaurants, setRestaurants] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [formData, setFormData] = useState({
    name: "",
    address: "",
    cuisine: "",
    rating: "",
    image: null, // For file input
  });

  const toggleModal = () => {
    setShowModal(!showModal);
  };

  const handleModalClose = () => {
    setShowModal(false);
  };

  // Handle form changes
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  // Handle file upload
  const handleFileChange = (e) => {
    setFormData({ ...formData, image: e.target.files[0] });
  };

  // Handle form submission
  const handleSubmit = async (formData) => {
    // e.preventDefault(); // Handled in RestaurantFormModal
  
    // Prepare form data
    const data = new FormData();
    data.append("name", formData.name);
    data.append("address", formData.address);
    data.append("cuisine", formData.cuisine);
    data.append("rating", formData.rating);
    if (formData.image) {
      data.append("image", formData.image);
    }
  
    try {
      const response = await authApi.post("/api/restaurants/", data, { 
        headers: { "Content-Type": "multipart/form-data" },
      });
  
      console.log("Restaurant added:", response.data);
  
      // Close modal & reset form
      setShowModal(false);
      setFormData({ name: "", address: "", cuisine: "", rating: "", image: null });
  
      // Optionally refresh restaurant list without full page reload
    //   window.location.reload(); // TODO: Consider replacing with state update instead
    } catch (error) {
      console.error("Error adding restaurant:", error.response?.data || error.message);
    }
  };

  useEffect(() => {
    // Fetch restaurants from backend
    publicApi
      .get("/api/restaurants/") // The trailing backslash is important!
      .then((response) => setRestaurants(response.data))
      .catch((error) => console.error("Error fetching restaurants:", error));
  }, []);

  return (
    <div className="restaurant-index">
      <div className="header">
        <h2>Restaurants</h2>
        <button className="add-button" onClick={toggleModal}>Add Restaurant</button>
      </div>
      <div className="restaurant-grid">
      {restaurants.map((restaurant) => (
          <RestaurantCard
            key={restaurant.ID}  
            id={restaurant.ID}
            name={restaurant.name}
            imageurl={restaurant.imageurl}
            cuisine={restaurant.cuisine}
            address={restaurant.address}
            rating={restaurant.rating}
          />
        )
      )}
      </div>

      {showModal && (
        <RestaurantFormModal
            show={showModal}
            onClose={toggleModal}
            isEdit={false}
            onSubmit={handleSubmit}/>
       )}

      {/* Add restaurant pop-up */}
      {/* {showModal && (
        <div className="modal-overlay">
          <div className="modal">
            <h2>Add a New Restaurant</h2>
            <form onSubmit={handleSubmit}>
              <label>Restaurant Name:</label>
              <input type="text" name="name" value={formData.name} onChange={handleChange} required />
              <label>Address:</label>
              <input type="text" name="address" value={formData.address} onChange={handleChange} required />
              <label>Cuisine:</label>
              <input type="text" name="cuisine" value={formData.cuisine} onChange={handleChange} required />
              <label>Rating (0-5):</label>
              <input type="number" name="rating" value={formData.rating} onChange={handleChange} min="0" max="5" step="0.1" required />
              <label>Upload Image:</label>
              <input type="file" accept="image/*" onChange={handleFileChange} />

              <button type="submit">Submit</button>
              <button type="button" onClick={toggleModal}>Cancel</button>
            </form>
          </div>
        </div>
      )} */}
    </div>
  );
};

export default RestaurantIndex;