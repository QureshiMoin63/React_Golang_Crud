import Nav from "../components/Nav";
import {withRouter} from 'react-router-dom';
import { useState, useEffect } from 'react';
import { NavItem } from "react-bootstrap";

function UpdateBook(props) {
    const [data, setData] = useState([])
    const[name,setName]=useState("")
    const[description,setDescription]=useState("")
     const formData = new FormData();
     console.warn("props", props.match.params.id)

    useEffect(async () => {
        let result = await fetch("http://localhost:8080/books/" + props.match.params.id);
        result = await result.json();
        console.log(name,description)
        setData(result)
        setName(result.name)
        setDescription(result.description)
    }
    ,[])    
    async function editBook(id){
        formData.append('name',name);
        formData.append('description',description);
        console.log(formData.get('name')) 
        console.log(formData.get('description')) 
        let result = await fetch("http://localhost:8080/books/" + props.match.params.id, {
            method: "PUT",
            body: formData
        })
         result = await result.json();
        
        alert(id)
    }
    return (
        <div>
            <Nav />
            <div className="col-sm-6 offset-sm-3">
                <h1>Update Book</h1>
                <input type="text" onChange={(e) => setName(e.target.value)} className="form-control" defaultValue={data.name} /><br />
                <input type="text" onChange={(e) => setDescription(e.target.value)} className="form-control" defaultValue={data.description} /><br />
                <button onClick={() => editBook(data.id)} className="w-100 btn btn-lg btn-primary">Update Book</button>
            </div>
        </div>
    )
}

export default withRouter(UpdateBook);