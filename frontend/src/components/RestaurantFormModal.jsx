import React, { useEffect, useState } from "react";

const RestaurantFormModal = ({ show, onClose, initialData, isEdit, onSubmit }) => {
    const [formData, setFormData] = useState({
    name: "",
    address: "",
    cuisine: "",
    rating: "",
    image: null,
});

useEffect(() => {
  if (isEdit && initialData) {
    setFormData({
      name: initialData.name,
      address: initialData.address,
      cuisine: initialData.cuisine,
      rating: initialData.rating,
      image: null, // image upload optional
    });
  }
}, [initialData, isEdit]);

const handleChange = (e) => {
  const { name, value } = e.target;
  setFormData((prev) => ({ ...prev, [name]: value }));
};

const handleFileChange = (e) => {
  setFormData((prev) => ({ ...prev, image: e.target.files[0] }));
};

const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData); // Parent handles API call
};

if (!show) return null;

  return (
    <div className="modal-overlay">
      <div className="modal">
        <h2>{isEdit ? "Edit Restaurant" : "Add a New Restaurant"}</h2>
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
          <button type="button" onClick={onClose}>Cancel</button>
        </form>
      </div>
    </div>
  );
};

export default RestaurantFormModal;