import { NavLink } from "react-router-dom";
import React from 'react';
import {BsFillBarChartFill} from "react-icons/bs";
import "./styles.css";

import * as data from './navbar.json'
import {HomePage} from "./pages/HomePage";
const linksString = JSON.stringify(data);
const links = JSON.parse(linksString).links;


type Link = {
  label: string;
  href: string;
}

const Links: React.FC<{links: Link[]}> = ({ links }) => {
  return (
      <div className='links-container'>
      
          {links.map((link: Link) => {
              return (
                  <div key={link.href} className='link'>
                      <a href={link.href}>
                          {link.label}
                      </a>
                  </div>
              )
          })}
      </div>
  )
}

export function NavBar() {

  return (
    <nav className="navbar-container">
      <div className='logo-container'>
          <NavLink to={'/'}>
         <span style={{font: 'roboto', }}>
             {" "}
            <BsFillBarChartFill />  Mind My Wallet  {""}
         </span>
      </NavLink>
      </div>
        <Links links={links}/>
    </nav>
  );
}


