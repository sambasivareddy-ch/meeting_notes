import "./App.css";
import { Route, Routes } from "react-router-dom";

import Signin from "./pages/Signin";
import Meetings from "./pages/Meetings";
import TextEditor from "./pages/TextEditor";
import NotFound from "./pages/NotFound";

function App() {
    return (
        <div className="App">
            <Routes>
                <Route
                    path="/my-meetings/add-notes/:id"
                    element={<TextEditor />}
                    exact
                />
                <Route path="/my-meetings" element={<Meetings />} exact />
                <Route path="/" element={<Signin />} exact />
                <Route path="*" element={<NotFound />} />
            </Routes>
        </div>
    );
}

export default App;
