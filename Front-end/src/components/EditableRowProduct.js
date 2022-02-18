import React from 'react';

const EditableRowProduct = ({item, editFormData, updateFormForText, updateFormForNum, EditFormSubmit, CancelEditForm}) =>{

  return (
        <tr key={item.Id}>
                <td>
                        <h3> </h3>
                </td>
                <td>
                        <input 
                        type="text"
                        name="productname"
                        required="required"
                        placeholder="Enter Product Name" 
                        value={editFormData.productname}
                        onChange={updateFormForText}></input>
                </td>
                <td>
                        <input 
                        type="number"
                        name="quantity"
                        required="required"
                        placeholder="Enter Quantity"
                        value={editFormData.quantity}
                        onChange={updateFormForNum}></input>
                </td>
                <td>
                        <input 
                        type="number"
                        name="sp"
                        required="required"
                        placeholder="Enter Selling Price"
                        value={editFormData.sp}
                        onChange={updateFormForNum}></input>
                </td>
                <td>
                        <input 
                        type="text"
                        name="supplierid"
                        required="required"
                        placeholder="Enter Supplier ID"
                        value={editFormData.supplierid}
                        onChange={updateFormForText}></input>
                </td> 
                
                <td>
                        <button type='submit' className="btn btn-success" onClick={(event) => EditFormSubmit(event, item)}>Save</button>
                        <button type='button' className="btn btn-secondary" onClick={(event) => CancelEditForm()}>Cancel</button>
                </td>
                
        </tr>
  );



};

export default EditableRowProduct;
