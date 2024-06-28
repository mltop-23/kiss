// src/components/daymenu/DayMenu.js
import React from 'react';
import { Link } from 'react-router-dom';
import './DayMenu.css';

const DayMenu = ({ day, meals }) => {
  console.log('DayMenu meals:', meals); // Логирование данных

  return (
    <div className="day-menu">
      <h2>{day}</h2>
      <div className="meal-columns">
        <div className="meal-time">
          <h3>Завтрак</h3>
          <div className="meals">
            {meals.breakfast && meals.breakfast.map((meal, index) => (
              <Link to={`/recipe/${meal.id}`} className="meal" key={meal.id || index}>
                <img src={meal.src} alt={meal.name} className="meal-image" />
                <div className="meal-info">
                  <h4>{meal.name}</h4>
                  <p>{meal.description}</p>
                </div>
              </Link>
            ))}
          </div>
        </div>
        <div className="meal-time">
          <h3>Обед</h3>
          <div className="meals">
            {meals.lunch && meals.lunch.map((meal, index) => (
              <Link to={`/recipe/${meal.id}`} className="meal" key={meal.id || index}>
                <img src={meal.src} alt={meal.name} className="meal-image" />
                <div className="meal-info">
                  <h4>{meal.name}</h4>
                  <p>{meal.description}</p>
                </div>
              </Link>
            ))}
          </div>
        </div>
        <div className="meal-time">
          <h3>Ужин</h3>
          <div className="meals">
            {meals.dinner && meals.dinner.map((meal, index) => (
              <Link to={`/recipe/${meal.id}`} className="meal" key={meal.id || index}>
                <img src={meal.src} alt={meal.name} className="meal-image" />
                <div className="meal-info">
                  <h4>{meal.name}</h4>
                  <p>{meal.description}</p>
                </div>
              </Link>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default DayMenu;
