import React  from 'react';

const ReadOnlyRowProduct= ({myitem, openEditForm, deleteProduct}) =>{
  
 
  return (
    
    
    // <tr key={item.Id}>
    //     <td>{item.Id}</td>
    //     <td>{item.productname}</td>
    //     <td>{item.quantity}</td>
    //     <td>{item.sp}</td>
    //     <td>{item.supplierid}</td>
    //     <td>
    //       <button className="btn btn-success" type="button" onClick={(event) => updateProduct(event, item)}>Edit</button>{'  '}
    //       <button className="btn btn-danger" >Delete</button>

    //     </td>
    // </tr>

    
    <tr key={myitem.Id}>
        <td>{myitem.Id}</td>
        <td>{myitem.productname}</td>
        <td>{myitem.quantity}</td>
        <td>{myitem.sp}</td>
        <td>{myitem.supplierid}</td>
        <td>
        <button className="btn btn-success" type="button" onClick={(event) => openEditForm(event, myitem)}>Edit</button>{'  '}
          <button className="btn btn-danger" type= "button" onClick={(event) => deleteProduct(myitem.Id)}>Delete</button>

        </td>
      </tr>
    
  );
};

export default ReadOnlyRowProduct;
