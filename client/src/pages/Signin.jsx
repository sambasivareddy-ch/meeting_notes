import React from "react";

import GoogleSigninButton from "../components/GoogleSigninButton";
import styles from '../styles/signin.module.css';

const Signin = (props) => {
    return (
        <div className={styles['signin-page_wrapper']}>
            <GoogleSigninButton/>
        </div>
    )
}

export default Signin;