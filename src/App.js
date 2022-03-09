import React from 'react';
import Login from "./pages/Login";
import './App.css'; import { BrowserRouter, Route, Routes } from "react-router-dom";
// import Home from "./pages/Home";
import Register from "./pages/Register";
import AddBook from "./pages/AddBook";
import UpdateBook from "./pages/UpdateBook";
import Protected from "./pages/Protected";
import AddAuthor from "./pages/AddAuthor";
import BookList from "./pages/BookList";
import AuthorList from "./pages/AuthorList";
import UpdateAuthor from "./pages/UpdateAuthor";



function App() {
    return (
        <div className="App">
            <BrowserRouter>

                <main className="form-signin">
                    <Route>
                        {/* <Route path="/" element={<Home />}></Route>
                        <Route path="/login" element={<Login />}></Route>
                        <Route path="/register" element={<Register />}></Route>
                        <Route path="/addauthor" element={<AddAuthor />}></Route>
                        <Route path="/authorlist" element={<AuthorList />}></Route>
                        <Route path="/updateauthor" element={<UpdateAuthor />}></Route>
                        <Route path="/addbook" element={<AddBook />}></Route>
                        <Route path="/booklist" element={<BookList />}></Route>
                        <Route path="/updatebook/" element={<UpdateBook />}></Route> */}
                        {/* <Route path="/" render={routeProps => (<Home routeProps={routeProps} animate={true} />)}/> */}
                        <Route path="/login/" render={routeProps => (<Login routeProps={routeProps} animate={true} />)}/>
                        <Route path="/register/" render={routeProps => (<Register routeProps={routeProps} animate={true} />)}/>
                        <Route path="/addauthor/" render={routeProps => (<AddAuthor routeProps={routeProps} animate={true} />)}/>
                        <Route path="/authorlist/" render={routeProps => (<AuthorList routeProps={routeProps} animate={true} />)}/>
                        <Route path="/updateauthor/:id" render={routeProps => (<UpdateAuthor routeProps={routeProps} animate={true} />)}/>
                        <Route path="/addbook/" render={routeProps => (<AddBook routeProps={routeProps} animate={true} />)}/>
                        <Route path="/booklist/" render={routeProps => (<BookList routeProps={routeProps} animate={true} />)}/>
                        <Route path="/updatebook/:id" render={routeProps => (<UpdateBook routeProps={routeProps} animate={true} />)}/>
                        </Route>

                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;


{/* <Route path="/" element={<App />}></Route> */}