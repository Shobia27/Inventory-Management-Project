import axios from "axios";
import React, { useEffect, useState } from "react";

const SupplierDetailUrl = "http://localhost:3000/api/getAllSuppliers";



function Supplier(){
    const [supplierData, setSupplierData]= useState([])
    const getProductData = async() =>{
        try{
        const data= await axios.get(SupplierDetailUrl)
        //console.log(data.data)
        setSupplierData(data.data)
        }
        catch(e){
        console.log(e)
        }
    }

    useEffect(()=>{
        getProductData()
    },[])


    let supplier_tb_data= supplierData.map((item, i)=> {
        return(
        <tr key={item.Id}>
            <td>{item.supplierid}</td>
            <td>{item.suppliername}</td>
            <td>{item.supplieraddress}</td>
            <td>{item.contactnum}</td>
            <td>{item.numproduct}</td>
        </tr>)
        })
    return <div className="container">
                <h1 style={{display: 'flex',  justifyContent:'center'}}>Supplier Details</h1>
                <br></br>

                {/*------------------------------------------------------- 
                    displaying the supplier details table 
                ----------------------------------------------------------*/}
                <form >
                <table className="table table-hover">
                    <thead>
                    <tr>
                        <th>Supplier ID</th>
                        <th>Supplier Name</th>
                        <th>Supplier Address</th>
                        <th>Contact No.</th>
                        <th>No. of Inventory Products</th>
                    </tr>
                    </thead>
                    <tbody>
                    {supplier_tb_data}
                    </tbody>
                </table>
                </form>
           </div>
}
export default Supplier;