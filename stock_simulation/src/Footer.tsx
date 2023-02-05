import { NavLink } from "react-router-dom";

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
                      <a href={link.href} target={link.target}>
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
            <span>Logo</span>
      </div>
      <Links links={links}/>
    </nav>
  );
}
