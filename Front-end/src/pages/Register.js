import axios from "axios";
import React, { useState } from "react";
import { Navigate } from 'react-router-dom';


// const Register= ()=>{
//     return(
//         <div>
//             Register
//         </div>
//     );
// }

const RegisterApi= "http://localhost:3000/api/registerUser"

function Register(){
    const [Name, setName]= useState('')
    const [Email, setEmail]= useState('')
    const [Password, setPassword]= useState('')
    const [redirect, setRedirect]= useState(false)

    const RegisterData= {
        name: '',
        email: '',
        password: ''
      }
    const submit=(e) =>{
        e.preventDefault();
        console.log({
            Name,
             Email,
              Password
        });


        RegisterData["name"]=Name;
        RegisterData["email"]=Email;
        RegisterData["password"]=Password;
        console.log(RegisterData)

        //writing to user table in db
        axios.post(RegisterApi, RegisterData)
        .then(function (response) {
        console.log(response)
        })
        .catch(function (error) {
        console.log(error);
        });

        alert('You are successfully Registered!!')
        setRedirect(true);
    }

    if(redirect){
        
        return <Navigate to="/login" />;
    }
    

    // renderRedirect = () => {
    //     if (redirect) {
    //       return <Redirect to='/login' />
    //     }
    //   }
    
    return <div>
                <form onSubmit={submit} className="form-signin"> 
                    {/* {this.renderRedirect()} */}

                    <h1 className="h3 mb-3 fw-normal">Please Register</h1>
                    <h5>Name</h5>
                    
                    <input  
                    type="text"
                    name="name"
                    required="required"
                    className="form-control"
                     placeholder="Name" 
                    onChange={e=> setName(e.target.value)}/>
                    
                    <h5>Email ID</h5>
                    
                    <input type="email" className="form-control"  placeholder="Email Address"
                    onChange={e=> setEmail(e.target.value)} />
                    
                    <h5>Password</h5>

                    <input type="password" className="form-control"  placeholder="Password"
                    onChange={e=> setPassword(e.target.value)}/>
                    
                    <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
                </form>
           </div>
}

export default Register;