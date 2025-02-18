import { useState, useEffect } from "react";
import axios from "axios";
import {data} from "react-router-dom";

function Tracks() {
    const [tracks, setTracks] = useState([]);
    const [selectedFile, setSelectedFile] = useState(null);
    const [error, setError] = useState("");

    useEffect(() => {
        const fetchTracks = async () => {
            try {
                const response = await axios.get("http://localhost:8080/tracks");
                if (response.data) {
                    setTracks(response.data);
                } else {
                    setTracks([])
                }
            } catch (err) {
                setError("Ошибка загрузки треков");
            }
        };

        fetchTracks();
    }, []);

    const handleFileChange = (event) => {
        setSelectedFile(event.target.files[0]);
    };

    const handleUpload = async () => {
        if (!selectedFile) return;

        const formData = new FormData();
        formData.append("file", selectedFile);
        formData.append("metadata", JSON.stringify({
            title: selectedFile.name.split("-")[1].split(".")[0],
            artist: selectedFile.name.split("-")[0]
        }));

        try {
            await axios.post("http://localhost:8080/upload", formData, {
                headers: { "Content-Type": "multipart/form-data" },
            });
            setSelectedFile(null);
            window.location.reload(); // Перезагрузка списка треков
        } catch (err) {
            setError("Ошибка загрузки файла");
        }
    };

    return (
        <div className="max-w-lg mx-auto mt-10 bg-white p-6 rounded-lg shadow-md">
            <h2 className="text-2xl font-bold text-gray-800 mb-4">Музыкальные треки</h2>
            {error && <p className="text-red-500">{error}</p>}
            <input type="file" onChange={handleFileChange} className="w-full p-2 border rounded mt-2" />
            <button onClick={handleUpload} className="w-full bg-blue-500 text-black py-2 mt-4 rounded">
                Загрузить трек
            </button>
            <ul className="mt-4 text-black">
                {tracks.map((track) => (
                    <li key={track.Id} className="p-2 border-b">
                        <p>{track.Artist}- {track.Title}</p>
                        <audio controls className="w-full mt-2">
                            <source src={`http://localhost:8080/uploads/${track.FileName}`} type="audio/mpeg" />
                        </audio>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Tracks;
