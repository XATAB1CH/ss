import React, { useState, useEffect } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import "./NavBar.css";
import FeedIcon from '../../assets/icons/feed.svg';
import ProfileIcon from '../../assets/icons/profile.svg';
import SuggestionsIcon from '../../assets/icons/suggestions.svg';

const NavBar = () => {
  const navigate = useNavigate();
  const location = useLocation();
  const [activePath, setActivePath] = useState(location.pathname); // Состояние для активного пути

  // Обновляем активный путь при изменении location
  useEffect(() => {
    setActivePath(location.pathname);
  }, [location]);

  return (
    <div className="navbar">
      <button
        onClick={() => navigate("/my-suggestions")}
        className={activePath === "/my-suggestions" ? "active" : "inactive"}
      >
        <img src={SuggestionsIcon} alt="Мои предложения" />
        <span>Мои предложения</span>
      </button>
      <button
        onClick={() => navigate("/feed")}
        className={activePath === "/feed" ? "active" : "inactive"}
      >
        <img src={FeedIcon} alt="Лента" />
        <span>Лента</span>
      </button>
      <button
        onClick={() => navigate("/profile")}
        className={activePath === "/profile" ? "active" : "inactive"}
      >
        <img src={ProfileIcon} alt="Профиль" />
        <span>Профиль</span>
      </button>
    </div>
  );
};

export default NavBar;