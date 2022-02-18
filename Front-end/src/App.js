import React, { Component } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import Nav from './components/Nav';
import Home from './pages/Home';
import Inventory from './pages/Inventory';
import Login from './pages/Login'
import Register from './pages/Register';
import Supplier from './pages/Supplier';


class App extends Component {
  
  render(){
 
  return (
    
    <BrowserRouter>
    <div className="App">
      
      <Nav />

      {/* <main> className="form-signin"> */}
        <main>
          <Routes>
            <Route path="/" exact element={<Home/>} />
            <Route path="/login" exact element= {<Login/>}/>
            <Route path="/register" exact element= {<Register/>}/>
            <Route path="/inventory" exact element= {<Inventory/>}/>
            <Route path="/supplier" exact element= {<Supplier/>}/>

          </Routes>
        

      </main>
    </div>
    </BrowserRouter>

  );
}
}

export default App;
