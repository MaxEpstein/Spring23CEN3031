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
  button: number;
  isLoadingS: boolean;
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
        error: "Incorrect username or password!",
        button: 1,
        isLoadingS: false
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
  isLoggedIn: false,
  button: 1,
  isLoadingS: false
};

const refreshPage = ()=>{
    window.location.reload();
 }

  

export function Login() {
  const [state, dispatch] = React.useReducer(loginReducer, initialState);
  const { username, password, isLoading, error, isLoggedIn, button, isLoadingS } = state;

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (state.button == 1){
      console.log("loggin in");
    
      dispatch({ type: "login" });

      try {
        await login({ username, password });
        dispatch({ type: "success" });

      } catch (error) {
        dispatch({ type: "error" });
      }
    }
    else if (state.button == 2){
      console.log("signing up");

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
            <button onClick={() => (state.button = 1)} type="submit" className="submit" id="action" value="login" disabled={isLoading}>
              {isLoading ? "Logging in....." : "Log in"}
            </button>
            <button className="submit" onClick={() => (state.button = 2)} disabled={isLoadingS}>
              {isLoadingS ? "Signing up....." : "Sign up"}
            </button>
           
          </form>
        )}
      </div>
    </div>
  );
}
