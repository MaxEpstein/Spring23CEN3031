import * as React from "react";
import "./pageStyles.css";
import { login } from "./login";
import { Dash } from "./Dash";
import { Redirect, Route, useHistory } from "react-router-dom";



interface LoginState {
  password: string;
  username: string;
  isLoading: boolean;
  error: string;
  isLoggedIn: boolean;
}

type LoginAction =
  | { type: "login" | "success" | "error" | "logout" }
  | { type: "field"; fieldName: string; payload: string };

const loginReducer = (state: LoginState, action: LoginAction): LoginState => {
  switch (action.type) {
    case "field": {
      return {
        ...state,
        [action.fieldName]: action.payload
      };
    }
    case "login": {
      return {
        ...state,
        error: "",
        isLoading: true
      };
    }
    case "success": {
      return { ...state, error: "", isLoading: false, isLoggedIn: true };
    }
    case "error": {
      return {
        ...state,
        isLoading: false,
        isLoggedIn: false,
        username: "",
        password: "",
        error: "Incorrect username or password!"
      };
    }
    case "logout": {
      return {
        ...state,
        isLoggedIn: false
      };
    }
    default:
      return state;
  }
};

const initialState: LoginState = {
  password: "",
  username: "",
  isLoading: false,
  error: "",
  isLoggedIn: false
};

const refreshPage = ()=>{
    window.location.reload();
 }

  

export function Login() {
  const [state, dispatch] = React.useReducer(loginReducer, initialState);
  const { username, password, isLoading, error, isLoggedIn } = state;

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    
    e.preventDefault();
    dispatch({ type: "login" });

    try {
      await login({ username, password });
      dispatch({ type: "success" });

    } catch (error) {
      dispatch({ type: "error" });
    }
  };

 


  return (
    
    <div className="App">
      <div className="login-container">
        {isLoggedIn ? (
           <Redirect to={"/dashboard"} ></Redirect>
) : (
          <form className="form" onSubmit={onSubmit}>
            {error && <p className="error">{error}</p>}
            <h1> Please Login!</h1>
            <input
              type="text"
              placeholder="username"
              value={username}
              onChange={(e) =>
                dispatch({
                  type: "field",
                  fieldName: "username",
                  payload: e.currentTarget.value
                })
              }
            />
            <input
              type="password"
              placeholder="password"
              autoComplete="new-password"
              value={password}
              onChange={(e) =>
                dispatch({
                  type: "field",
                  fieldName: "password",
                  payload: e.currentTarget.value
                })
              }
            />
            <button type="submit" className="submit" id="action" value="login" disabled={isLoading}>
              {isLoading ? "Logging in....." : "Log in / Sign up"}
            </button>
           
          </form>
        )}
      </div>
    </div>
  );
}
