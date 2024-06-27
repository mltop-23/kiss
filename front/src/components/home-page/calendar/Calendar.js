import React, { useState } from 'react';
import './Calendar.css';

const Calendar = ({ onDateChange }) => {
  const [selectedDate, setSelectedDate] = useState(new Date());

  const handleDateClick = (date) => {
    setSelectedDate(date);
    onDateChange(date);
  };

  // Примерный вывод календаря
  const renderCalendar = () => {
    const daysInMonth = new Date(selectedDate.getFullYear(), selectedDate.getMonth() + 1, 0).getDate();
    const firstDayOfMonth = new Date(selectedDate.getFullYear(), selectedDate.getMonth(), 1).getDay();
    const calendarDays = [];

    for (let i = 0; i < firstDayOfMonth; i++) {
      calendarDays.push(<div key={`empty-${i}`} className="calendar-day empty"></div>);
    }

    for (let i = 1; i <= daysInMonth; i++) {
      calendarDays.push(
        <div
          key={i}
          className={`calendar-day ${selectedDate.getDate() === i ? 'selected' : ''}`}
          onClick={() => handleDateClick(new Date(selectedDate.getFullYear(), selectedDate.getMonth(), i))}
        >
          {i}
        </div>
      );
    }

    return calendarDays;
  };

  return (
    <div className="calendar">
      <div className="calendar-header">
        <button onClick={() => handleDateClick(new Date(selectedDate.getFullYear(), selectedDate.getMonth() - 1, 1))}>
          &lt;
        </button>
        <h2>{selectedDate.toLocaleDateString('ru-RU', { month: 'long', year: 'numeric' })}</h2>
        <button onClick={() => handleDateClick(new Date(selectedDate.getFullYear(), selectedDate.getMonth() + 1, 1))}>
          &gt;
        </button>
      </div>
      <div className="calendar-body">
        {renderCalendar()}
      </div>
    </div>
  );
};

export default Calendar;
