import React, { useState } from "react";

import MeetingCard from "../components/MeetingCard";
import Button from "../components/Button";
import styles from "../styles/meetings.module.css";

const DUMMY_MEETINGS = [
  {
    id: Math.random(),
    title: "Meeting 1",
    start_time: `${new Date(2024, 8, 13).toLocaleString()}`,
    url: "http://localhost:3000",
    description: "Google Meet 1 Description",
    isNotesTaken: false,
    notes: "",
  },
  {
    id: Math.random(),
    title: "Meeting 2",
    start_time: `${new Date(2024, 8, 14).toLocaleString()}`,
    url: "http://localhost:3000",
    description: "Google Meet 2 Description",
    isNotesTaken: false,
    notes: "",
  },
  {
    id: Math.random(),
    title: "Meeting 3",
    start_time: `${new Date(2024, 8, 15).toLocaleString()}`,
    url: "http://localhost:3000",
    description: "Google Meet 3 Description",
    isNotesTaken: false,
    notes: "",
  },
  {
    id: Math.random(),
    title: "Meeting 4",
    start_time: `${new Date(2024, 8, 16).toLocaleString()}`,
    url: "http://localhost:3000",
    description: "Google Meet 4 Description",
    isNotesTaken: false,
    notes: "",
  },
  {
    id: Math.random(),
    title: "Meeting 0",
    start_time: `${new Date(2024, 7, 16).toLocaleString()}`,
    url: "http://localhost:3000",
    description: "Google Meet 0 Description",
    isNotesTaken: false,
    notes: "",
  },
];

const Meetings = (props) => {
  const [shouldShowPastMeetings, setShouldShowPastMeetings] = useState(false);

  const toggleMeetingsDisplayStateToFalse = () => {
    setShouldShowPastMeetings(false);
  }

  const toggleMeetingsDisplayStateToTrue = () => {
    setShouldShowPastMeetings(true);
  }

  const getScheduledMeetings = (meetings) => {
    return meetings.filter((meeting) => {
      return checkIsStartTimeToday(meeting.start_time) >= 1;
    });
  };

  const checkIsStartTimeToday = (string) => {
    let [day, month, year] = string.split(",")[0].split("/").map(Number);
    let parsedDate = new Date(year, month - 1, day);
    let today = new Date();
    today.setHours(0, 0, 0, 0);

    if (parsedDate.getTime() === today.getTime()) {
      return 1;
    } else if (parsedDate.getTime() < today.getTime()) {
      return 0;
    }
    return 2;
  };

  return (
    <div className={styles["meetings-page_wrapper"]}>
      <div className={styles["meetings-wrapper"]}>
        <div className={styles["meetings-header"]}>
          <h1>My Meetings</h1>
          <div className={styles["user-info"]}>
            <p>on sambasivareddy@gmail.com</p>
            <a href="/">logout</a>
          </div>
        </div>
        <div className={styles["nav-buttons-wrapper"]}>
          <Button
            text="Scheduled Meetings"
            className={shouldShowPastMeetings === true ? styles["nav-button"]: styles["nav-button_selected"]}
            onClickHandler={toggleMeetingsDisplayStateToFalse}
          />
          <Button
            text="Past Meetings"
            className={shouldShowPastMeetings === false ? styles["nav-button"]: styles["nav-button_selected"]}
            onClickHandler={toggleMeetingsDisplayStateToTrue}
          />
        </div>
        <div className={styles["meetings"]}>
          {!shouldShowPastMeetings &&
            DUMMY_MEETINGS.length > 0 &&
            getScheduledMeetings(DUMMY_MEETINGS).length === 0 && (
              <h2 className={styles["no-meetings_message"]}>
                Hurray! No Meetings Scheduled
              </h2>
            )}
          {shouldShowPastMeetings &&
            DUMMY_MEETINGS.length > 0 &&
            getScheduledMeetings(DUMMY_MEETINGS).length ===
              DUMMY_MEETINGS.length && (
              <h2 className={styles["no-meetings_message"]}>
                No Meetings held in the past
              </h2>
            )}
          {!shouldShowPastMeetings &&
            DUMMY_MEETINGS.length > 0 &&
            getScheduledMeetings(DUMMY_MEETINGS).length > 0 &&
            getScheduledMeetings(DUMMY_MEETINGS).map((meeting) => (
              <MeetingCard
                meeting_id={meeting.id}
                title={meeting.title}
                start_time={meeting.start_time}
                url={meeting.url}
                description={meeting.description}
                isNotesTaken={meeting.isNotesTaken}
                isDisabled={checkIsStartTimeToday(meeting.start_time) !== 1}
                is_today={checkIsStartTimeToday(meeting.start_time) === 1}
                key={Math.random()}
              />
            ))}
          {shouldShowPastMeetings &&
            DUMMY_MEETINGS.length > 0 &&
            DUMMY_MEETINGS.filter((meeting) => {
              return checkIsStartTimeToday(meeting.start_time) === 0;
            }).map((meeting) => (
              <MeetingCard
                meeting_id={meeting.id}
                title={meeting.title}
                start_time={meeting.start_time}
                url={meeting.url}
                description={meeting.description}
                isNotesTaken={meeting.isNotesTaken}
                isDisabled={true}
                is_today={false}
                key={Math.random()}
              />
            ))}
        </div>
      </div>
    </div>
  );
};

export default Meetings;
