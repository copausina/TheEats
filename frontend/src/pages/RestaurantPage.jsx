import { useParams, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import { publicApi, authApi } from "../api/axiosInstance";
import "./RestaurantPage.css";
import RestaurantFormModal from "../components/RestaurantFormModal";

const RestaurantPage = () => {
    const { id } = useParams();
    const [restaurant, setRestaurant] = useState(null);
    const [showModal, setShowModal] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchRestaurant = async () => {
            try {
                const response = await publicApi.get(`/api/restaurants/${id}`);
                setRestaurant(response.data);
            } catch (err) {
                console.error("Error fetching restaurant", err);
            }
        };
        fetchRestaurant();
    }, [id]);

    const handleDelete = async () => {
        const confirmed = window.confirm("Are you sure you want to delete this restaurant?");
        if (!confirmed) return;

        try {
            await authApi.delete(`/api/restaurants/${id}`);
            navigate(-1, { replace: true }) || navigate("/restaurants");
        } catch (err) {
            console.error("Error deleting restaurant", err);
        }
    };

    const handleEdit = () => {
        setShowModal(true);
    };

    const handleModalClose = () => {
        setShowModal(false);
    };

    const handleFormSubmit = async (formData) => {
        try {
            // Prepare multipart/form-data if there's an image file
            const data = new FormData();
            data.append("name", formData.name);
            data.append("address", formData.address);
            data.append("cuisine", formData.cuisine);
            data.append("rating", formData.rating);
            if (formData.image) {
            data.append("image", formData.image);
            }

            const response = await authApi.put(`/api/restaurants/${id}`, data, {
            headers: { "Content-Type": "multipart/form-data" },
            });

            setRestaurant(response.data);
            setShowModal(false);
        } catch (err) {
            console.error("Error updating restaurant", err);
        }
    };

    if (!restaurant) return <div>Loading...</div>;

    return (
        <div className="restaurant-detail">
        <h2>{restaurant.name}</h2>
        <img src={restaurant.imageurl} alt={restaurant.name} />
        <p>Cuisine: {restaurant.cuisine}</p>
        <p>Location: {restaurant.address}</p>
        <p>Rating: ⭐ {restaurant.rating}</p>

        <button onClick={handleEdit}>Edit</button>
        <button onClick={handleDelete}>Delete</button>

        <hr />
        <h3>Reviews</h3>
        <p>Reviews not implemented yet</p>

        {showModal && (
            <RestaurantFormModal
                show={showModal}
                onClose={handleModalClose}
                initialData={restaurant}
                isEdit={true}
                onSubmit={handleFormSubmit}
            />
        )}
        </div>

        
    );
};

export default RestaurantPage;