import React, { useState } from 'react';
import './UserProfile.css';
import {
  Avatar,
  Button
} from '@material-ui/core';
import { connect } from 'react-redux';

function Profile(props) {
  const [displayName] = useState(localStorage.getItem("displayName"));
  const [username] = useState(localStorage.getItem("cn"));
  const [teamName] = useState(localStorage.getItem("teamName"));

  const logout = () => {
    localStorage.clear();
    window.location.href = '/';
  }

  const timelog = () => {
    var redirectURL = process.env.REACT_APP_TIMELOG + "?";
    redirectURL += "cn=" + username + "&";
    redirectURL += "displayName=" + displayName + "&"; 
    redirectURL += "teamName=" + teamName; 
    window.open(redirectURL);
  }

  return (
    <center>
      <div className="profile-box">
        <div className="profile-split"></div>
        <Avatar className="avatar-name" alt={displayName} src="/broken-image.jpg"/>
        <div className="profile-split"></div>
        <div>
          <p>{displayName}</p>
          <p>{username}</p>
        </div>
        <div className="profile-split"></div>
        <div className="btn-div">
          <Button 
            className = "ams-btn" 
            variant="contained" 
            color="primary" 
            style = {{minWidth : "6vw"}}
            onClick = {timelog}
            >
              TIMELOG
          </Button>
          <Button 
            className = "logout-btn" 
            variant="contained" 
            color="primary" 
            style = {{minWidth : "6vw", margin : "10px"}}
            onClick = {logout}
            >
              LOGOUT
          </Button>
        </div>
        
      </div>
    </center>
  )
};


export default (Profile);
