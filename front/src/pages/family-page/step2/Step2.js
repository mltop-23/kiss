import React from 'react';
import './Step2.css';

const Step2 = ({ handleCookSelection, selectedCook, handlePrevStep, handleNextStep }) => {
  return (
    <div className="step-container">
      
      <div className="step-header">
      <h2>Шаг 2: Кто будет готовить</h2>
      </div>

      <div className="cooks-step2">
        <div className="cook-container-step2" onClick={() => handleCookSelection('Мама')}>
          <img src="/mama.jpg" alt="Мама" className={`cook-step2 ${selectedCook === 'Мама' ? 'selected-step2' : ''}`} />
        </div>
        <div className="cook-container-step2" onClick={() => handleCookSelection('Папа')}>
          <img src="/papa.png" alt="Папа" className={`cook-step2 ${selectedCook === 'Папа' ? 'selected-step2' : ''}`} />
        </div>
      </div>
      <div className="buttons-step">
          <button onClick={handlePrevStep}>Назад</button>
          <button onClick={handleNextStep}>Далее</button>
      </div>
    </div>
  );
};

export default Step2;
