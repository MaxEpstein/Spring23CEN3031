import React from 'react';
import "./pageStyles.css";




export function AboutPage() {
    return  (
        <div className={"About"}>
            <div className={"AboutTop"}>
                <h1 style={{textAlign: "center", color: "white", font : "roboto" , alignItems: "center"
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
                <h1 style={{textAlign: "center", color: "midnightblue", font : "roboto" , alignItems: "center",
                }}> Meet the Makers!</h1>

                <div className = "jeanetteInfo">


                    <img src ={require("../images/plsWork.jpeg")} height = "200" width = "200"/>
                    <p style={{textAlign: "center", font: "roboto"}}>
                        Jeanette
                    </p>


                </div>

                <div className = 'maxInfo'>
                    <img src = {require("../images/plsWork.jpeg")} height = "200" width = "200"/>
                    <p style={{textAlign: "center", font: "roboto"}}>
                        Max
                    </p>

                </div>
                <div className = 'leandroInfo'>
                    <img src = {require("../images/plsWork.jpeg")}  height = "200" width = "200"/>
                    <p style={{textAlign: "center", font: "roboto"}}>
                        Leandro
                    </p>

                </div>
                <div className = 'braydenInfo'>
                    <img src = {require("../images/plsWork.jpeg")}  height = "200" width = "200"/>
                    <p style={{textAlign: "center", font: "roboto"}}>
                        Brayden
                    </p>

                </div>

            </div>
        </div>
    );
}

