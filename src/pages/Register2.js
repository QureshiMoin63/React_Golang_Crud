// import Nav from "../components/Nav"
// import { useState, useEffect } from 'react';
// function Register() {
//     useEffect(() => {
//         if (localStorage.getItem('user-info')) {
//             return <redirect to="/addBook" />
//         }
//     }, [])
//     const [email, setEmail] = useState("")
//     const [username, setName] = useState("")
//     const [password, setPassword] = useState("")
//     //    const[redirect, setRedirect]=useState("")

//     async function signUp() {
//         try {
//             let item = { email, username, password }
//             console.warn(item)

//             let result = await fetch("http://localhost:8080/register", {
//                 method: 'POST',
//                 body: JSON.stringify(item),
//             })
//             result = result.json()
//             window.location.href = "/addBook"

//             //      console.warn("result",result)
//             localStorage.setItem("name", username)
//             localStorage.setItem("password", password)
//             localStorage.setItem("email", email)
//             //  setRedirect(item);
//         }
//         catch (err) {
//             console.log(err)
//         }
//     }
//     return (
//         <>
//             <Nav />
//             <form>
//                 <div className="col-sm-6 offset-sm-3">
//                     <h1 className="h3 mb-3 fw-normal">Register</h1>
//                     <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} className="form-control" placeholder="Email" />
//                     <input type="text" value={username} onChange={(e) => setName(e.target.value)} className="form-control" placeholder="Name" />
//                     <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} className="form-control" placeholder="Password" />
//                     <button onClick={signUp} className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
//                 </div>
//             </form>
//         </>
//     )


// }

// export default Register;