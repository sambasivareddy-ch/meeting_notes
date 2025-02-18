import React from "react";
import { Link } from "react-router-dom";

import "../App.css";

const NotFound = (props) => {
    return (
        <div className="not-found_wrapper">
            <h1 className="not-found_warning">Ooops! 404</h1>
            <p className="not-found_message">The page you are looking for does not exist or <i>login to see.</i></p>
            <Link to="/" className="home-link">
                Click Here for Home
            </Link>
        </div>
    );
};

export default NotFound;
