import Nav from "../components/Nav";
import React, { useState, useEffect } from "react";
import { Table } from 'react-bootstrap';
import { Link } from "react-router-dom";


function AuthorList(){
    const [data, setData] = useState([]);
    useEffect( () => {
        getData();
    }, [])

async function  deleteOperation(ID){
    let result = await fetch("http://localhost:8080/authors/"+ID,{
        method:'DELETE'
    });
    result = await result.json();
    console.warn("result",data)
    getData()
}   

async function getData(){
    let result = await fetch("http://localhost:8080/authors");
    result = await result.json();
    setData(result)
}   

 return(
    <div>
        <Nav />

        <div className="col-sm-8 offset-sm-2">
            <h1>AuthorsList</h1>
            <Table>
                <tbody>
                <tr>
                    <th>Name</th>
                    <th>Description</th>
                    <th>Delete Entry</th>
                    <th>Update Authors</th>
                </tr>
                {
                    data.map((item) => 
                        <tr>
                            <td>{item.name}</td>
                            <td>{item.description}</td>
                            <td><span onClick={() =>deleteOperation(item.ID) } className="delete">Delete</span></td>
                            <td>
                                <Link to ={"updateauthor/"+item.ID}>
                                <span  className="update">Update</span>
                                </Link>
                                </td>
                        </tr>
                    )
                }
                </tbody>
            </Table>
            </div>
    </div>
)
}

export default AuthorList;





