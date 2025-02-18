import { useState, useEffect } from "react";
import axios from "axios";

function Playlists() {
    const [playlists, setPlaylists] = useState([]);
    const [newPlaylist, setNewPlaylist] = useState("");
    const [error, setError] = useState("");

    useEffect(() => {
        const fetchPlaylists = async () => {
            try {
                const token = localStorage.getItem("token");
                const response = await axios.get("http://localhost:8080/playlists", {
                    headers: { Authorization: token },
                });
                if (response.data) {
                    setPlaylists(response.data);
                } else {
                    setPlaylists([])
                }
            } catch (err) {
                setError("Ошибка загрузки плейлистов");
            }
        };

        fetchPlaylists();
    }, []);

    const handleAddPlaylist = async () => {
        if (!newPlaylist) return;
        try {
            const token = localStorage.getItem("token");
            const response = await axios.post(
                "http://localhost:8080/playlists",
                { name: newPlaylist },
                { headers: { Authorization: token } }
            );
            setPlaylists([...playlists, response.data]);
            setNewPlaylist("");
        } catch (err) {
            setError("Ошибка создания плейлиста");
        }
    };

    return (
        <div className="max-w-lg mx-auto mt-10 bg-white p-6 rounded-lg shadow-md">
            <h2 className="text-2xl font-bold text-gray-800 mb-4">Мои плейлисты</h2>
            {error && <p className="text-red-500">{error}</p>}
            <input
                type="text"
                placeholder="Название плейлиста"
                value={newPlaylist}
                onChange={(e) => setNewPlaylist(e.target.value)}
                className="w-full p-2 border rounded mt-2"
            />
            <button onClick={handleAddPlaylist} className="w-full bg-blue-500 text-black py-2 mt-4 rounded">
                Добавить плейлист
            </button>
            <ul className="mt-4">
                {playlists.map((playlist) => (
                    <li key={playlist.Id} className="p-2 border-b">{playlist.Name}</li>
                ))}
            </ul>
        </div>
    );
}

export default Playlists;
