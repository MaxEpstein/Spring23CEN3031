import React from "react";
import {useRef} from "react";
import "./pageStyles.css";
import emailjs from '@emailjs/browser';


export function ContactPage() {
    const form = useRef<HTMLFormElement>(null) ;

    const sendEmail = (e: React.FormEvent) => {
        e.preventDefault();

        emailjs.sendForm(
            'service_jv4htf3',
            'template_waiq5pl',
            form.current!,
            'Vztn8121DqbQGav4q')
            .then((result) => {
                console.log(result.text);
                console.log("Message sent!")
            }, (error) => {
                console.log(error.text);
            });
    };
    return (
        <div className = "contact">
            <div className = "center">

                <h1 style={{textAlign: "center"}}>Contact Us</h1>

                <form ref= {form} onSubmit={sendEmail}>
                    <label>Name</label>
                    <input type="text" name="name" placeholder = "Enter name..."/>
                    <label>Email</label>
                    <input type="email" name="email" placeholder = "Enter email..." />
                    <label>Message</label>
                    <textarea name="message" />
                    <input type="submit" value="Send" />
                </form>

            </div>
        </div>
    )
}