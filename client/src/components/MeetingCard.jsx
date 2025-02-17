import React from "react";

import LinkButton from "./LinkButton";
import NavLinkButton from "./NavLinkButton";
import Button from "./Button";
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
}) => {
    const [date, timeFrame] = start_time.split("T");
    const time = timeFrame.split("+")[0];
    const [hour, minute, second] = time.split(":");

    return (
        <div className={styles["meeting-card_wrapper"]}>
            <h2 className={styles["meeting-title"]}>
                {title} {is_today ? "(Today)" : ""}
            </h2>
            <div className={styles["meeting-meta"]}>
                <p>
                    <b>Scheduled Time:</b> {date} at {Math.abs(parseInt(hour) - 12)}:
                    {minute}:{second}
                    {parseInt(hour) >= 12 ? " PM" : " AM"}
                </p>
                <p>{description}</p>
                <p><b>Organizer:</b> {organizer.email}{ organizer.self === true ? " (You)": "" }</p>
            </div>
            <div className={styles["meeting-buttons"]}>
                {isNotesTaken ? (
                    <NavLinkButton
                        className={styles["notes-btn"]}
                        meeting_id={meeting_id}
                        isEdit={true}
                        text="Show Notes"
                    />
                ) : (
                    <NavLinkButton
                        className={styles["notes-btn"]}
                        meeting_id={meeting_id}
                        isEdit={false}
                        text="Take Notes"
                    />
                )}
                <Button text="Delete" className={styles["notes-delete_btn"]} isDisabled={ organizer.self === false } />
                <LinkButton url={url} text="Join" isDisabled={isDisabled} />
            </div>
        </div>
    );
};

export default MeetingCard;
