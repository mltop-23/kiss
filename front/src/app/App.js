import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Header from '../components/header/Header';
import AuthPage from '../pages/auth-page/AuthPage';
import HomePage from '../pages/home-page/HomePage';
import RecipesPage from '../pages/recipes-page/RecipesPage';
import MenuPage from '../pages/menu-page/MenuPage';
import Recipe from '../components/1-recipe-page/Recipe';
import FamilyPage from '../pages/family-page/FamilyPage';
import './App.css';

const App = () => {
  const [selectedDate, setSelectedDate] = useState(new Date());

  const handleDateChange = (date) => {
    setSelectedDate(date);
  };

  return (
    <Router>
      <div className="App">
        <Header />
        <div className="content">
          <div className="container">
              <Routes>
              <Route path="/" element={<HomePage selectedDate={selectedDate} onDateChange={handleDateChange} />} />
              <Route path="/auth" element={<AuthPage />} />
              <Route path="/recipe/:id" element={<Recipe />} />
              <Route path="/menu" element={<MenuPage />} />
              <Route path="/recipes" element={<RecipesPage />} />
              <Route path="/family" element={<FamilyPage selectedDate={selectedDate} onDateChange={handleDateChange} />} />
              <Route path="/recipe/:id" component={Recipe} />
              </Routes>
          </div>
        </div>
      </div>
    </Router>
  );
};

export default App;
