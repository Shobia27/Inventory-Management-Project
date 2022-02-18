import axios from "axios";
import React, { Fragment, useEffect, useState } from "react";
import EditableRowProduct from "../components/EditableRowProduct";
import ReadOnlyRowProduct from "../components/ReadOnlyRowProduct";

// const InventoryProductUrl = "http://localhost:3000/api/getAllProducts";
const InventoryProductUrl = "http://localhost:3000/api/redis/getAllProductCache";

const AddProductApi="http://localhost:3000/api/addProduct";

function Inventory(){
  
  //-----------------------------------------------------------------
  // to recieve all product data fro api
  //-----------------------------------------------------------------
  const [productData, setProduct]= useState([])
  const getProductData = async() =>{
    try{
      const data= await axios.get(InventoryProductUrl)
      //console.log(data.data)
      setProduct(data.data)
    }
    catch(e){
      console.log(e)
    }
  }

  useEffect(()=>{
    getProductData()
  },[])

  //----------------------------------------------------------------
  //to add record to product database
  //----------------------------------------------------------------
  const [addFormData, setAddFormData]= useState({
    productname: '',
    quantity: 0,
    sp: 0,
    supplierid: ''
  })

  //to take input as string
  const onFormChanged=(event) =>{
    event.preventDefault();
    const fieldname= event.target.getAttribute('name');
    const fieldvalue= event.target.value;
   
    const newFormData= {...addFormData};
    newFormData[fieldname]= fieldvalue;
    setAddFormData(newFormData);
  }

  //to take input as number
  const onFormChanged1=(event) =>{
    event.preventDefault();
    const fieldname= event.target.getAttribute('name');
    const fieldvalue= event.target.valueAsNumber;

    const newFormData= {...addFormData};
    newFormData[fieldname]= fieldvalue;
    setAddFormData(newFormData);
  }

  //adding the record to db
  const addProduct=(event) =>{
    event.preventDefault();
    console.log(addFormData);

    //adding record in postgresql product db
    axios.post(AddProductApi, addFormData)
    .then(function (response) {
      console.log(response)
    })
    .catch(function (error) {
      console.log(error);
    });
    
    alert(`A new Record ${addFormData.productname} is Inserted!!`)
    window.location.reload(false);

  }

  //----------------------------------------------------------------
  //to delete record from product database
  //----------------------------------------------------------------
  function deleteProduct(id){

    //deleting record in postgresql product db
    axios.delete(`http://localhost:3000/api/deleteProduct/${id}`)
    .then(function (response) {
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });

    //deleting record in redis db
    axios.delete(`http://localhost:3000/api/redis/deleteProductCache/${id}`)
    .then(function (response) {
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });


    alert(`Record with Id ${id} is Deleted!!`)
    window.location.reload(false);
  }

  //----------------------------------------------------------------
  //to update record to product database
  //----------------------------------------------------------------
  const [editFormData, setEditFormData]= useState({
    productname: '',
    quantity: 0,
    sp: 0,
    supplierid: ''
  })
  const [editProducId, setEditProductId]= useState(null);

  const openEditForm= (event, item)=> {
    event.preventDefault();
    setEditProductId(item.Id);
    console.log(item.Id);

    const formvalues={
      productname: item.productname,
      quantity: item.quantity,
      sp: item.sp,
      supplierid: item.supplierid,
    }

    setEditFormData(formvalues);
  }

  //for text
  const updateFormForText =(event) => {
    event.preventDefault();

    const fieldname= event.target.getAttribute("name");
    const fieldvalue= event.target.value;

    const newFormData= {...editFormData};
    newFormData[fieldname]= fieldvalue;

    setEditFormData(newFormData);
    console.log(newFormData);
  }

  //for number
  const updateFormForNum =(event) => {
    event.preventDefault();

    const fieldname= event.target.getAttribute("name");
    const fieldvalue= event.target.valueAsNumber;

    const newFormData= {...editFormData};
    newFormData[fieldname]= fieldvalue;

    setEditFormData(newFormData);
    console.log(newFormData);
  }

  const EditFormSubmit =(event, item) =>{
    event.preventDefault();

    const editedProduct= {
      productname: editFormData.productname,
      quantity: editFormData.quantity,
      sp: editFormData.sp,
      supplierid: editFormData.supplierid
    }

    // console.log(editedProduct);
    // console.log(item.Id);

    //updating record in postgresql product db
    axios.put(`http://localhost:3000/api/updateProduct/${item.Id}`, editedProduct)
    .then(function (response) {
      console.log(response);
    })
    .catch(function (error) {
      console.log(error);
    });

    alert(`Record with Id ${item.Id} is Updated!!`)
    window.location.reload(false);
  }

  const CancelEditForm =() =>{
    setEditProductId(null);
    window.location.reload(false);
  }

  //----------------------------------------------------------------
  //read the records from product database 
  //----------------------------------------------------------------
  let product_tb_data= productData.map((item, i)=> {
    return(
      // <tr key={item.Id}>
      //   <td>{item.Id}</td>
      //   <td>{item.productname}</td>
      //   <td>{item.quantity}</td>
      //   <td>{item.sp}</td>
      //   <td>{item.supplierid}</td>
      //   <td>
      //     <button className="btn btn-success" >Edit</button>{'  '}
      //     <button className="btn btn-danger" onClick={()=> deleteProduct(item.Id)}>Delete</button>

      //   </td>
      // </tr>

      
      <Fragment key={i}>
         {editProducId === item.Id?(
           <EditableRowProduct 
           editFormData= {editFormData}
           updateFormForText= {updateFormForText}
           updateFormForNum= {updateFormForNum}
          EditFormSubmit= {EditFormSubmit}
          CancelEditForm= {CancelEditForm}
           item={item}
            />
         ):(
           <ReadOnlyRowProduct 
           myitem={item}
           deleteProduct= {deleteProduct}
           openEditForm={openEditForm}/>
         )
         }
        
      </Fragment>
      
    )
  })

  return <div className="container">
              <h1 style={{display: 'flex',  justifyContent:'center'}}>Inventory Product Details</h1>              
              <br></br>

              {/*------------------------------------------------------- 
                  displaying the table 
              ----------------------------------------------------------*/}
              <form >
              <table className="table table-hover">
                <thead>
                  <tr>
                    <th>Product ID</th>
                    <th>Product Name</th>
                    <th>Quantity</th>
                    <th>Selling Price</th>
                    <th>Supplier ID</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {product_tb_data}
                </tbody>
              </table>
              </form>
              
              {/*------------------------------------------------------- 
                  form to get input for adding the record 
              ----------------------------------------------------------*/}     
              <br></br>
              <form onSubmit={addProduct} >
                <input 
                  type="text"
                  name="productname"
                  required="required"
                  placeholder="Enter Product Name"
                  onChange={onFormChanged}/>{'   '}
                <input 
                  type="number"
                  name="quantity"
                  required="required"
                  placeholder="Enter Quantity"
                  onChange={onFormChanged1}/>{'   '}
                <input 
                  type="number"
                  name="sp"
                  required="required"
                  placeholder="Enter Selling Price"
                  onChange={onFormChanged1}/>{'   '}
                <input 
                  type="text"
                  name="supplierid"
                  required="required"
                  placeholder="Enter Supplier ID"
                  onChange={onFormChanged}/> {'   '}
                <button type="submit" className="btn btn-primary">Add Product</button>{'     '}
                {/* <button type="button" className="btn btn-warning" onClick={()=> setInputFieldValues()}>Update</button> */}

              </form>

          </div>

    
}
export default Inventory;