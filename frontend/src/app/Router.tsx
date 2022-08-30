import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import BucketItems from 'pages/BucketItems';
import Top from '../pages/Top';

const Router = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Top />} />
        <Route path='bucket-items' element={<BucketItems />} />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
