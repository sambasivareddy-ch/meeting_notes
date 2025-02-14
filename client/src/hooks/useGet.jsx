const useGet = (url, includeCredentials) => {
    const fetchUrl = async () => {
        const response = await fetch(url, {
            method: "GET",
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const data = await response.json();
        return data;
    };

    return fetchUrl;
};

export default useGet;
