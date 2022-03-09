import Nav from "../components/Nav";
import { useState, useEffect } from "react";

function Login() {

    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [redirect, setRedirect] = useState("")

    useEffect(() => {
        if (localStorage.getItem('user-info')) {
            return <redirect to="/addbook"/>
        }
    }, [])

   
    async function login() {
        try{
            let item = { email, password  }
            console.warn(item)
          let response = await fetch("http://localhost:3000/login", {
               method: "POST",
                body: JSON.stringify(item),
            })
            response =  response.json()
            localStorage.setItem("pwd",password )
            localStorage.setItem("email",email )
            window.location.href="/addBook"
            setRedirect(item);
            console.log(redirect)
            if (redirect) {
                return <redirect to="/addbook"/>
            }
            if (response.status === 200) {
              console.log("success")
              } else {
               console.log("failure")
              }
        }
        catch(err){
            console.log(err)
        }
        }
    return (
        <>
            <Nav />
            <form>
                <div className="col-sm-6 offset-sm-3">
                    <h1 className="h3 mb-3 fw-normal">Enter Credentials</h1>
                    <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} className="form-control" placeholder="Email address" />
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} className="form-control" placeholder="Password" />
                    <button onClick={login} className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
                </div>
            </form>
        </>
    );
};
export default Login;