import React, { useState, useEffect } from "react";
import { createPortal } from "react-dom";

import LinkButton from "./LinkButton";
import NavLinkButton from "./NavLinkButton";
import Button from "./Button";
import NotesModel from "./NotesModel";
import styles from "../styles/meetingcard.module.css";

const MeetingCard = ({
    meeting_id,
    title,
    start_time,
    url,
    description,
    isNotesTaken,
    isDisabled,
    is_today,
    organizer,
    modelShowHandler
}) => {
    const [date, timeFrame] = start_time.split("T");
    const time = timeFrame.split("+")[0];
    const [hour, minute, second] = time.split(":");

    const [notes, setNotes] = useState("");
    const [showNotes, setShowNotes] = useState(false);

    // useEffect(() => {
    //     modelShowHandler();
    // }, [showNotes, modelShowHandler])

    const showNotesButtonHandler = (e) => {
        setShowNotes(true);
    };

    const closesNotesButtonHandler = () => {
        setShowNotes(false);
    };

    useEffect(() => {
        const getNotes = async () => {
            try {
                const response = await fetch(
                    `${import.meta.env.VITE_APP_SERVER_URI}/meetings/${meeting_id}/notes`,
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

    return (
        <div className={styles["meeting-card_wrapper"]}>
            <h2 className={styles["meeting-title"]}>
                {title} {is_today ? "(Today)" : ""}
            </h2>
            <div className={styles["meeting-meta"]}>
                <p>
                    <b>Scheduled Time:</b> {date} at{" "}
                    {Math.abs(parseInt(hour) - 12)}:{minute}:{second}
                    {parseInt(hour) >= 12 ? " PM" : " AM"}
                </p>
                <p>{description}</p>
                <p>
                    <b>Organizer:</b> {organizer.email}
                    {organizer.self === true ? " (You)" : ""}
                </p>
            </div>
            <div className={styles["meeting-buttons"]}>
                {isNotesTaken ? (
                    <Button
                        text="Show Notes"
                        className={styles["show_notes_btn"]}
                        onClickHandler={showNotesButtonHandler}
                    />
                ) : (
                    <NavLinkButton
                        className={styles["notes-btn"]}
                        meeting_id={meeting_id}
                        isEdit={false}
                        text="Take Notes"
                    />
                )}
                <Button
                    text="Delete"
                    className={styles["notes-delete_btn"]}
                    isDisabled={organizer.self === false}
                />
                <LinkButton url={url} text="Join" isDisabled={isDisabled} />
            </div>
            {showNotes &&
                createPortal(
                    <NotesModel
                        notes={notes}
                        notesCloseHandler={closesNotesButtonHandler}
                    />,
                    document.getElementById("notes-model")
                )}
        </div>
    );
};

export default MeetingCard;
