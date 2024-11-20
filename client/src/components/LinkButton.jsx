import React from "react";

import styles from '../styles/button.module.css';

const LinkButton = (props) => {
    return (
        <button className={styles['link-button']} disabled={props.isDisabled}>
            <a href={props.url}>{props.text}</a>
        </button>
    )
}

export default LinkButton;