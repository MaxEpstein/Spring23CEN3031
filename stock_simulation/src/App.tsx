import React from 'react';
import { BrowserRouter, Route, Switch } from "react-router-dom";
import './App.css';

import { AboutPage } from './pages/AboutPage';
import { ContactPage } from './pages/ContactPage';
import { HomePage } from './pages/HomePage';
import { NavBar } from './NavBar';
import { LogInPage } from './pages/LogInPage';

export function App() {
  return (
    <>
      <BrowserRouter>
        <NavBar />
        <Switch>
          <Route path="/" exact><HomePage /></Route>
          <Route path="/about"><AboutPage /></Route>
          <Route path="/contact"><ContactPage /></Route>
          <Route path="/login"><LogInPage /></Route>
        </Switch>
      </BrowserRouter>
    </>
  );
}

export default App;
