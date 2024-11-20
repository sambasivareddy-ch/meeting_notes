import React from "react";
import { Link } from "react-router-dom";

import styles from '../styles/button.module.css';

const NavLinkButton = (props) => {
    const classes = `${props.className} ${styles['nav-link_btn']}`;

    // console.log(routeId)

    return <button className={classes}>
        <Link to={`/my-meetings/add-notes/${props.meeting_id}`}>{props.text}</Link>
    </button>
}

export default NavLinkButton;