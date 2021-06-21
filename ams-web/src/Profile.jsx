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

    var amsURL = process.env.REACT_APP_AMS_LOGIN_URL
    amsURL += '?' + encodeURIComponent('redirect_url=' + process.env.REACT_APP_WEB)
    window.location.replace(amsURL);
  }

  const timelog = () => {
    var redirectURL = process.env.REACT_APP_TIMELOG + "?";
    redirectURL += "uid=" + localStorage.getItem('uid') + "&";
    redirectURL += "cn=" + localStorage.getItem('cn') + "&";
    redirectURL += "sn=" + localStorage.getItem('sn') + "&";
    redirectURL += "givenName=" + localStorage.getItem('givenName') + "&";
    redirectURL += "displayName=" + localStorage.getItem('displayName') + "&";
    redirectURL += "mail=" + localStorage.getItem('mail');
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
