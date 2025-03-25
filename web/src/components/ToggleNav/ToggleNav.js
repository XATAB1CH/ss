import React from "react";
import { useNavigate, useLocation } from "react-router-dom";
import "./ToggleNav.css";

const ToggleNav = () => {
  const navigate = useNavigate();
  const location = useLocation();

  const isMySuggestions = location.pathname === "/my-suggestions";

  return (
    <div className="toggle-nav">
      <button
        className={`toggle-btn ${isMySuggestions ? "active" : ""}`}
        onClick={() => navigate("/my-suggestions")}
      >
        Мои предложения
      </button>
      <button
        className={`toggle-btn ${!isMySuggestions ? "active" : ""}`}
        onClick={() => navigate("/my-likes")}
      >
        Лайки
      </button>
    </div>
  );
};

export default ToggleNav;