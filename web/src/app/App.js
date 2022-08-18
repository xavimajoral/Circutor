import React, { useContext, createContext, useState } from "react";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";

import LayoutRegister from '../layouts/register';
import LayoutLogin from '../layouts/login';
import BuildingList from '../layouts/buildings_list';
import BuildingMetrics from '../layouts/buildings_metrics'
import BookmarksList from "../layouts/bookmarks_list";


function App() {
  return (
    <BrowserRouter basename={"/"}>
      <Routes>
        <Route path="/bookmarks" element={<BookmarksList />} />
        <Route path="/buildings" element={<BuildingList />} />
        <Route path="/buildings/:buildingId" element={<BuildingMetrics />} />
        <Route path="/register" element={<LayoutRegister />} />
        <Route path="home" element={<LayoutRegister />} />
        <Route path="/login" element={<LayoutLogin />} />
        <Route path="/" element={<LayoutRegister />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App
