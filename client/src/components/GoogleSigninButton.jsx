import React, { useState } from 'react';

import { getGoogleAuthURL } from '../utils/urls';
import GoogleIcon from '@mui/icons-material/Google';
import styles from '../styles/googlesigninbutton.module.css';

const GoogleSigninButton = (props) => {
    const [authUrl, SetAuthUrl] = useState(getGoogleAuthURL());

    const formSubmitHandler = async (e) => {
        e.preventDefault();

        try {
            const response = fetch(authUrl, {
                method: "GET"
            })

            if (!response.ok) {
                console.log("Authentication Failed")
            }

            const data = (await response).json()
            console.log(data)
            if (data.redirect_url) {
                window.location.href = data.redirect_url;
            }
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <button type="submit" className={styles['signin-button']}>
            <a href={authUrl}>
                <GoogleIcon/> Continue with Google
            </a>
        </button>
    )
}

export default GoogleSigninButton;