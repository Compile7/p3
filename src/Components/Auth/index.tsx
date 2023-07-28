import {useEffect} from "react";
import "./auth.scss";
import Logo from "../../assets/logo.svg";
import useFetch from "../../Hooks/useFetch";
import { FALLBACK } from "../../utils";
import { useStateContext } from "../../Contexts/contextProvider";
import { Navigate } from "react-router-dom";

const Auth = () => {
  const { handleGoogle, loading} = useFetch(
    `${import.meta.env.VITE_APP_BACKEND_ENDPOINT}/login`
  );

  const {profile} = useStateContext();
  useEffect(() => {
    /* global google */
    const google = (window as any).google;
    if (google) {
      google.accounts.id.initialize({
        client_id: import.meta.env.VITE_APP_GG_APP_ID,
        callback: handleGoogle,
      });

      google.accounts.id.renderButton(document.getElementById("loginDiv"), {
        // type: "standard",
        theme: "filled_blue",
        // size: "small",
        text: "signin_with",

        // shape: "pill",
      });

      google.accounts.id.prompt()
      console.log(google.accounts.id)
    }
  }, [handleGoogle]);

  return (
    <>
    {loading?FALLBACK:
      profile? profile.OrgId ===0?
      <>{console.log("here")}
      <Navigate replace to="/onboarding" />
      </>:<> 
      {console.log("there")}
      <Navigate replace to="/view-review" />
      </>:
        <div className={`App ${profile ? "hide" : ""}`}>
          <div className="login">
            <div className="logo">
              <a href="/">
                <img src={Logo} alt="Logo" />
              </a>
            </div>
            <div className="header">
              <h1 className="title">Welcome to P3</h1>
              <p>Login with your Google account to continue</p>
            </div>
            <div id="loginDiv" className="bg-gray" data-type="standard"></div>
          </div>
        </div>
      }
    </>
  );
};

export default Auth;
