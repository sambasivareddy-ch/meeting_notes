import React from 'react';

import { getGoogleAuthURL } from '../utils/urls';
import GoogleIcon from '@mui/icons-material/Google';
import styles from '../styles/googlesigninbutton.module.css';

const GoogleSigninButton = (props) => {
    const authUrl = getGoogleAuthURL()
    return (
        <button type="button" className={styles['signin-button']}>
            <a href={authUrl} className={styles['auth-url']}>
                <GoogleIcon/> Continue with Google
            </a>
        </button>
    )
}

export default GoogleSigninButton;