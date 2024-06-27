// src/components/weekmenu/WeekMenu.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import DayMenu from './daymenu/DayMenu';
import './WeekMenu.css';

const WeekMenu = () => {
  const [weekMeals, setWeekMeals] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:3001/api/week-menu')
      .then(response => {
        console.log('Week menu data:', response.data.weekMenu); // Логирование данных
        setWeekMeals(response.data.weekMenu);
      })
      .catch(error => {
        console.error('There was an error fetching the week menu!', error);
      });
  }, []);

  return (
    <div className="week-menu">
      {Object.keys(weekMeals).map((day, index) => (
        <DayMenu key={index} day={day} meals={weekMeals[day]} />
      ))}
    </div>
  );
};

export default WeekMenu;
