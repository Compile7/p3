import { useCallback } from "react";
import Logo from "../../assets/logo.svg";
import { useStateContext } from "../../Contexts/contextProvider";
import { useNavigate } from "react-router-dom";
const Header = () => {

  const{setProfile} = useStateContext();
  const navigate = useNavigate()
  const onLogoutSuccess = useCallback(() => {
    setProfile(null);
    localStorage.removeItem("P3AccessToken")
    navigate("/")
  }, []);
  
  return (
    <header>
      <div>
        <div className="logo">
          <a href="/">
            <img src={Logo} alt="Logo" />
          </a>
        </div>
        <nav>
          <ul>
            <li>
              <a href="#">Menu01</a>
            </li>
            <li>
              <a href="#">Menu02</a>
            </li>
            <li>
              <a href="#" onClick={onLogoutSuccess}>Logout</a>
            </li>
          </ul>
        </nav>
      </div>
    </header>
  );
};
export default Header;
