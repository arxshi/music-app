import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post("http://localhost:8080/auth/login", { email, password });
            localStorage.setItem("token", response.data.token);
            navigate("/profile");
        } catch (err) {
            setError(err.response?.data?.error || "Ошибка входа. Проверьте данные.");
        }
    };

    return (
        <div className="max-w-md mx-auto mt-10 bg-white p-6 rounded-lg shadow-md">
            <h2 className="text-2xl font-bold mb-4">Вход</h2>
            {error && <p className="text-red-500">{error}</p>}
            <form onSubmit={handleLogin}>
                <input type="email" placeholder="Email" className="w-full p-2 border rounded mt-2"
                       value={email} onChange={(e) => setEmail(e.target.value)} />
                <input type="password" placeholder="Пароль" className="w-full p-2 border rounded mt-2"
                       value={password} onChange={(e) => setPassword(e.target.value)} />
                <button type="submit" className="w-full bg-blue-500 text-white py-2 mt-4 rounded">
                    Войти
                </button>
            </form>
        </div>
    );
}

export default Login;
