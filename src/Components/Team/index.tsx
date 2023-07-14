import Header from "../Header";
import Footer from "../Footer";
import Hitesh from "../../assets/team/hitesh.png";
const Team = () => {
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
                <tr>
                  <td data-label="Member detail">
                    <div className="user-detail">
                      <div className="user-icon">
                        <img src={Hitesh} alt="Hitesh" />
                      </div>
                      <div className="user-info">
                        <h3>Hitesh.kumawat@loginradius.com</h3>
                        <p>Product Designer</p>
                      </div>
                    </div>
                  </td>
                  <td data-label="Invite Status">
                    <span className="badge badge-success">Approved</span>
                  </td>
                  <td data-label="Manage by">
                    <select name="" id="" className="primary sm">
                      <option value="govind">Govind</option>
                      <option value="sudhanshu">Sudhanshu</option>
                      <option value="sudhey">Sudhey</option>
                    </select>
                  </td>
                  <td data-label="Action">
                    <div className="btn-group">
                      <button className="btn btn-primary sm">Add review</button>
                    </div>
                  </td>
                </tr>
                <tr>
                  <td data-label="Member detail">
                    <div className="user-detail">
                      <div className="user-icon">
                        <img src={Hitesh} alt="Hitesh" />
                      </div>
                      <div className="user-info">
                        <h3>Hitesh.kumawat@loginradius.com</h3>
                        <p>Product Designer</p>
                      </div>
                    </div>
                  </td>
                  <td data-label="Invite Status">
                    <span className="badge badge-error">Declined</span>
                  </td>
                  <td data-label="Manage by">
                    <select name="" id="" className="primary sm">
                      <option value="govind">Govind</option>
                      <option value="sudhanshu">Sudhanshu</option>
                      <option value="sudhey">Sudhey</option>
                    </select>
                  </td>
                  <td data-label="Action">
                    <div className="btn-group">
                      <button className="btn btn-primary sm">Add review</button>
                    </div>
                  </td>
                </tr>
                <tr>
                  <td data-label="Member detail">
                    <div className="user-detail">
                      <div className="user-icon">
                        <img src={Hitesh} alt="Hitesh" />
                      </div>
                      <div className="user-info">
                        <h3>Hitesh.kumawat@loginradius.com</h3>
                        <p>Product Designer</p>
                      </div>
                    </div>
                  </td>
                  <td data-label="Invite Status">
                    <span className="badge badge-wraning">Pending</span>
                  </td>
                  <td data-label="Manage by">
                    <select name="" id="" className="primary sm">
                      <option value="govind">Govind</option>
                      <option value="sudhanshu">Sudhanshu </option>
                      <option value="sudhey">Sudhey</option>
                    </select>
                  </td>
                  <td data-label="Action">
                    <div className="btn-group">
                      <button className="btn btn-primary sm">Add review</button>
                    </div>
                  </td>
                </tr>
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
