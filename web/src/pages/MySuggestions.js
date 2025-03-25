import React from "react";
import { useNavigate, useLocation } from "react-router-dom";
import ToggleNav from "../components/ToggleNav/ToggleNav.js";
import SuggestionCard from "../components/SuggestionCard/SuggestionCard";
import './MySuggestions.css';
import img1 from "../assets/images/img1.png";
import img2 from "../assets/images/img2.png";
import img3 from "../assets/images/img3.png";

const suggestions = [
  {
    title: "Попинать хуи сегодня вечером... Попинать хуи сегодня вечером...",
    image: img1, 
    description: "Сегодня я предлагаю вообще ничего не делать, просто вместе ничего...",
    category: "Культура и искусство",
    isActive: "Активно",
    likes: 69
  },
  {
    title: "Второе предложение",
    image: img2,
    description: "Описание второго предложения...",
    category: "Активности",
    isActive: "Активно",
    likes: 42
  },
  {
    title: "Третье предложение",
    image: img3,
    description: "Описание третьего предложения...",
    category: "Активности",
    isActive: "Активно",
    likes: 42
  },
];

const MySuggestions = () => {
  const navigate = useNavigate();

  return (
    <div>
      <div className="container-top">
        <ToggleNav />
      </div>
      <div className="container">
        <div className="suggestions-container">
          {suggestions.map((suggestion, index) => (
            <SuggestionCard key={index} {...suggestion} />
          ))}
        </div>
        
        {/* Кнопка "Создать" */}
        <button className="create-suggestion-button" onClick={() => navigate("/create-suggestion")}>
          Создать
        </button>
      </div>
    </div>
    
  );
};

export default MySuggestions;