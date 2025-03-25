import React, { useEffect, useState } from "react";
import ToggleNav from "../components/ToggleNav/ToggleNav";
import "./LikedSuggestions.css";
import telegramIcon from "../assets/icons/telegram.svg";
import avatar from "../assets/images/avatar.png";

const LikedSuggestions = () => {
  const [likedSuggestions, setLikedSuggestions] = useState([]);

  // Заглушка: данные вместо бэкенда
  useEffect(() => {
    setLikedSuggestions([
      { id: 1, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
      { id: 2, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
      { id: 3, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
      { id: 4, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
      { id: 5, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
      { id: 6, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
      { id: 7, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
      { id: 8, user: "Артём Шарапов, 23 года, г. Москва", text: "Пинать хуи сегодня вечером" },
    ]);
  }, []);

  return (
    <div>

      <div className="container-top">
        <ToggleNav />
      </div>
      <div className="liked-container">
        <div className="liked-grid">
            {likedSuggestions.map((suggestion) => (
            <div key={suggestion.id} className="liked-card">
                <img src={avatar} alt="User" className="avatar" />
                <div className="liked-info">
                <p className="liked-user">{suggestion.user}</p>
                <p className="liked-text">Предложение: {suggestion.text}</p>
                </div>
                <button className="send-icon">
                    <img src={telegramIcon} alt="Send" className="send-icon-img" />
                </button>
            </div>
            ))}
            </div>
        </div>
    </div>
  );
};

export default LikedSuggestions;
