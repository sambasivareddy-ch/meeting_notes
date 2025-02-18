import React, { useState, useEffect } from "react";
import { useDispatch } from "react-redux";

import MeetingCard from "../components/MeetingCard";
import Button from "../components/Button";
import { logout } from "../store/sessionSlice";

import styles from "../styles/meetings.module.css";

const Meetings = (props) => {
    const actionDispatcher = useDispatch();
    const [shouldShowPastMeetings, setShouldShowPastMeetings] = useState(false);
    const [userEmail, setUserEmail] = useState(null);
    const [meetings, setMeetings] = useState([]);

    useEffect(() => {
        const getMeetings = async () => {
            const response = await fetch("http://localhost:8080/meetings/", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            setUserEmail(data.email);
            setMeetings(data.meetings.items);
        };

        getMeetings();
    }, []);

    const refreshButtonClickHandler = async () => {
        try {
            const response = await fetch("http://localhost:8080/meetings/reload", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data = await response.json();
            setMeetings(data.meetings.items);
        } catch (error) {
            console.error(error);
        }
    };

    const logoutButtonClickHandler = async () => {
        try {
            const response = await fetch("http://localhost:8080/logout", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            });
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            await response.json();

            actionDispatcher(logout()); // Logout the session
            window.location.href = "/"; // Redirect to the signin page
        } catch (error) {
            console.error(error);
        }
    };

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
                    </div>
                    <div className={styles["refresh-logout"]}>
                        <Button text="refresh" onClickHandler={refreshButtonClickHandler} />
                        <Button text="logout" onClickHandler={logoutButtonClickHandler} />
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
                        meetings &&
                        meetings.length > 0 &&
                        getScheduledMeetings(meetings).length === 0 && (
                            <h2 className={styles["no-meetings_message"]}>
                                Hurray! No Meetings Scheduled
                            </h2>
                        )}
                    {shouldShowPastMeetings &&
                        meetings &&
                        meetings.length > 0 &&
                        getScheduledMeetings(meetings).length ===
                            meetings.length && (
                            <h2 className={styles["no-meetings_message"]}>
                                No Meetings held in the past
                            </h2>
                        )}
                    {!shouldShowPastMeetings &&
                        meetings &&
                        meetings.length > 0 &&
                        getScheduledMeetings(meetings).length > 0 &&
                        getScheduledMeetings(meetings).map((meeting) => (
                            <MeetingCard
                                meeting_id={meeting.id}
                                title={meeting.summary}
                                start_time={meeting.start.dateTime}
                                url={meeting.hangoutLink}
                                description={meeting.summary}
                                isNotesTaken={meeting.notes !== ""}
                                organizer={meeting.organizer}
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
                        meetings &&
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
                                    isNotesTaken={meeting.notes !== ""}
                                    isDisabled={true}
                                    is_today={false}
                                    organizer={meeting.organizer}
                                    key={Math.random()}
                                />
                            ))}
                </div>
            </div>
        </div>
    );
};

export default Meetings;
