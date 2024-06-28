import React from 'react';
import './Step3.css';

const Step3 = ({ handleMealSelection, selectedMeal, handlePrevStep, handleAddDish }) => {
  return (
    <div className="step-container">

      <div className="step-header">
      <h2>Шаг 3: Выбор приема пищи</h2>
      </div>

      <div className="meals-step3">
        <div
          onClick={() => handleMealSelection('breakfast')}
          className={`meal-step3 ${selectedMeal === 'breakfast' ? 'selected-step3' : ''}`}
        >
          Завтрак
        </div>
        <div
          onClick={() => handleMealSelection('lunch')}
          className={`meal-step3 ${selectedMeal === 'lunch' ? 'selected-step3' : ''}`}
        >
          Обед
        </div>
        <div
          onClick={() => handleMealSelection('dinner')}
          className={`meal-step3 ${selectedMeal === 'dinner' ? 'selected-step3' : ''}`}
        >
          Ужин
        </div>
      </div>
      <div className="buttons-step">
        <button onClick={handlePrevStep}>Назад</button>
        <button onClick={handleAddDish}>Добавить блюдо</button>
      </div>
    </div>
  );
};

export default Step3;
