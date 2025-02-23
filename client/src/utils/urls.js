export const getGoogleAuthURL = () => {
    const requestOptions = {
        redirect_uri: import.meta.env.VITE_APP_REDIRECT_URI,
        client_id: import.meta.env.VITE_APP_CLIENT_ID,
        access_type: "offline",
        response_type: "code",
        prompt: "consent",
        scope: [
            "https://www.googleapis.com/auth/userinfo.profile",
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/calendar.readonly",
        ].join(" "),
    };

    const searchParams = new URLSearchParams(requestOptions);

    return `${import.meta.env.VITE_APP_AUTH_URI}?${searchParams.toString()}`;
};
