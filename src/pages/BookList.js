import Nav from "../components/Nav";
import React, { useState, useEffect } from "react";
import { Table } from 'react-bootstrap';
import { Link } from "react-router-dom";

function BookList() {

    const [data, setData] = useState([]);
    useEffect(() => {
        getData();
    }, [])

    async function deletebooks(ID) {
        let result = await fetch("http://localhost:8080/books/" + ID, {
            method: "DELETE"
        });
        result = await result.json();
        console.warn("result", data)
        getData()
    }
    async function getData() {
        let result = await fetch("http://localhost:8080/books");
        result = await result.json();
        setData(result)
    }

    return (
        <div>
            <Nav />
            <div className="col-sm-6 offset-sm-3">

                <h1>BooksList</h1>
                <Table>
                    <tbody>
                        <tr>
                            <th>Name</th>
                            <th>Description</th>
                            <th>Delete Books</th>
                            <th>Update Books</th>

                        </tr>
                        {
                            data.map((item) =>
                                <tr>
                                    
                                    <td>{item.name}</td>
                                    <td>{item.description}</td>
                                    <td><span onClick={() => deletebooks(item.ID)} className="delete">Delete</span></td>
                                    <td>
                                        <Link to={"updatebook/" + item.ID}>
                                            <span className="update">Update</span>
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

export default BookList;
