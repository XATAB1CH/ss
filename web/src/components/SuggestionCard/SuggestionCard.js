import React from "react";
import "./SuggestionCard.css";
import LikeIcon from '../../assets/icons/like.svg'; 

const SuggestionCard = ({ title, image, description, category, isActive, likes }) => {
  return (
    <div className="suggestion-card">
      <h3 className="suggestion-title">{title}</h3>
      <img src={image} alt={title} className="suggestion-image" />
      <p className="suggestion-description">{description}</p>

      <div className="suggestion-footer">
        <span className="suggestion-activity">{isActive}</span>
        <span className="suggestion-category">{category}</span>

        <div className="suggestion-likes">
          <img src={LikeIcon} alt="Лайк" />
          <span> 76 </span>
        </div>
      </div>
    </div>
  );
};

export default SuggestionCard;