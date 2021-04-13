import logo from './logo.svg';
import './App.css';
import TeamManage from './teamManage';
import React, { useState, useEffect }  from 'react';
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import axios from 'axios'
import moment from 'moment'
import {Button} from '@material-ui/core'

const useStyles = makeStyles((theme) => ({
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
}));




function App() {
  const classes = useStyles();
  const [team, setTeam] = React.useState('');
  const [teamList, setTeamList] = React.useState([]);
  const username = "ssl1321ois"


  const handleSelectTeam = (event) => {
    setTeam(event.target.value);
  };

  // const getGroups = () => {
  //   const data = {
  //     Username: username
  //   }
  //   axios.post("http://localhost:8080/team/get/member", data)
  //       .then(res => {
  //           // const queryStr = decodeURIComponent(window.location.search.substring(1)).split("&")
  //           //     const query = {}
  //           //     queryStr.forEach((item, i) => {
  //           //         itemPair = item.split("=")
  //           //         query[itemPair[0]] = itemPair[1]
  //           //     })
  //           console.log(res);
  //           setTeamList(res);
  //           console.log(teamList);
  //       })
  //       .catch(err => {
  //           console.log(err);
  //       })
  // }

  useEffect(() => {
    fetch("http://localhost:8080/team/get/member")
      .then(res => res.json())
      .then(
        (result) => {
          console.log(result);
            setTeamList(result);
            console.log(teamList);
        },
        // Note: it's important to handle errors here
        // instead of a catch() block so that we don't swallow
        // exceptions from actual bugs in components.
        (error) => {
          console.log(error);
        }
      )
  }, [])

  return (
    <div className="App">
      {/* <Button onClick={getGroups(username)}
                >
                Add Log
              </Button> */}
      <div className="selector">
         <FormControl className={classes.formControl}>
          <InputLabel id="demo-simple-select-label">Team</InputLabel>
          <Select
            value={team}
            onChange={handleSelectTeam}
          >
            {/* {
              props.groupList.map((group,index) => {
                return(
                  <MenuItem key={index} value={group}>{group.teamName}</MenuItem>
                )
              })
            } */}
          </Select>
        </FormControl>
      </div>
      <div className="team-name">
        OIS
      </div>
      <div className="table">
        <TeamManage />
      </div>
    </div>
  );
}

export default App;
