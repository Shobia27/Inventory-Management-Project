import React from "react";
import homeImg from "./images/InventoryManagement.png"
// const Home= ()=>{
//     return(
//         <div>
//             Home
//         </div>
//     );
// }

function Home(){
    return <div className="container">
                <h1 style={{display: 'flex',  justifyContent:'center'}}>Inventory Management Project</h1>
                <img src={homeImg} width="85%"  alt="Home_Page_Image"/>
           </div>
}
export default Home;