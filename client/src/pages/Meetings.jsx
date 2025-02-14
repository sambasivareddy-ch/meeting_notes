import React, { useState, useEffect } from "react";

import MeetingCard from "../components/MeetingCard";
import Button from "../components/Button";

import styles from "../styles/meetings.module.css";

const Meetings = (props) => {
    const [shouldShowPastMeetings, setShouldShowPastMeetings] = useState(false);
    const [userEmail, setUserEmail] = useState(null);
    const [meetings, setMeetings] = useState([]);

    useEffect(() => {
        const getData = async () => {
            const response = await fetch(
                "http://localhost:8080/user/getemail",
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
            setUserEmail(data.email);
        };

        const getMeetings = async () => {
            const response = await fetch(
                "http://localhost:8080/meeting/getmeetings",
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
            setMeetings(data.meetings.items);
        };

        getData();
        getMeetings();
    }, []);

    const toggleMeetingsDisplayStateToFalse = () => {
        setShouldShowPastMeetings(false);
    };

    const toggleMeetingsDisplayStateToTrue = () => {
        setShouldShowPastMeetings(true);
    };

    const getScheduledMeetings = (meetings) => {
        return meetings.filter((meeting) => {
            return checkIsStartTimeToday(meeting.start.dateTime) >= 1;
        });
    };

    const checkIsStartTimeToday = (string) => {
        let [year, month, day] = string.split("T")[0].split("-").map(Number);
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
                        <p>on {userEmail}</p>
                        <a href="/">logout</a>
                    </div>
                </div>
                <div className={styles["nav-buttons-wrapper"]}>
                    <Button
                        text="Scheduled Meetings"
                        className={
                            shouldShowPastMeetings === true
                                ? styles["nav-button"]
                                : styles["nav-button_selected"]
                        }
                        onClickHandler={toggleMeetingsDisplayStateToFalse}
                    />
                    <Button
                        text="Past Meetings"
                        className={
                            shouldShowPastMeetings === false
                                ? styles["nav-button"]
                                : styles["nav-button_selected"]
                        }
                        onClickHandler={toggleMeetingsDisplayStateToTrue}
                    />
                </div>
                <div className={styles["meetings"]}>
                    {!shouldShowPastMeetings &&
                        meetings.length > 0 &&
                        getScheduledMeetings(meetings).length === 0 && (
                            <h2 className={styles["no-meetings_message"]}>
                                Hurray! No Meetings Scheduled
                            </h2>
                        )}
                    {shouldShowPastMeetings &&
                        meetings.length > 0 &&
                        getScheduledMeetings(meetings).length ===
                            meetings.length && (
                            <h2 className={styles["no-meetings_message"]}>
                                No Meetings held in the past
                            </h2>
                        )}
                    {!shouldShowPastMeetings &&
                        meetings.length > 0 &&
                        getScheduledMeetings(meetings).length > 0 &&
                        getScheduledMeetings(meetings).map((meeting) => (
                            <MeetingCard
                                meeting_id={meeting.id}
                                title={meeting.summary}
                                start_time={meeting.start.dateTime}
                                url={meeting.hangoutLink}
                                description={meeting.summary}
                                isNotesTaken={false}
                                isDisabled={
                                    checkIsStartTimeToday(
                                        meeting.start.dateTime
                                    ) !== 1
                                }
                                is_today={
                                    checkIsStartTimeToday(
                                        meeting.start.dateTime
                                    ) === 1
                                }
                                key={Math.random()}
                            />
                        ))}
                    {shouldShowPastMeetings &&
                        meetings.length > 0 &&
                        meetings
                            .filter((meeting) => {
                                return (
                                    checkIsStartTimeToday(
                                        meeting.start.dateTime
                                    ) === 0
                                );
                            })
                            .map((meeting) => (
                                <MeetingCard
                                    meeting_id={meeting.id}
                                    title={meeting.summary}
                                    start_time={meeting.start.dateTime}
                                    url={meeting.hangoutLink}
                                    description={meeting.summary}
                                    isNotesTaken={false}
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
