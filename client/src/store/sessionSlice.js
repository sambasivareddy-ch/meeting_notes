import { createSlice } from "@reduxjs/toolkit";

const sessionSlice = createSlice({
    name: "session",
    initialState: {
        sessionStillExists: false,
    },
    reducers: {
        login: (state) => {
            state.sessionStillExists = true;
        },
        logout: (state) => {
            state.sessionStillExists = false;
        }
    }
})

export const { login, logout } = sessionSlice.actions;
export default sessionSlice;