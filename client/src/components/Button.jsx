import React from "react";

import styles from "../styles/button.module.css";

const Button = (props) => {
  const classes = `${styles["button"]} ${props.className}`;

  const buttonClickHandler = () => {
    if (props.onClickHandler) props.onClickHandler();
  };

  return (
    <button className={classes} onClick={buttonClickHandler}>
      {props.text}
    </button>
  );
};

export default Button;
