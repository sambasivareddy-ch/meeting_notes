import React, { useCallback, useState, useEffect } from "react";
import { useParams } from "react-router-dom";

import Quill from "quill";
import Button from "../components/Button";

import "quill/dist/quill.snow.css";
import "../styles/texteditor.css";

const TextEditor = (props) => {
    const { id } = useParams();
    const [notes, setNotes] = useState("");

    useEffect(() => {
        const getNotes = async () => {
            try {
                const response = await fetch(
                    `http://localhost:8080/meetings/${id}/notes`,
                    {
                        method: "GET",
                        headers: {
                            "Content-Type": "application/json",
                        },
                        credentials: "include",
                    }
                );
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                setNotes(data.notes);
            } catch (error) {
                console.error(error);
            }
        };

        getNotes();
    }, []);

    const saveBtnClickHandler = async () => {
        try {
            const response = await fetch(
                `http://localhost:8080/meetings/${id}/notes`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        notes: document.querySelector(".ql-editor").innerHTML,
                    }),
                    credentials: "include",
                }
            );
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.error(error);
        }
    };

    const printNotesHandler = () => {
        window.print();
    };

    const goBackHandler = () => {
        window.history.back();
    };

    const wrapperRef = useCallback((wrapper) => {
        if (wrapper == null) return;

        wrapper.innerHTML = "";
        const editor = document.createElement("div");
        editor.ariaPlaceholder = "Write your notes here...";
        if (props.isEditMode) editor.innerHTML = notes;
        wrapper.append(editor);
        new Quill(editor, { theme: "snow" });
    }, []);

    return (
        <div className="text-editor_wrapper">
            <div className="text-editor" ref={wrapperRef}></div>
            <div className="editor-btn_wrapper">
                <Button
                    text={props.isEditMode ? "Save Changes" : "Save Notes"}
                    className="save-btn"
                    onClickHandler={saveBtnClickHandler}
                />
                <Button
                    text="Print Notes"
                    className="print-btn"
                    onClickHandler={printNotesHandler}
                />
                <Button
                    text="Go Back"
                    className="go-back_btn"
                    onClickHandler={goBackHandler}
                />
            </div>
        </div>
    );
};

export default TextEditor;
