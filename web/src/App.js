import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Feed from './pages/Feed';
import MySuggestions from './pages/MySuggestions';
import UserProfile from './pages/UserProfile';
import Navbar from './components/NavBar/NavBar';
import './App.css';
import UserBar from './components/UserBar/UserBar';
import LikedSuggestions from './pages/LikedSuggestions';
import CreateSuggestion from './pages/CreateSuggestion';

export default function App() {
  return (
    <Router>
      <div className="container">
        <UserBar />
        
        {/* Маршруты */}
        <Routes>
          <Route path="*" element={<Feed />} /> {/* По умолчанию открываем Ленту */}
          <Route path="/feed" element={<Feed />} /> {/* Лента */}

          <Route path="/my-suggestions" element={<MySuggestions />} /> {/* Мои предложения */}
          <Route path="/create-suggestion" element={<CreateSuggestion />} /> {/* Создание предложения */}
          <Route path="/my-likes" element={<LikedSuggestions />} /> {/* Мои лайки */}

          <Route path="/profile" element={<UserProfile />} /> {/* Профиль */}
        </Routes>

        <Navbar />
      </div>
    </Router>
  );
}