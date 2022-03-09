import { useEffect, useState } from "react";
import { withRouter } from 'react-router-dom';
import Nav from "../components/Nav";

function UpdateAuthor(props) {
    const [data, setData] = useState([])
    const [name, setName] = useState("")
    const [description, setDescription] = useState("")
    const formData = new FormData();
    console.warn("props", props.match.params.id)

    useEffect(async () => {
        let result = await fetch("http://localhost:8080/authors/" + props.match.params.id);

        result = await result.json();
        console.log(name, description)
        setData(result)
        setName(result.name)
        setDescription(result.description)
    }, [])
    

    async function editAuthor(ID){
        try{
            formData.append('name',name);
            formData.append('description',description);
            console.log(formData.get('name')) 
            console.log(formData.get('description')) 
            let result = await fetch ("http://localhost:8080/authors/"+ props.match.params.id,{
                method: "PATCH",
                body:formData
            })
            result =await result.json();
            alert(ID)
        }
        catch(err){
            console.log(err)
        }
    }
  
    return (
        <div>
            <Nav />
            <div className="col-sm-6 offset-sm-3">
                <h1>Update Author</h1>
                <input type="text" onChange={(e) => setName(e.target.value)} className="form-control" defaultValue={data.name} /> <br /> <br />
                <input type="text" onChange={(e) => setDescription(e.target.value)} className="form-control" defaultValue={data.description} /> <br /> <br />

                <button onClick={(e) => editAuthor(data.ID)} className="w-100 btn btn-lg btn-primary">Update Author</button>
            </div>
        </div>
    )
}

export default withRouter(UpdateAuthor);


