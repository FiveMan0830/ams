import React,{ seCallback, useState , useEffect} from 'react';
import Toolbar from '@material-ui/core/Toolbar';
import AppBar from '@material-ui/core/AppBar';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import './Appbar.css';
import Popover from '@material-ui/core/Popover';
import Avatar from '@material-ui/core/Avatar';
import Profile from './Profile';
import { Component } from 'react';
import { withStyles } from '@material-ui/core';
import { Link } from 'react-router-dom';


const useStyles = (theme) => ({
  root: {
    display: 'flex',
  },
  appBar: {
    [theme.breakpoints.up('sm')]: {
      width: `calc(100%)`,
      zIndex: theme.zIndex.drawer + 1,
      background:'#303030',
    },
  },
  menuButton: {
    marginRight: theme.spacing(2),
    [theme.breakpoints.up('sm')]: {
      display: 'none',
    },
  },
  iconColor: {
    color: '#fff' ,
    backgroundColor: '#00C6CF' ,
  },
  toolbar: theme.mixins.toolbar,
  toolBar: {
    justifyContent:'space-between',
  },
  popover: {
    width: '150%',
    height: '150%',
  },
  select: {
    width: '150px',
    height: '40px',
    margin: '10px'
  }

});

class MyAppBar extends Component {
  constructor(props) {
    super(props)
    this.handleDrawerToggle = this.handleDrawerToggle.bind(this)
    this.handleProfileClick = this.handleProfileClick.bind(this)
    this.handleProfileClose = this.handleProfileClose.bind(this)
    this.state = { 
      anchorProfile: null, 
      team: localStorage.getItem("teamName"),
      displayName : localStorage.getItem("displayName"),
      
    };
  }

  handleDrawerToggle = () => {
    this.props.handleDrawerToggle();
  };

  handleProfileClick = (event) => {
    this.setState({anchorProfile: event.currentTarget});
  };

  handleProfileClose = () => {
    this.setState({anchorProfile: null});
  };

  render() {
    const { classes , history } = this.props;
    
    return (
      <div>
        <AppBar position="fixed" className={classes.appBar}>
          <Toolbar className={classes.toolBar}> 
            <IconButton
              color="inherit"
              aria-label="open drawer"
              edge="start"
              onClick={this.handleDrawerToggle}
              className={classes.menuButton}
            >
            <MenuIcon />
            </IconButton>
            <div className="collapse navbar-collapse" id="navbarContent">
              <ul className="navbar-nav">
                <li className="nav-item">
                  <Link className="nav-link" to="/">
                    Home
                  </Link>
                </li>
                <li className="nav-item">
                  <Link className="nav-link" to="/groupManage">
                    GroupManage
                  </Link>
                </li>
              </ul>
            </div>
            <div className="profile-btn" >
              <Avatar className={classes.iconColor}  alt={this.state.displayName} src="/broken-image.jpg" onClick={this.handleProfileClick} id="profile-icon"/>
              <Popover
                open={Boolean(this.state.anchorProfile)}
                margin='100px'
                anchorEl={this.state.anchorProfile}
                onClose={this.handleProfileClose}
                anchorOrigin={{
                  vertical: 'bottom',
                  horizontal: 'right',
                }}
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                className={classes.popover}
              >
                <Profile></Profile>
              </Popover>
            </div>
          </Toolbar>
        </AppBar>

      </div>
    
    );
  }

}

export default withStyles(useStyles,{withTheme: true})(MyAppBar)

