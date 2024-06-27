import React from 'react';
import Calendar from '../../components/home-page/calendar/Calendar';
import MealCard from '../../components/home-page/mealcard/MealCard';
import Gallery from '../../components/home-page/gallery/Gallery';
import './HomePage.css';

const HomePage = ({ selectedDate, onDateChange }) => {
  return (
    <>
      <div className="main-content">
        <div className="meal-cards">
          <MealCard title="Мама" date={selectedDate} className="meal-card" />
          <MealCard title="Папа" date={selectedDate} className="meal-card" />
        </div>
        <Calendar onDateChange={onDateChange} className="calendar" />
      </div>
      <div className="gallery-content">
        <Gallery />
      </div>
    </>
  );
};

export default HomePage;
