import React from "react";
import { Link } from "react-router-dom";

import '../App.css';

const NotFound = (props) => {
    return (
        <div className="not-found_wrapper">
            <h1 className="not-found_warning">Page Not Found</h1>
            <Link to="/" className="home-link">Click Here for Home</Link>
        </div>
    )
}

export default NotFound;