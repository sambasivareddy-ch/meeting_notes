import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";


import GoogleSigninButton from "../components/GoogleSigninButton";
import styles from "../styles/signin.module.css";
import { login } from "../store/sessionSlice";

const Signin = (props) => {
    const dispatch = useDispatch();
    const isLoggedIn = useSelector((state) => state.sessionStillExists);

    useEffect(() => {
        const params = new URLSearchParams(window.location.search);

        const token = params.get("token");

        if (token) {
            fetch(`${import.meta.env.VITE_APP_SERVER_URI}/set-cookie`, {
                method: "POST",
                credentials: "include",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ token })
            }).then(() => {
                dispatch(login());
                window.location.href = "/my-meetings"
            })
        }
    }, [])

    return (
        <div className={styles["signin-page_wrapper"]}>
            <GoogleSigninButton />
        </div>
    );
};

export default Signin;
