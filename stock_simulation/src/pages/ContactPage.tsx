

export function ContactPage() {
    return (
        <div className = "contact">
            <h1 style={{textAlign: "center"}}>Contact Us</h1>

            <form id= "contact-form" method = "Post">
                <label htmlFor={"name"}>Full Name </label>
                <br/>
                <input
                    name = "name"
                    placeholder = "Enter full name..."
                    type = "text"
                />
                <br/> <br/>
                <label htmlFor={"email"}>Email </label>
                <br/>
                <input
                    name = "email"
                    placeholder = "Enter email..."
                    type = "email"
                />
                <br/>  <br/>
                <label htmlFor= "message" > Message</label>
                <br/>
                <textarea
                    rows={6}
                    placeholder = "Enter message"
                > </textarea>
                <br/>
                <button type = "submit"> Send Message</button>
            </form>
        </div>
        )
  }
  