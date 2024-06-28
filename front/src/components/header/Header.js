import React, { useState } from 'react';
import { useMediaQuery } from 'react-responsive';
import { FaBars } from 'react-icons/fa';
import { Link, useLocation } from 'react-router-dom';
import './Header.css';

const Header = () => {
  const isMobile = useMediaQuery({ query: '(max-width: 1000px)' });
  const [menuOpen, setMenuOpen] = useState(false);
  const location = useLocation();
  

  const toggleMenu = () => {
    setMenuOpen(!menuOpen);
  };

  return (
    <header className="header">
      <div className="nav-container">
        <div className="logo">
          <Link to="/">
            <img src="/logo.png" alt="Logo" className="logo-image" />
          </Link>
        </div>
        {isMobile ? (
          <>
            <FaBars className="burger-icon" onClick={toggleMenu} />
            {menuOpen && (
              <nav className="mobile-menu">
                <Link to="/menu">Меню</Link>
                <Link to="/recipes">Рецепты</Link>
                <Link to="/family">Семья</Link>
                <Link to="/auth" className="button">Войти</Link>
                <Link to="/register" className="button">Зарегистрироваться</Link>
              </nav>
            )}
          </>
        ) : (
          <>
            <nav className="nav-links">
              <Link to="/menu">Меню</Link>
              <Link to="/recipes">Рецепты</Link>
              <Link to="/family">Семья</Link>
            </nav>
            <Link to="/" className="title">
              {location.pathname === '/menu' ? 'Меню' : 
              location.pathname === '/recipes' ? 'Рецепты' : 
              location.pathname === '/auth' ? 'Заходи, не стесняйся' : 
              location.pathname === '/family' ? 'Ваша семья' : 
              'A Cooking Schedule'}
            </Link>
          </>
        )}
      </div>
      {!isMobile && (
        <div className="buttons">
          <Link to="/auth" className="button">Войти</Link>
          <Link to="/register" className="button">Зарегистрироваться</Link>
        </div>
      )}
    </header>
  );
};

export default Header;