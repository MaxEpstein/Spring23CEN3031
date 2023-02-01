import { NavLink } from "react-router-dom";

import "./styles.css";

import * as data from './links.json'
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
            <span>Logo</span>
      </div>
      <Links links={links}/>
    </nav>
  );
}
