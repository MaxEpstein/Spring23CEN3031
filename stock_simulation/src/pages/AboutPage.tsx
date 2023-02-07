

export function AboutPage() {
    return  (
        <div className={"About"}>
            <div className={"AboutTop"}>
                <h1 style={{textAlign: "center", color: "white", font : "roboto"}}>About Us</h1>
            </div>
            <div className = "AboutBottom">
                <p style = {{font: "roboto"}}>
                    The purpose of this project is to allow users
                    to participate in the stock market risk free while
                    keeping up with the latest news and events.
                </p>
            </div>
        </div>
    );
  }