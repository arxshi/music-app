import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function Profile() {
    const [user, setUser] = useState(null);
    const [error, setError] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        const fetchUser = async () => {
            const token = localStorage.getItem("token");
            if (!token) {
                navigate("/login");
                return;
            }

            try {
                const response = await axios.get("http://localhost:8080/users", {
                    headers: { Authorization: `Bearer ${token}` },
                });
                setUser(response.data);
            } catch (err) {
                setError("Ошибка загрузки профиля");
            }
        };

        fetchUser();
    }, [navigate]);

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/login");
    };

    return (
        <div className="max-w-lg mx-auto mt-10 bg-white p-6 rounded-lg shadow-md">
            <h2 className="text-2xl font-bold mb-4">Профиль</h2>
            {error && <p className="text-red-500">{error}</p>}
            {user ? (
                <>
                    <p><strong>Email:</strong> {user.email}</p>
                    <button onClick={handleLogout} className="mt-4 bg-red-500 text-white px-4 py-2 rounded">
                        Выйти
                    </button>
                </>
            ) : (
                <p>Загрузка...</p>
            )}
        </div>
    );
}

export default Profile;
