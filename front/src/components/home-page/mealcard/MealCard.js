import React from 'react';
import './MealCard.css';

const MealCard = ({ title, date }) => {
  return (
    <div className="meal-card">
      <img src="/food.jpg" alt={`${title}'s meal`} className="meal-image" />
      <h2>{title}</h2>
      <p>Дата: {date.toLocaleDateString()}</p>
      <p>Информация о блюде для {title}</p>
    </div>
  );
};

export default MealCard;
