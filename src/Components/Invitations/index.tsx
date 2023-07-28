import Header from "../Header";
import Footer from "../Footer";
import Hitesh from "../../assets/team/hitesh.png";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useStateContext } from "../../Contexts/contextProvider";
const Team = () => {
const navigate = useNavigate();

const [loading, setLoading] = useState(false)
const [email, setEmail] = useState("")
const { profile} = useStateContext();
  useEffect(() => {
    //fetch team list
    
    return () => {
      
    }
  }, [])
  
  const resendInvite = (email:string) =>{

    setLoading(true)

    console.log(email)

    // resend invite API call

    setLoading(false)
  }
  const inviteEmployees = (e:any) =>{
    e.preventDefault();

    setLoading(true)

    console.log(email)

    //apicall with organization name

    setLoading(false)
    navigate("/team")
  }

  const teamlist = [
    {
      email: 'hitesh.kumawat@loginradius.com',
      inviteStatus:'approved',


    },
    {
      email: 'Govind.malviya@loginradius.com',
      inviteStatus:'approved',
      

    },
    {
      email: 'ankit.aabad@loginradius.com',
      inviteStatus:'pending',
      

    },
    {
      email: 'Hemant.manwani@loginradius.com',
      inviteStatus:'rejected',
      

    },
    {
      email: 'ankit.aabad@loginradius.com',
      inviteStatus:'pending',

    }
    ,
    {
      email: 'ankit.aabad@loginradius.com',
      inviteStatus:'pending',
      

    }
    ,
    {
      email: 'ankit.aabad@loginradius.com',
      inviteStatus:'pending',
      

    }
    ,
    {
      email: 'ankit.aabad@loginradius.com',
      inviteStatus:'pending',
    }
  ]
  return (
    <>
      <Header />
      <main>
        <div>
          <div className="header">
            <h1>Invitations</h1>
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
                        {/* <p>{member.name}</p> */}
                      </div>
                    </div>
                  </td>
                  <td data-label="Invite Status">
                    <span className="badge badge-success">{member.inviteStatus}</span>
                  </td>
                  {/* <td data-label="Manage by">
                    
                    <select name="" id="" className="primary sm" >
                      {teamlist.map((teamMember)=>{
                        return member.name!==teamMember.name? <option key={teamMember.name} value="govind" selected={teamMember.name===member.managedBy}>{teamMember.name}</option>:null
                      })}
                      
                    </select>
                  </td> */}
                  <td data-label="Action">
                    <div className="btn-group">
                      <button className="btn btn-primary sm" onClick={()=>resendInvite(member.email)}>Resend Invitation</button>
                    </div>
                  </td>
                </tr>
                })
              }
              </tbody>
            </table>
          </div>
          <h1 style={{marginTop: '40px', marginBottom:'40px'}}>Add Members</h1>
          <div className="onboarding" >
            <div className="description" style={{width:'100%'}}>
          <div className="import-data" style={{width:'100%'}}>
          <div className="form-control" style={{width:'67%', height:'60px'}}>
                    
                    <textarea
                      name="review"
                      id="review"
                      required
                      placeholder="Ex. hitesh@loginradius.com, ankit@loginradius.com"
                      style={{width:'100%', height:'60px'}}
                    ></textarea>
                  </div>
          <div className="form-control">
                    <div className="btn-group">
                      <button onClick={inviteEmployees} type="submit" className="btn btn-primary">
                        Add members
                      </button>
                      
                    </div>
                  </div>
                  </div>
                  </div>
                  </div>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Team;
