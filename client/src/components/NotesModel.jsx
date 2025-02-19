import React from 'react';

import Button from './Button';
import CloseIcon from '@mui/icons-material/Close';
import styles from '../styles/shownotes.module.css';

const NotesModel = (props) => {
    const buttonClickHandler = () => {
        props.notesCloseHandler()
    }
    return (
        <div className={styles['notes-wrapper']}>
            <div className={styles['main-container']}>
                <div className={styles['close-button']}>
                    <Button text={<CloseIcon />} onClickHandler={buttonClickHandler} />
                </div>
                <div className={styles['main']} dangerouslySetInnerHTML={{ __html: props.notes}}></div>
            </div>
        </div>
    )
}

export default NotesModel;