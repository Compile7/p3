import Header from "../Header";
import Footer from "../Footer";
import Hitesh from "../../assets/team/hitesh.png";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
const Team = () => {
const navigate = useNavigate();
  useEffect(() => {
    //fetch team list
    
    return () => {
      
    }
  }, [])
  
  const teamlist = [
    {
      email: 'hitesh.kumawat@loginradius.com',
      name:'Hitesh',
      profilePic: '',
      inviteStatus:'approved',
      managedBy:'Govind',


    },
    {
      email: 'Govind.malviya@loginradius.com',
      name:'Govind',
      profilePic: '',
      inviteStatus:'approved',
      managedBy:'',
      

    },
    {
      email: 'ankit.aabad@loginradius.com',
      name:'Ankit',
      profilePic: '',
      inviteStatus:'pending',
      managedBy:'Govind',
      

    },
    {
      email: 'Hemant.manwani@loginradius.com',
      name:'Hemant',
      profilePic: '',
      inviteStatus:'rejected',
      managedBy:'Govind',
      

    },
    {
      email: 'ankit.aabad@loginradius.com',
      name:'Ankit',
      profilePic: '',
      inviteStatus:'pending',
      managedBy:'Govind',
      

    }
    ,
    {
      email: 'ankit.aabad@loginradius.com',
      name:'Ankit',
      profilePic: '',
      inviteStatus:'pending',
      managedBy:'Govind',
      

    }
    ,
    {
      email: 'ankit.aabad@loginradius.com',
      name:'Ankit',
      profilePic: '',
      inviteStatus:'pending',
      managedBy:'Govind',
      

    }
    ,
    {
      email: 'ankit.aabad@loginradius.com',
      name:'Ankit',
      profilePic: '',
      inviteStatus:'pending',
      managedBy:'Govind',
      

    }
    ,
    {
      email: 'ankit.aabad@loginradius.com',
      name:'Ankit',
      profilePic: '',
      inviteStatus:'pending',
      managedBy:'Govind',
      

    },
    {
      email: 'ankit.aabad@loginradius.com',
      name:'Ankit',
      profilePic: '',
      inviteStatus:'pending',
      managedBy:'Govind',
      

    }
  ]
  return (
    <>
      <Header />
      <main>
        <div>
          <div className="header">
            <h1>Team Members</h1>
            <p>
              Lorem Ipsum is simply dummy text of the printing and typesetting
              industry.
            </p>
          </div>
          <div className="table">
            <table>
              {/* <caption>Statement Summary</caption> */}
              <thead>
                <tr>
                  <th scope="col">Member detail</th>
                  <th scope="col">Invite Status</th>
                  <th scope="col">Manage by</th>
                  <th scope="col">Action</th>
                </tr>
              </thead>
              <tbody>
              {
                teamlist.map((member:any)=>{
                  {console.log(member.managedBy)}
                  return <tr key={member.email}>
                  <td data-label="Member detail">
                    <div className="user-detail">
                      <div className="user-icon">
                        <img src={Hitesh} alt="Hitesh" />
                      </div>
                      <div className="user-info">
                        <h3>{member.email}</h3>
                        <p>{member.name}</p>
                      </div>
                    </div>
                  </td>
                  <td data-label="Invite Status">
                    <span className="badge badge-success">{member.inviteStatus}</span>
                  </td>
                  <td data-label="Manage by">
                    
                    <select name="" id="" className="primary sm" >
                      {teamlist.map((teamMember)=>{
                        return member.name!==teamMember.name? <option key={teamMember.name} value="govind" selected={teamMember.name===member.managedBy}>{teamMember.name}</option>:null
                      })}
                      
                    </select>
                  </td>
                  <td data-label="Action">
                    <div className="btn-group">
                      <button className="btn btn-primary sm" onClick={()=>navigate(`/${member.name}/add-review`)}>Add review</button>
                    </div>
                  </td>
                </tr>
                })
              }
              </tbody>
            </table>
          </div>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Team;
