import { useState } from 'react';
import { FormatCodeSettings } from 'typescript';

export function Search() {
    const [message, setMessage] = useState('');
    const [updated, setUpdated] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement> ) => {
    if (event.key === 'Enter') {
      handleClick();
    }
};

  const handleClick = () => {
    setMessage("");
    console.log(message.toUpperCase())
  };

    return  (
        <div className="SearchTop">
                <input type="text" placeholder="Stock Ticker" onChange={handleChange} value={message} name="message" id="message" onKeyDown={handleKeyDown}/>
                <button className="submit" type="submit" onClick={handleClick}>Search</button>
        </div>
    );
}