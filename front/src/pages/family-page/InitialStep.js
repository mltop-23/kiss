// src/components/InitialStep.js
import React from 'react';
import './InitialStep.css';

const InitialStep = ({ selectedDate, meals, handleNextStep, handleDeleteDish }) => {
  const formatDate = (date) => {
    const today = new Date();
    if (date.toDateString() === today.toDateString()) {
      return 'Блюда на сегодня';
    }
    return `Блюда на ${date.toLocaleDateString('ru-RU')}`;
  };
  
  return (
    <div className="initial-step">
      <h2 className="step-header">{formatDate(selectedDate)}</h2>
      <div className="step-meals">
        <div className="step-meal">
          <div className="step-meal-header">Завтрак:</div>
          {meals.breakfast ? (
            <>
              <img src={meals.breakfast.dish.src} alt={meals.breakfast.dish.name} className="step-meal-image" />
              <div className="step-meal-info">
                <div>{meals.breakfast.dish.name}</div>
                <div>({meals.breakfast.cook})</div>
              </div>
            </>
          ) : <div className="step-no-meal">...</div>}
        </div>
        <div className="step-meal">
          <div className="step-meal-header">Обед:</div>
          {meals.lunch ? (
            <>
              <img src={meals.lunch.dish.src} alt={meals.lunch.dish.name} className="step-meal-image" />
              <div className="step-meal-info">
                <div>{meals.lunch.dish.name}</div>
                <div>({meals.lunch.cook})</div>
              </div>
            </>
          ) : <div className="step-no-meal">...</div>}
        </div>
        <div className="step-meal">
          <div className="step-meal-header">Ужин:</div>
          {meals.dinner ? (
            <>
              <img src={meals.dinner.dish.src} alt={meals.dinner.dish.name} className="step-meal-image" />
              <div className="step-meal-info">
                <div>{meals.dinner.dish.name}</div>
                <div>({meals.dinner.cook})</div>
              </div>
            </>
          ) : <div className="step-no-meal">...</div>}
        </div>
      </div>
      <div className="buttons-step">
        <button onClick={handleDeleteDish} className="step-button">Удалить</button>
        <button onClick={handleNextStep} className="step-button">Добавить</button>
      </div>
    </div>
  );
};

export default InitialStep;
