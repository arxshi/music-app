import { Link } from "react-router-dom";
import { useEffect, useState } from "react";

function Navbar() {
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        const token = localStorage.getItem("token");
        setIsAuthenticated(!!token); // Проверяемa наличие токена
    }, []);

    return (
        <nav className="bg-white shadow-md p-4 flex justify-center gap-6">
            <Link to="/" className="text-gray-700 font-semibold hover:text-blue-500">Главная</Link>
            <Link to="/tracks" className="text-gray-700 font-semibold hover:text-blue-500">Треки</Link>
            <Link to="/playlists" className="text-gray-700 font-semibold hover:text-blue-500">Плейлисты</Link>

            {!isAuthenticated ? (
                <>
                    <Link to="/register" className="text-gray-700 font-semibold hover:text-blue-500">Регистрация</Link>
                    <Link to="/login" className="text-gray-700 font-semibold hover:text-blue-500">Вход</Link>
                </>
            ) : (
                <Link to="/profile" className="text-gray-700 font-semibold hover:text-blue-500">Профиль</Link>
            )}
        </nav>
    );
}

export default Navbar;
