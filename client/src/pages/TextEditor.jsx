import React, { useCallback } from "react";

import Quill from "quill";
import Button from "../components/Button";

import "quill/dist/quill.snow.css";
import "../styles/texteditor.css";

const TextEditor = (props) => {
    const saveBtnClickHandler = () => {
        console.log(document.querySelector(".ql-editor"));
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
        wrapper.append(editor);
        new Quill(editor, { theme: "snow" });
    }, []);

    return (
        <div className="text-editor_wrapper">
            <div className="text-editor" ref={wrapperRef}></div>
            <div className="editor-btn_wrapper">
                <Button
                    text="Save Notes"
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
