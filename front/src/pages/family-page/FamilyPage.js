import React, { useState, useEffect } from 'react';
import axios from 'axios';
// import Calendar from 'react-calendar';
// import 'react-calendar/dist/Calendar.css';
import 'primereact/resources/themes/saga-blue/theme.css'; // Theme
import 'primereact/resources/primereact.min.css';         // Core CSS
import 'primeicons/primeicons.css';                      // Icons
import { Timeline } from 'primereact/timeline';
import Step1 from './step1/Step1';
import Step2 from './step2/Step2';
import Step3 from './step3/Step3';
import InitialStep from './InitialStep';
import Calendar from '../../components/home-page/calendar/Calendar';
import './FamilyPage.css';

const FamilyPage = () => {
  const [selectedDate, setSelectedDate] = useState(new Date());
  const [currentStep, setCurrentStep] = useState(0);
  const [selectedDish, setSelectedDish] = useState(null);
  const [hoveredDish, setHoveredDish] = useState(null);
  const [selectedCook, setSelectedCook] = useState(null);
  const [selectedMeal, setSelectedMeal] = useState(null);
  const [meals, setMeals] = useState({ breakfast: null, lunch: null, dinner: null });
  const [dishes, setDishes] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:3001/api/recipes')
      .then(response => {
        setDishes(response.data.recipes);
      })
      .catch(error => {
        console.error('There was an error fetching the recipes!', error);
      });
  }, []);

  const handleNextStep = () => {
    setCurrentStep((prevStep) => prevStep + 1);
  };

  const handlePrevStep = () => {
    setCurrentStep((prevStep) => prevStep - 1);
  };

  const handleDishSelection = (dish) => {
    setSelectedDish(dish);
  };

  const handleCookSelection = (cook) => {
    setSelectedCook(cook);
  };

  const handleMealSelection = (meal) => {
    setSelectedMeal(meal);
  };

  const handleAddDish = () => {
    setMeals((prevMeals) => ({
      ...prevMeals,
      [selectedMeal]: { dish: selectedDish, cook: selectedCook },
    }));
    setSelectedDish(null);
    setSelectedCook(null);
    setSelectedMeal(null);
  };

  const handleDeleteDish = () => {
    setMeals({ breakfast: null, lunch: null, dinner: null });
  };

  const handleDateChange = (date) => {
    setSelectedDate(date);
  };

  const handleMouseOver = (dish) => {
    setHoveredDish(dish);
  };

  const handleMouseOut = () => {
    setHoveredDish(null);
  };

  const events = [
    { status: 'Блюда на сегодня', content: 
      <InitialStep
        selectedDate={selectedDate}
        meals={meals}
        handleNextStep={handleNextStep}
        handleDeleteDish={handleDeleteDish}
      /> 
    },
    { status: 'Шаг 1: Выбор блюда', content: 
      <Step1
        handlePrevStep={handlePrevStep}
        handleNextStep={handleNextStep}
        selectedDish={selectedDish}
        hoveredDish={hoveredDish}
        handleMouseOver={handleMouseOver}
        handleMouseOut={handleMouseOut}
        handleDishSelection={handleDishSelection}
        dishes={dishes}
      />
    },
    { status: 'Шаг 2: Кто будет готовить', content: 
      <Step2
        handlePrevStep={handlePrevStep}
        handleNextStep={handleNextStep}
        selectedCook={selectedCook}
        handleCookSelection={handleCookSelection}
      />
    },
    { status: 'Шаг 3: Выбор приема пищи', content: 
      <Step3
        handlePrevStep={handlePrevStep}
        handleAddDish={handleAddDish}
        selectedMeal={selectedMeal}
        handleMealSelection={handleMealSelection}
      />
    }
  ];

  return (
    <div className="family-page">
      <div className="calendar-container">
        <Calendar 
          onChange={handleDateChange} 
          value={selectedDate} 
          tileClassName={({ date, view }) =>
            date.toDateString() === selectedDate.toDateString() ? 'selected-date' : ''
          }
        />
      </div>
      <div className="form-container">
        <Timeline value={events} content={(item) => item.content} />
      </div>
    </div>
  );
};

export default FamilyPage;
