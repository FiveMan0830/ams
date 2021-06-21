import React, { Component } from "react";
import axios from 'axios'

class Create extends Component {
  constructor(props) {
    super(props)
    this.state = {
      teamName : "",
      memberList: [],
      leaderName:"",
      addMemberOpen:false,
    }
    this.initialize = this.initialize.bind(this)
  }

  initialize(){
    console.log("init")
    const data = {
      GroupName: this.props.teamName
    }
    axios.post("http://localhost:8080/team/get/member/name", data)
        .then(res => {
          const result = []
          res.data.map((member)=>{
            result.push({username:member.username,name:member.displayname})
          })
          this.setState({memberList:result})
        })
        .catch(err => {
            console.log(err);
        })

    axios.post("http://localhost:8080/team/get/leader", data)
        .then(res => {
          this.setState({leaderName:res.data})
        })
        .catch(err => {
            console.log(err);
        })
  }
  componentWillMount() {
    this.initialize();
  }

  componentDidUpdate(prevProps) {
    if (this.props.teamName !== prevProps.teamName) {
      this.initialize();
    }
  }
  
  render() {
    return (
      <div>
          
        
      </div>
    );
  }
}

export default Create