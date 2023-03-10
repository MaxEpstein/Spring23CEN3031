import { NavLink } from "react-router-dom";
import React from 'react';
import {BsFillBarChartFill} from "react-icons/bs";
import "./styles.css";

import * as data from './footer.json'
const linksString = JSON.stringify(data);
const links = JSON.parse(linksString).links;

type Link = {
  label: string;
  href: string;
  target: string;
}

const Links: React.FC<{links: Link[]}> = ({ links }) => {
  return (
      <div className='links-container'>
      
          {links.map((link: Link) => {
              return (

                  <div key={link.href} className='link'>
                      <a href={link.href} target={link.target} style={{textDecoration: "none", color: "midnightblue"}}>
                          {link.label}
                      </a>
                  </div>
              )
          })}
      </div>
  )
}

export function Footer() {
  return (
    <nav className="Footer-container">
      <div className='logo-container'>
              <a href = "/" style={{textDecoration: "none", color: "midnightblue"}}>

          {" "}
              <BsFillBarChartFill />  Mind My Wallet  {""}
              </a>
      </div>
      <Links links={links}/>
    </nav>
  );
}
