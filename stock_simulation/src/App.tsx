import React from 'react';
import { BrowserRouter, Route, Switch } from "react-router-dom";

import { AboutPage } from './pages/AboutPage';
import { ContactPage } from './pages/ContactPage';
import { HomePage } from './pages/HomePage';
import { NavBar } from './NavBar';
import { Login } from './pages/LogInPage';
import { Footer } from './Footer';
import { Search } from './pages/Search';


export function App() {
  return (
    <>
    <div className="body">
    <BrowserRouter>
        <NavBar />
        <Switch>
          <Route path="/" exact><HomePage /></Route>
          <Route path="/about"><AboutPage /></Route>
          <Route path="/contact"><ContactPage /></Route>
          <Route path="/search"><Search /></Route>
          <Route path="/login"><Login /></Route>
        </Switch>
      </BrowserRouter>

        <Footer />
    </div>
      

    </>
  );
}

export default App;
