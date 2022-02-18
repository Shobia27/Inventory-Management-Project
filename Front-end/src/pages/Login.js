import axios from "axios";
import React, { useState } from "react";
import { Navigate } from 'react-router-dom';

const LoginApi= "http://localhost:3000/api/login"
const Login= ()=>{
    const [email, setEmail]= useState('')
    const [password, setPassword]= useState('')
    const [redirect, setRedirect]= useState(false)

    const LoginData= {
      email: '',
      password: ''
    }

    const login=(e) =>{
        e.preventDefault();
        console.log({
             email,
              password
        });
        LoginData["email"]=email;
        LoginData["password"]=password;
        console.log(LoginData);

        //writing to user table in db
        axios.post(LoginApi, LoginData)
        .then(function (response) {

          console.log(response.data)
          alert("Login Successfull!!")
          setRedirect(true);
          })
        .catch(function (error) {
          console.log(error.response.data);
          alert("OOPS!! Incorrect Credentials! Try again.. :(")
        });

        // alert('You are successfully Registered!!')
        //setRedirect(true);
    }

    if(redirect){
        
        return <Navigate to="/" />;
    }
    return(
        <form onSubmit={login} className="form-signin">
          <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
          <h5>Email ID</h5>
          <div className="form-floating">
            <input type="email" className="form-control"  placeholder="Email Address" 
            onChange={e=> setEmail(e.target.value)}/>
          </div>
          <h5>Password</h5>

          <div className="form-floating">
            <input type="password" className="form-control"  placeholder="Password"
            onChange={e=> setPassword(e.target.value)}/>
          </div>

          <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        </form>
    );
}

export default Login;