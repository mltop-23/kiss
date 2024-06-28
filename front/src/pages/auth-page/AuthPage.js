import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './AuthPage.css';

const AuthPage = () => {
  const [isLogin, setIsLogin] = useState(true);

  const toggleForm = () => {
    setIsLogin(!isLogin);
  };

  const generatePassword = () => {
    const chars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    let password = '';
    for (let i = 0; i < 12; i++) {
      password += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return password;
  };

  return (
    <div className="auth-container">
      {isLogin ? (
        <div className="auth-form">
          <h2 className="auth-title">Вход</h2>
          <label>Номер телефона или email</label>
          <input type="text" placeholder="Введите номер телефона или email" />
          <label>Пароль</label>
          <input type="password" placeholder="Введите пароль" />
          <button className="auth-button">Войти</button>
          <p className="auth-toggle">
            Если вы не зарегистрированы - <span onClick={toggleForm} className="auth-link">зарегистрироваться</span>
          </p>
        </div>
      ) : (
        <div className="auth-form">
          <h2 className="auth-title">Регистрация</h2>
          <label>Имя</label>
          <input type="text" placeholder="Введите имя" />
          <label>Фамилия</label>
          <input type="text" placeholder="Введите фамилию" />
          <label>Почта</label>
          <input type="email" placeholder="Введите почту" />
          <label>Пол</label>
          <div className="gender-selection">
            <label>
              <input type="radio" name="gender" value="male" /> Мужчина
            </label>
            <label>
              <input type="radio" name="gender" value="female" /> Женщина
            </label>
          </div>
          <label>Пароль</label>
          <input type="password" placeholder="Введите пароль" />
          <button onClick={() => alert(`Сгенерированный пароль: ${generatePassword()}`)} className="generate-button">Сгенерировать пароль</button>
          <label>Подтверждение пароля</label>
          <input type="password" placeholder="Подтвердите пароль" />
          <label>ID семьи</label>
          <input type="text" placeholder="Введите ID семьи" />
          <button className="auth-button">Зарегистрироваться</button>
          <p className="auth-toggle">
            Уже зарегистрированы? <span onClick={toggleForm} className="auth-link">Войти</span>
          </p>
        </div>
      )}
    </div>
  );
};

export default AuthPage;
