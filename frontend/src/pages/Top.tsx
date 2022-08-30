import React from 'react';
import './Top.css';
import { Link } from 'react-router-dom';

function Top() {
  return (
    <>
      <p>Hello</p>
      <nav>
        <Link to='bucket-items'>Items</Link>
      </nav>
    </>
  );
}

export default Top;
