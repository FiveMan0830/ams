import logo from './logo.svg';
import './App.css';
import TeamManage from './teamManage';
import React, { Component, forwardRef } from "react";
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import {Button} from '@material-ui/core'
import AddMember from './addMember';
import {TeamGetMemberOf} from './middleware';
import axios from 'axios'


// const useStyles = makeStyles((theme) => ({
//   formControl: {
//     margin: theme.spacing(1),
//     minWidth: 120,
//   },
//   selectEmpty: {
//     marginTop: theme.spacing(2),
//   },
// }));

// const classes = useStyles();


class App extends Component{
  constructor(props) {
    super(props)
    this.state = {
      teamList: [],
      teamName : 'SSL LAB',
      username : "ssl1321ois",
      addMemberOpen: false,
      memberList : [],
    }
    this.handleAddMemberClose = this.handleAddMemberClose.bind(this);
    this.handleAddMemberOpen = this.handleAddMemberOpen.bind(this);
    // this.getTeam = this.getTeam.bind(this);
  }

  componentWillMount() {
    const data = {
      Username: this.state.username
    }
    axios.post("http://localhost:8080/team/get/memberOf", data)
      .then(res => {
          this.setState({teamList: res.data})
      })
      .catch(err => {
          console.log(err);
      })

    // axios.post("http://localhost:8080/team/get/name", String Team ID)
    //     .then(res => {
    //         this.setState({team: res.data})
    //     })
    //     .catch(err => {
    //         console.log(err);
    //     })
  }

  handleAddMemberOpen() {
    this.setState({addMemberOpen:true});
  };
  handleAddMemberClose() {
    this.setState({addMemberOpen:false});
  };
  // useEffect(() => {
  //   let temp =  new Array();
  //   const GetTeam = new Promise((resolve,reject) =>{
  //     temp = TeamGetMemberOf(username);
  //     console.log(temp);
  //     resolve();
  //   })
  //   GetTeam.then(setTeamList(temp)).then(console.log(teamList));
  // },[])

  // useEffect(() => {
  //   // STEP 1：在 useEffect 中定義 async function 取名為 fetchData
  //   const fetchData = async () => {
  //     // STEP 2：使用 Promise.all 搭配 await 等待兩個 API 都取得回應後才繼續
  //     const data = await Promise.all([
  //       TeamGetMemberOf(username),
  //     ]);
  //     setTeamList(data);
  //     console.log('data', data);
  //   };

  //   // STEP 5：呼叫 fetchData 這個方法
  //   fetchData();
  // }, []);


render(){
  // const { teamList } = this.state;
  return (
    <div className="App">
      <div className="selector">
         <FormControl>
          <InputLabel id="demo-simple-select-label">Team</InputLabel>
          <Select
            value={this.state.teamName}
            onChange={(e) => {this.setState({teamName: e.target.value})}}
          >
            {
              this.state.teamList.map((team,index) => {
                return(
                  <MenuItem key={index} value={team}>{team}</MenuItem>
                )
              })
            }
          </Select>
        </FormControl>
      </div>
      <div className="team-name">
        {this.state.teamName}
        <Button onClick = {this.handleAddMemberOpen}>ADD</Button>
        <AddMember open={this.state.addMemberOpen} handleClose={this.handleAddMemberClose} memberList={this.state.memberList}/>
      </div>
      <div className="table">
        <TeamManage teamName = {this.state.teamName} username = {this.state.username} > </TeamManage>
      </div>
    </div>
  );
}
  
}

export default App;
