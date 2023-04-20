import { NavLink } from "react-router-dom";

import React, {useEffect, useState} from 'react';
import {BsFillBarChartFill} from "react-icons/bs";
import "./styles.css";
import { sendMsg } from './server';
import { useRef } from 'react';


import * as data from './navbar.json'
import { HomePage } from "./pages/HomePage";
const linksString = JSON.stringify(data);
const links = JSON.parse(linksString).links;

let loggedIn = false;


type Link = {
  label: string;
  href: string;
}

const delay = async (ms: number) => new Promise(
    resolve => setTimeout(resolve, ms)
)

const Links: React.FC<{links: Link[]}> = ({ links }) => {
    const dataFetchedRef = useRef(false);
    const[forceRefresh, setForceRefresh] = useState('');

    const logged = async() => {
        await delay(100);
        let incomming = await sendMsg("LOG");
        console.log("Incomming about log " + incomming);

        if (incomming == 1){
            loggedIn = true;
        }
        else if (incomming == 0){
            loggedIn = false;
        }

        console.log("Loggedin: " + loggedIn);
    }

    useEffect(() => {
        if (dataFetchedRef.current) return;

        else
        {
            logged();
            dataFetchedRef.current = true;
            setForceRefresh("10");
        }
    }) ;

    const logOut = async() => {
        console.log("logging out");

        sendMsg("LOGO");
        let incomming = await sendMsg("LOG");

        if (incomming == 1){
            console.log("Still logged in");
            loggedIn = true;
        }
        else if (incomming == 0){
            console.log("Logged out");
            loggedIn = false;
        }
    }
  return (
      <div className='links-container'>
          {links.map((link: Link) => {
              return (
                  <div key={link.href} className='link'>
                          <a href={link.href} style={{textDecoration: "none", color: "midnightblue"}}>
                              {link.label}
                          </a>
                  </div>
              )
          })}

          <div className={'link'}>
              {loggedIn === true &&
                  <a href={"/dashboard"} style={{textDecoration: "none", color: "midnightblue"}}>
                        {"Dashboard"}
                  </a> &&
                  <a href={"/Login"} style={{textDecoration: "none", color: "midnightblue"}} onClick={logOut}>
                      {"Log Out"}
                  </a>
              }
              {loggedIn === false &&
                  <a href={"/Login"} style={{textDecoration: "none", color: "midnightblue"}}>
                      {"Log In"}
                  </a>
              }
          </div>
      </div>
  )
}

export function NavBar() {

  return (

    <nav className="navbar-container">
      <div className='logo-container' key={"/"}>
        <a href = "/" style={{textDecoration: "none", color: "midnightblue"}}>

                 {" "}
                <BsFillBarChartFill />  Mind My Wallet  {""}
        </a>
      </div>
        <Links links={links}/>
    </nav>
  );
}


