import React from 'react';
import './UserBar.css';
import avatar from "../../assets/images/avatar.png";

function UserBar() {
  return (
    <div className="user-bar">
      <img 
        src={avatar} // Замените на реальную ссылку на фото
        alt="User Avatar" 
        className="user-avatar"
      />
      <span className="user-name">Артём Шарапов</span>
    </div>
  );
}

export default UserBar;