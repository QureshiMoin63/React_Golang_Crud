// import React, { useState, useEffect } from "react";
// import Nav from "../components/Nav";

// function Login() {
//     useEffect(() => {
//         if (localStorage.getItem('email','pwd')) {
//         window.location.href="/addBook"
//         }
//     }, [])

//     const [email, setEmail] = useState("")
//     const [password, setPassword] = useState("")
//     const[redirect, setRedirect] = useState("")

//     async function login() {
//         try {
//             let item = { email, password }
//             console.warn(item)
           
//             let response = await fetch("http://localhost:8080/login", {
//                 method: "POST",
//                 body: JSON.stringify(item),

//             })
//             response = response.json()
//             localStorage.setItem("email",email)
//             localStorage.setItem("pwd",password)
//             window.location.href="/addbook"
           
//         }
//         catch (err) {
//             console.log(err)
//         }
//     }


//     return (
//         <div>
//             <Nav />
//             <form>
//                 <div className="col-sm-6 offset-sm-3">

//                     <h1 className="h3 mb-3 fw-normal">Please sign in</h1>

//                     <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} className="form-control" placeholder="Email address" required />

//                     <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} className="form-control" placeholder="Password" required />

//                     <button onClick={login} className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
//                 </div>
//             </form>
//         </div>

//     );
// }
// export default Login;
