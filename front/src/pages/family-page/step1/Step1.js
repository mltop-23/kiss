import React from 'react';
import './Step1.css';

const Step1 = ({ dishes, handleDishSelection, selectedDish, handleMouseOver, handleMouseOut, handlePrevStep, handleNextStep }) => {
  return (
    <div className="step-container">
      

      <div className="step-header">
      <h2>Шаг 1: Выбор блюда</h2>
      </div>
      
      <div className="dishes-step1">
        {dishes.map(dish => (
          <div
            key={dish.id}
            onClick={() => handleDishSelection(dish)}
            onMouseOver={() => handleMouseOver(dish)}
            onMouseOut={handleMouseOut}
            className={`dish-step1 ${selectedDish && selectedDish.id === dish.id ? 'selected-step1' : ''}`}
          >
            <img src={dish.src} alt={dish.name} className="dish-image-step1" />
            {selectedDish && selectedDish.id === dish.id && (
              <div className="dish-description-step1">
                <h3>{dish.name}</h3>
                <p>{dish.description}</p>
              </div>
            )}
          </div>
        ))}
      </div>
      <div className="buttons-step">
        <button onClick={handlePrevStep}>Назад</button>
        <button onClick={handleNextStep}>Далее</button>
      </div>
    </div>
  );
};

export default Step1;
