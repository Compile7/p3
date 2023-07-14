import { useCallback, useEffect, useState } from "react";
import "./auth.scss";
import Logo from "../../assets/logo.svg";
import { User } from "../User"; // component display user (see detail on /example directory)
// import { LoginSocialGoogle } from "reactjs-social-login";

// // CUSTOMIZE ANY UI BUTTON
// import {
//   //   FacebookLoginButton,
//   GoogleLoginButton,
//   //   GithubLoginButton,
//   //   AmazonLoginButton,
//   //   InstagramLoginButton,
//   //   LinkedInLoginButton,
//   //   MicrosoftLoginButton,
//   //   TwitterLoginButton,
//   //   AppleLoginButton,
// } from "react-social-login-buttons";
import useFetch from "../../Hooks/useFetch";

// import { ReactComponent as PinterestLogo } from './assets/pinterest.svg'
// import { ReactComponent as TiktokLogo } from './assets/tiktok.svg'

// REDIRECT URL must be same with URL where the (reactjs-social-login) components is locate
// MAKE SURE the (reactjs-social-login) components aren't unmounted or destroyed before the ask permission dialog closes
// const REDIRECT_URI = window.location.href;

const Auth = () => {
  const { handleGoogle, loading, error } = useFetch(
    "http://localhost:5152/login"
  );

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

      // google.accounts.id.prompt()
    }
  }, [handleGoogle]);

  const [provider, setProvider] = useState("");
  const [profile, setProfile] = useState<any>();

  const onLogoutSuccess = useCallback(() => {
    setProfile(null);
    setProvider("");
    alert("logout success");
  }, []);

  return (
    <>
      {provider && profile ? (
        <User
          provider={provider}
          profile={profile}
          onLogout={onLogoutSuccess}
        />
      ) : (
        <div className={`App ${provider && profile ? "hide" : ""}`}>
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
            {/* <LoginSocialGoogle
              isOnlyGetToken
              typeResponse="idToken"
              
              client_id={import.meta.env.VITE_APP_GG_APP_ID || ""}
              onLoginStart={onLoginStart}
              onResolve={({ provider, data }: any) => {
                console.log(provider);
                console.log(data);
                setProvider(provider);
                setProfile(data);
              }}
              onReject={(err: any) => {
                console.log(err);
              }}
            >
              <GoogleLoginButton />
            </LoginSocialGoogle> */}
          </div>
        </div>
      )}
    </>
  );
};

export default Auth;
