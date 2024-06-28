// src/pages/menupage/MenuPage.js
import React from 'react';
import WeekMenu from '../../components/menu-page/weekmenu/WeekMenu';
import './MenuPage.css';

const MenuPage = () => {
  return (
    <div className="menu-page">

      <WeekMenu />
    </div>
  );
};

export default MenuPage;
