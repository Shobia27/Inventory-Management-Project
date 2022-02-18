import React from "react";
import { Link } from "react-router-dom";

// const Nav= ()=>{
//     return(
//         <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
//         <div className="container-fluid">
//           <Link to="/" className="navbar-brand" >Home</Link>
          
//           {/* <div class="collapse navbar-collapse" id="navbarCollapse">
//             <ul class="navbar-nav me-auto mb-2 mb-md-0">
//               <li class="nav-item">
//                 <a class="nav-link active" aria-current="page" href="#">Inventory products</a>
//               </li>
//             </ul>
//           </div> */}

//           <div >
//             <ul className="navbar-nav me-auto mb-2 mb-md-0">
//               <li className="nav-item">
//                 <Link to="/login" className="nav-link active" aria-current="page" >Login</Link>
//               </li>
//               <li className="nav-item">
//                 <Link to="/register" className="nav-link active" aria-current="page" >Sign Up</Link>
//               </li>
//             </ul>
//           </div>
//         </div>
//       </nav>
//     );
// }


function Nav(){
  return <div>
              <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
              <Link to="/" className="navbar-brand" >Home</Link>
              
              <div className="collapse navbar-collapse" id="navbarCollapse">
                <ul className="navbar-nav me-auto mb-2 mb-md-0">
                  <li className="nav-item">
                    <Link to="/inventory" className="nav-link active" aria-current="page" >Inventory products</Link>
                  </li>
                  <li className="nav-item">
                    <Link to="/supplier" className="nav-link active" aria-current="page" >Supplier Details</Link>
                  </li>
                </ul>
              </div>

              <div >
                <ul className="navbar-nav me-auto mb-2 mb-md-0">
                  <li className="nav-item">
                    <Link to="/login" className="nav-link active" aria-current="page" >Login</Link>
                  </li>
                  <li className="nav-item">
                    <Link to="/register" className="nav-link active" aria-current="page" >Sign Up</Link>
                  </li>
                </ul>
              </div>
            </div>
          </nav>
         </div>
}

export default Nav