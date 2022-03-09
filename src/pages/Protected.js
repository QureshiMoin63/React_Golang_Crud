import React, { useState, useEffect } from "react";
import Nav from "../components/Nav";
function Protected() {
    useEffect(() => {
        if (!localStorage.getItem('user-info')) {
            return <redirect to="/register" />

        }
    })
    return (
        <div>
            <Nav />

        </div>

    );
};
export default Protected;
