import React from "react";

import { getGoogleAuthURL } from "../utils/urls";
import GoogleIcon from "@mui/icons-material/Google";
import styles from "../styles/googlesigninbutton.module.css";

const GoogleSigninButton = (props) => {
    const authUrl = getGoogleAuthURL();

    const handleGoogleSignin = () => {
        window.location.href = authUrl;
    };

    return (
        <button
            type="button"
            className={styles["signin-button"]}
            onClick={handleGoogleSignin}
        >
            <div className={styles["submit-text"]}>
                <GoogleIcon />
                <span>Sign in with Google</span>
            </div>
        </button>
    );
};

export default GoogleSigninButton;
