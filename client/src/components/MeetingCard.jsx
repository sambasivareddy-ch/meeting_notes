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
}) => {
  return (
    <div className={styles["meeting-card_wrapper"]}>
      <h2 className={styles["meeting-title"]}>
        {title} {is_today ? "(Today)" : ""}
      </h2>
      <div className={styles["meeting-meta"]}>
        <p>Scheduled Time: {start_time}</p>
        <p>{description}</p>
      </div>
      <div className={styles["meeting-buttons"]}>
        {isNotesTaken ? (
          <Button className={styles["notes-btn"]} text="Show Notes" />
        ) : (
          <NavLinkButton
            className={styles["notes-btn"]}
            meeting_id={meeting_id}
            text="Take Notes"
          />
        )}
        <LinkButton url={url} text="Join" isDisabled={isDisabled} />
      </div>
    </div>
  );
};

export default MeetingCard;
