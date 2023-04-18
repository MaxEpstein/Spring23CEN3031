import React from 'react';

import "./pageStyles.css";




export function AboutPage() {
    return  (
        <div className={"About"}>
            <div className={"AboutTop"}>
                <h1 style={{textAlign: "center", color: "white", alignItems: "center"
                    }}>About Us</h1>
            </div>
            <div className = "rightSide">
                <p style = {{font: "roboto", textAlign: "center"}}>
                    The purpose of this project is to allow users
                    to participate in the stock market risk free while
                    keeping up with the latest news and events.
                </p>
            </div>

            <div className ="splitScreen">
                <h1
                    style={{textAlign: "center", color: "midnightblue",
                        alignItems: "center", paddingBottom: "50px",}}>
                    Meet the Developers!
                </h1>
                <div className = "row">
                    <div className = "columnL">
                        <img src ={require("../images/ForemanAbout.jpg")} height = "300" width = "400"/>
                    </div>
                    <div className="columnR">
                        <img src = {require("../images/IMG_3471.jpg")} height = "300" width = "400" />
                    </div>
                </div>

                <div className = "row">
                    <div className = "columnL" style = {{marginLeft: 230}}>
                        <h2>Jeanette</h2>
                    </div>
                    <div className = "columnR" style = {{marginLeft: 160}}>
                        <h2>Max</h2>
                    </div>
                </div>

                <div className = "row" style = {{marginTop: 10}}>
                    <div className = "columnL">
                        <img src = {require("../images/bkg3.jpg")}  height = "400" width = "400"/>
                    </div>
                    <div className={"columnR"}>
                        <img src = {require("../images/bkg2.jpg")}  height = "400" width = "400"/>
                    </div>
                </div>

                <div className = "row">
                    <div className = "columnL" style = {{marginLeft: 230}} >
                        <h2>Leandro</h2>
                    </div>
                    <div className = "columnR" style = {{marginLeft: 130}}>
                        <h2>Brayden</h2>
                    </div>
                </div>

            </div>
        </div>
    );
}

