import React, { useEffect } from "react";
import { Route, Routes, useLocation } from "react-router-dom";
import { useSelector, useDispatch } from "react-redux";
import { Analytics } from "@vercel/analytics/react"

import "./App.css";

import Signin from "./pages/Signin";
import Meetings from "./pages/Meetings";
import TextEditor from "./pages/TextEditor";
import NotFound from "./pages/NotFound";
import PrivacyPolicy from "./pages/PrivacyPolicy";
import { login } from "./store/sessionSlice";

function App() {
    const dispatch = useDispatch();
    const location = useLocation();
    const isLoggedIn = useSelector((state) => state.sessionStillExists);

    useEffect(() => {
        const getAndSetLoginStatus = async () => {
            try {
                const response = await fetch(`https://meeting-notes-7izd.onrender.com/loginstatus`, {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    credentials: "include"
                })
                if (!response.ok) {
                    console.log("An error occcurred")
                }

                const data = await response.json();

                if (data.isLoggedIn) {
                    dispatch(login())
                    window.location.href = '/my-meetings'
                }
                
            } catch (e) {
                console.log(e)
            }
        }

        if (location.pathname === '/')
            getAndSetLoginStatus()
    }, [])

    return (
        <div className="App">
            <Routes>
                {isLoggedIn &&
                    <Route
                        path="/my-meetings/add-notes/:id"
                        element={<TextEditor isEditMode={false} />}
                        exact
                    />
                }
                {isLoggedIn &&
                    <Route
                        path="/my-meetings/edit-notes/:id"
                        element={<TextEditor isEditMode={true} />}
                        exact
                    />
                }
                {isLoggedIn && <Route path="/my-meetings" element={<Meetings />} exact />}
                {<Route path="/privacy-policy" element={<PrivacyPolicy/>} exact />}
                {!isLoggedIn && <Route path="/" element={<Signin />} exact />}
                <Route path="*" element={<NotFound />} />
            </Routes>
            <Analytics/>
        </div>
    );
}

export default App;
